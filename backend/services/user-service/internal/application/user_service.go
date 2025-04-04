package application

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/shopnest/user-service/config"
	"github.com/shopnest/user-service/internal/domain"
	"github.com/shopnest/user-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// UserService defines the interface for user-related operations
type UserService interface {
	Register(ctx context.Context, name, email, password string) (string, error)
	VerifyEmail(ctx context.Context, token string) error
	Login(ctx context.Context, email, password string) (string, error)
}

// userService implements UserService
type userService struct {
	repo       repository.UserRepository
	redis      *redis.Client
	rabbitMQ   *amqp091.Channel
	smtpConfig config.SMTPConfig
	jwtSecret  string
}

// NewUserService creates a new service instance
func NewUserService(
	repo repository.UserRepository,
	redis *redis.Client,
	rabbitMQ *amqp091.Channel,
	smtpConfig config.SMTPConfig,
	jwtSecret string,
) UserService {
	return &userService{
		repo:       repo,
		redis:      redis,
		rabbitMQ:   rabbitMQ,
		smtpConfig: smtpConfig,
		jwtSecret:  jwtSecret,
	}
}

// Register handles user registration
func (u *userService) Register(ctx context.Context, name, email, password string) (string, error) {
	// Generate UUIDs for user and auth
	userID := uuid.New().String()
	authID := uuid.New().String()

	// Hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Create user and auth models
	user := &domain.User{
		UserID:    userID,
		Name:      name,
		Email:     email,
		Status:    "inactive", // User starts as inactive until email is verified
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	auth := &domain.Auth{
		AuthID:       authID,
		UserID:       userID,
		UserType:     "customer",
		Email:        email,
		PasswordHash: string(passwordHash),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Save to database
	if err := u.repo.CreateUser(ctx, user, auth); err != nil {
		log.Printf("Failed to save a user to database: %v", err)
		return "", err
	}

	// Generate verification token
	token := uuid.New().String()

	// Store token in Redis with 1-hour expiration
	redisKey := "verification:token:" + token
	if err := u.redis.SetEx(ctx, redisKey, userID, 1*time.Hour).Err(); err != nil {
		log.Printf("Failed to store verification token: %v", err)
		return "", err
	}

	// Publish email task to RabbitMQ
	emailBody := []byte(`{"email": "` + email + `", "token": "` + token + `"}`)
	err = u.rabbitMQ.Publish("", "email_queue", false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        emailBody,
	})
	if err != nil {
		log.Printf("Failed to publish email: %v", err)
		return "", err
	}

	return userID, nil
}

// VerifyEmail verifies the user's email using the token
func (u *userService) VerifyEmail(ctx context.Context, token string) error {
	// Check token in Redis
	redisKey := "verification:token:" + token
	userID, err := u.redis.Get(ctx, redisKey).Result()
	if err == redis.Nil {
		return errors.New("invalid or expired token")
	} else if err != nil {
		return err
	}

	// Update user status to active
	if err := u.repo.UpdateUserStatus(ctx, userID, "active"); err != nil {
		return err
	}

	// Delete token from Redis
	if err := u.redis.Del(ctx, redisKey).Err(); err != nil {
		return err
	}

	return nil
}

// Login handles user login
func (u *userService) Login(ctx context.Context, email, password string) (string, error) {
	// Get user and auth data
	user, auth, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Check if email is verified (status must be "active")
	if user.Status != "active" {
		return "", errors.New("email not verified")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(auth.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Create JWT claims
	claims := jwt.MapClaims{
		"sub":   user.UserID,
		"email": user.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}
