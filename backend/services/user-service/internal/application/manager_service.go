package application

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"github.com/shopnest/user-service/config"
	"github.com/shopnest/user-service/internal/domain"
	"github.com/shopnest/user-service/internal/repository"
)

type ManagerService interface {
	Register(ctx context.Context, authId, name, email string) (string, error)
	ApproveRegistration(ctx context.Context, authId string) (string, error)
	GetManagerProfile(ctx context.Context, managerId string) (*domain.Manager, error)
}

type managerService struct {
	repo       repository.ManagerRepository
	authRepo   repository.AuthRepository
	rabbitMQ   *amqp091.Channel
	smtpConfig config.SMTPConfig
}

func NewManagerService(
	repo repository.ManagerRepository,
	authRepo repository.AuthRepository,

) ManagerService {
	return &managerService{
		repo:     repo,
		authRepo: authRepo,
	}
}

// manager signs up and waits for the super admin approval.
func (mgr *managerService) Register(ctx context.Context, authId, name, email string) (string, error) {
	// Generate UUIDs for user and auth
	managerID := uuid.New().String()

	// This might need to make a call to the authDB to verifiy
	// the authID+email. This would mean the email is not to be passed
	// explicitly but obtained from the access token and compared against
	// authDB entries
	auth, err := mgr.authRepo.GetAccount(ctx, authId)
	if err != nil {
		return "", err
	}

	if auth.Email != email {
		return "", errors.New("account email does not match")
	}

	// redundant as this is already checked on login
	// user needs to be logged in before registering
	if !auth.Verified {
		return "", errors.New("email not verified")
	}

	// Create user and auth models
	manager := &domain.RegisterManager{
		ManagerID: managerID,
		Name:      name,
		Email:     email,
		Role:      "ADMIN",
		Approved:  false, // set to false, enforced in lower layers (infra.)
	}

	// Save to database
	if err := mgr.repo.RegisterManager(ctx, authId, manager); err != nil {
		log.Printf("Failed to save a manager to database: %v", err)
		return "", err
	}

	return managerID, nil
}

// Callable only by super-admin
func (mgr *managerService) ApproveRegistration(ctx context.Context, authId string) (string, error) {

	auth, err := mgr.authRepo.GetAccount(ctx, authId)
	if err != nil {
		return "", err
	}

	if !auth.Verified {
		return "", errors.New("cannot approve: manager yet to verify")
	}
	if err := mgr.repo.ApproveManager(ctx, auth.UserID); err != nil {
		return "", err
	}

	return fmt.Sprintf("manager with email address %+s has been approved\n", auth.Email), nil

}

func (mgr *managerService) GetManagerProfile(ctx context.Context, managerId string) (*domain.Manager, error) {
	manager, err := mgr.repo.GetManager(ctx, managerId)
	if err != nil {
		log.Printf("Failed to retrieve manager profile: %v", err)
		return nil, err
	}
	return manager, nil
}
