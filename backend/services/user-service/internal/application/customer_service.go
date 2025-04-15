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
type CustomerService interface {
	RegisterCustomer(ctx context.Context, authId, name, email string) (string, error)
	ActivateCustomer(ctx context.Context, authId string) (string, error)
	GetCustomerProfile(ctx context.Context, customerId string) (*domain.Customer, error)
}

// userService implements UserService
type customerService struct {
	repo     repository.CustomerRepository
	authRepo repository.AuthRepository
}

// NewUserService creates a new service instance
func NewCustomerService(
	repo repository.CustomerRepository,
	authRepo repository.AuthRepository,
) CustomerService {
	return &customerService{
		repo:     repo,
		authRepo: authRepo,
	}
}

// Register handles user registration
func (u *customerService) RegisterCustomer(ctx context.Context, authId, name, email string) (string, error) {
	// Generate UUIDs for user and auth
	customerID := uuid.New().String()

	auth, err := u.authRepo.GetAccount(ctx, authId)
	if err != nil {
		return "", err
	}

	if auth.Email != email {
		return "", errors.New("account email does not match")
	}

	// Create customer and auth models
	customer := &domain.Customer{
		CustomerID:    customerID,
		Name:      name,
		Email:     email,
		Status:    "inactive", // customer starts as inactive until email is verified
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save to database
	if err := u.repo.CreateCustomer(ctx, authId, customer); err != nil {
		log.Printf("Failed to save a customer to database: %v", err)
		return "", err
	}
	return customerID, nil
}

func (u *customerService) ActivateCustomer(ctx context.Context, authId string) (string, error) {

	auth, err := u.authRepo.GetAccount(ctx, authId)
	if err != nil {
		return "", err
	}

	if !auth.Verified {
		return "", errors.New("cannot approve: customer yet to verify email")
	}
	if err := u.repo.UpdateCustomerStatus(ctx, auth.UserID, "active"); err != nil {
		return "", err
	}

	return fmt.Sprintf("customer with email address %+s has been approved\n", auth.Email), nil

}

func (u *customerService) GetCustomerProfile(ctx context.Context, customerId string) (*domain.Customer, error) {
	customer, err := u.repo.GetCustomer(ctx, customerId)
	if err != nil {
		log.Printf("Failed to retrieve customer profile: %v", err)
		return nil, err
	}
	return customer, nil
}
