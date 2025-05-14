package application

import (
	"context"
	"errors"
	"log"
	"time"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"

	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/shopnest/user-service/config"
	"github.com/shopnest/user-service/internal/domain"
	"github.com/shopnest/user-service/internal/repository"
)

type AuthService interface {
	CreateAccount(ctx context.Context, email, password, role string) (string, error)
	Login(ctx context.Context, email, password string) (string, error)
	VerifyEmail(ctx context.Context, token string) error
	GetAccount(ctx context.Context, authId string) (*domain.Auth, error)
	RefreshToken(ctx context.Context, authId string) (string, error)
}

type authService struct {
	repo       repository.AuthRepository
	redis      *redis.Client
	rabbitMQ   *amqp091.Channel
	smtpConfig config.SMTPConfig
	jwtSecret  string
}

func NewAuthService(
	repo repository.AuthRepository,
	redis *redis.Client,
	rabbitMQ *amqp091.Channel,
	smtpConfig config.SMTPConfig,
	jwtSecret string,

) AuthService {
	return &authService{
		repo:       repo,
		redis:      redis,
		rabbitMQ:   rabbitMQ,
		smtpConfig: smtpConfig,
		jwtSecret:  jwtSecret,
	}
}

func (a *authService) CreateAccount(ctx context.Context, email, password, role string) (string, error) {
	authID := uuid.New().String()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	auth := &domain.Auth{
		AuthID:       authID,
		UserID:       "", // already forced in the repository
		UserType:     role,
		Email:        email,
		PasswordHash: string(passwordHash),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Save to database
	if err := a.repo.CreateAccount(ctx, auth); err != nil {
		log.Printf("Failed to save a auth to database: %v", err)
		return "", err
	}

	// Generate verification token
	token := uuid.New().String()

	// Store token in Redis with 1-hour expiration
	redisKey := "verification:token:" + token
	if err := a.redis.SetEx(ctx, redisKey, auth.AuthID, 1*time.Hour).Err(); err != nil {
		log.Printf("Failed to store verification token: %v", err)
		return "", err
	}

	emailBody := []byte(`{"email": "` + email + `", "token": "` + token + `"}`)
	// NOTE: change this "email_queue" to manager_approval_queue
	err = a.rabbitMQ.Publish("", "email_queue", false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        emailBody,
	})
	if err != nil {
		log.Printf("Failed to publish email: %v", err)
		return "", err
	}

	return auth.AuthID, nil
}

func (a *authService) Login(ctx context.Context, email, password string) (string, error) {
	// Get user and auth data
	auth, err := a.repo.GetAccountByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("invalid credentials: ", err)
	}

	// Initially, this was enforcing user to verify on login, but the new flow
	// is to allow them execute actions like browse, add to cart but critical
	// endpoints would check for verifcation status
	// // Check if email is verified (status must be "active")
	// if !auth.Verified {
	// 	return "", errors.New("email not verified")
	// }

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(auth.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Create JWT claims
	claims := jwt.MapClaims{
		"sub":      auth.AuthID,
		"userId":   auth.UserID,
		"userType": auth.UserType,
		"email":    auth.Email,
		"verified": auth.Verified,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	tokenString, err := token.SignedString([]byte(a.jwtSecret))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	// this is where we deal with session data

	return tokenString, nil
}

// VerifyEmail verifies the user's email using the token
func (u *authService) VerifyEmail(ctx context.Context, token string) error {
	// Check token in Redis
	redisKey := "verification:token:" + token
	authID, err := u.redis.Get(ctx, redisKey).Result()
	if err == redis.Nil {
		return errors.New("invalid or expired token")
	} else if err != nil {
		return err
	}

	// Update user status to active
	if err := u.repo.UpdateVerificationStatus(ctx, authID, true); err != nil {
		return err
	}

	// Delete token from Redis
	if err := u.redis.Del(ctx, redisKey).Err(); err != nil {
		return err
	}

	return nil
}

func (a *authService) GetAccount(ctx context.Context, authId string) (*domain.Auth, error) {
	auth, err := a.repo.GetAccount(ctx, authId)
	if err != nil {
		log.Printf("Failed to save a auth to database: %v", err)
		return nil, err
	}
	return auth, nil
}

func (a *authService) RefreshToken(ctx context.Context, authId string) (string, error) {

	auth, err := a.repo.GetAccount(ctx, authId)
	if err != nil {
		return "", fmt.Errorf("auth cannont be retrieved: ", err)
	}

	// Create JWT claims
	claims := jwt.MapClaims{
		"sub":      auth.AuthID,
		"userId":   auth.UserID,
		"userType": auth.UserType,
		"email":    auth.Email,
		"verified": auth.Verified,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	tokenString, err := token.SignedString([]byte(a.jwtSecret))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	// this is where we deal with session data

	return tokenString, nil
}

