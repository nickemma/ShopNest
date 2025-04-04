package application

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"github.com/shopnest/user-service/config"
	"github.com/shopnest/user-service/internal/domain"
	"github.com/shopnest/user-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type ManagerService interface {
	Register(ctx context.Context, name, email, password string) (string, error)
	ApproveRegistration(ctx context.Context, email string) (string, error)
}

type managerService struct {
	repo       repository.ManagerRepository
	rabbitMQ   *amqp091.Channel
	smtpConfig config.SMTPConfig
}

func NewManagerService(
	repo repository.ManagerRepository,
	rabbitMQ *amqp091.Channel,
	smtpConfig config.SMTPConfig,

) ManagerService {
	return &managerService{
		repo:       repo,
		rabbitMQ:   rabbitMQ,
		smtpConfig: smtpConfig,
	}
}

// manager signs up and waits for the super admin approval.
func (mgr *managerService) Register(ctx context.Context, name, email, password string) (string, error) {
	// Generate UUIDs for user and auth
	managerID := uuid.New().String()
	authID := uuid.New().String()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Create user and auth models
	manager := &domain.RegisterManager{
		ManagerID: managerID,
		Name:      name,
		Email:     email,
		Role:      "ADMIN",
		Approved:  false, // set to false, enforced in lower layers (infra.)
	}
	auth := &domain.Auth{
		AuthID:       authID,
		UserID:       managerID,
		UserType:     "MANAGER",
		Email:        email,
		PasswordHash: string(passwordHash),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Save to database
	if err := mgr.repo.RegisterManager(ctx, manager, auth); err != nil {
		log.Printf("Failed to save a manager to database: %v", err)
		return "", err
	}

	// publish the manager email and name to the super admin/permissioned manager
	emailBody := []byte(`{"email": "` + email + `", "name": "` + manager.Name + `"}`)
	// NOTE: change this "email_queue" to manager_approval_queue
	err = mgr.rabbitMQ.Publish("", "email_queue", false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        emailBody,
	})
	if err != nil {
		log.Printf("Failed to publish email: %v", err)
		return "", err
	}

	return managerID, nil
}

// Callable only by super-admin
func (mgr *managerService) ApproveRegistration(ctx context.Context, email string) (string, error) {
	if err := mgr.repo.ApproveManager(ctx, email); err != nil {
		return "", err
	}

	return fmt.Sprintf("manager with email address %+s has been approved\n", email), nil

}
