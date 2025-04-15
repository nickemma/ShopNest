package application

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/shopnest/user-service/internal/domain"
	"github.com/shopnest/user-service/internal/repository"
)

// UserService defines the interface for user-related operations
type UserService interface {
	Register(ctx context.Context, authId, name, email string) (string, error)
	ActivateUser(ctx context.Context, authId string) (string, error)
	GetUserProfile(ctx context.Context, userId string) (*domain.User, error)
}

// userService implements UserService
type userService struct {
	repo     repository.UserRepository
	authRepo repository.AuthRepository
}

// NewUserService creates a new service instance
func NewUserService(
	repo repository.UserRepository,
	authRepo repository.AuthRepository,
) UserService {
	return &userService{
		repo:     repo,
		authRepo: authRepo,
	}
}

// Register handles user registration
func (u *userService) Register(ctx context.Context, authId, name, email string) (string, error) {
	// Generate UUIDs for user and auth
	userID := uuid.New().String()

	auth, err := u.authRepo.GetAccount(ctx, authId)
	if err != nil {
		return "", err
	}

	if auth.Email != email {
		return "", errors.New("account email does not match")
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

	// Save to database
	if err := u.repo.CreateUser(ctx, authId, user); err != nil {
		log.Printf("Failed to save a user to database: %v", err)
		return "", err
	}
	return userID, nil
}

func (u *userService) ActivateUser(ctx context.Context, authId string) (string, error) {

	auth, err := u.authRepo.GetAccount(ctx, authId)
	if err != nil {
		return "", err
	}

	if !auth.Verified {
		return "", errors.New("cannot approve: customer yet to verify email")
	}
	if err := u.repo.UpdateUserStatus(ctx, auth.UserID, "active"); err != nil {
		return "", err
	}

	return fmt.Sprintf("customer with email address %+s has been approved\n", auth.Email), nil

}

func (u *userService) GetUserProfile(ctx context.Context, userId string) (*domain.User, error) {
	user, err := u.repo.GetUser(ctx, userId)
	if err != nil {
		log.Printf("Failed to retrieve customer profile: %v", err)
		return nil, err
	}
	return user, nil
}
