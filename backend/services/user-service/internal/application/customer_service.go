package application

import (
	"context"
	"errors"
	// "fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/shopnest/user-service/internal/domain"
	"github.com/shopnest/user-service/internal/repository"
)

// UserService defines the interface for user-related operations
type CustomerService interface {
	RegisterCustomer(ctx context.Context, authId, name, email string) (string, error)
	// ActivateCustomer(ctx context.Context, authId string) (string, error)
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
// This is created separately from update user because it needs to reissue token to reflect the userId
// Update customer does not need that. This prevents the log in/out flow for updating token 
// Instead of returning token here, the best approach is to have Frontend call the refresh token endpoint
// after creating a customer. This is clear separation of concern but with increased latency.
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
		// status field not needed. The email verified field on auth is enough
		// setting to active by default (why: changing/refactoring of code to remove
		// status is tiring)
		Status:    "active", // enforced in db
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

// Why are we activating a customer ? No need
// Once they verified their email; thats enough activation. This is an order system.
// func (u *customerService) ActivateCustomer(ctx context.Context, authId string) (string, error) {

// 	auth, err := u.authRepo.GetAccount(ctx, authId)
// 	if err != nil {
// 		return "", err
// 	}

// 	if !auth.Verified {
// 		return "", errors.New("cannot approve: customer yet to verify email")
// 	}
// 	if err := u.repo.UpdateCustomerStatus(ctx, auth.UserID, "active"); err != nil {
// 		return "", err
// 	}

// 	return fmt.Sprintf("customer with email address %+s has been approved\n", auth.Email), nil
// }

func (u *customerService) GetCustomerProfile(ctx context.Context, customerId string) (*domain.Customer, error) {
	customer, err := u.repo.GetCustomer(ctx, customerId)
	if err != nil {
		log.Printf("Failed to retrieve customer profile: %v", err)
		return nil, err
	}
	return customer, nil
}
