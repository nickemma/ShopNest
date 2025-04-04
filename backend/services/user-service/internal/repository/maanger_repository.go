package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/shopnest/user-service/internal/domain"
)

type ManagerRepository interface {
	RegisterManager(ctx context.Context, manager *domain.RegisterManager) error
	GetManagerByMail(ctx context.Context, email string) error
}

type PostgresManagerRepository struct {
	db *pgx.Conn
}


func NewPostgresManagerRepository(db *pgx.Conn) *PostgresManagerRepository {
	return &PostgresManagerRepository{db: db}
}


func (mgr *PostgresManagerRepository) RegisterManager(ctx context.Context, manager *domain.RegisterManager) error {
	return nil
}
func (mgr *PostgresManagerRepository) GetManagerByMail(ctx context.Context, email string) error {
	return nil
}