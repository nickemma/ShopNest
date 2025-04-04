package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/shopnest/user-service/internal/domain"
)

type ManagerRepository interface {
	RegisterManager(ctx context.Context, manager *domain.RegisterManager, auth *domain.Auth) error
	GetManagerByEmail(ctx context.Context, email string) error
	ApproveManager(ctx context.Context, email string) error
}

type PostgresManagerRepository struct {
	db *pgx.Conn
}

func NewPostgresManagerRepository(db *pgx.Conn) *PostgresManagerRepository {
	return &PostgresManagerRepository{db: db}
}

func (mgr *PostgresManagerRepository) RegisterManager(ctx context.Context, manager *domain.RegisterManager, auth *domain.Auth) error {
	tx, err := mgr.db.Begin(ctx)
	if err != nil {
		return err
	}
	// rolling back our transaction in case of failed transactions or error
	defer tx.Rollback(ctx)

	// Insert auth
	authQuery := `
        INSERT INTO auth (auth_id, user_id, user_type, email, password_hash, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = tx.Exec(ctx, authQuery, auth.AuthID, auth.UserID /*auth.UserType*/, "MANAGER", auth.Email,
		auth.PasswordHash, time.Now(), time.Now())
	if err != nil {
		return err
	}

	mgrQuery := `
		INSERT INTO managers (manager_id, name, email, role, approved)
		VALUES ($1, $2, $3, $4, $5)`

	_, err = tx.Exec(ctx, mgrQuery, manager.ManagerID, manager.Name, manager.Email, manager.Role, /*manager.Approved*/false)
	if err != nil {
		return err
	}
	// Commiting the transaction
	return tx.Commit(ctx)
}


func (mgr *PostgresManagerRepository) GetManagerByEmail(ctx context.Context, email string) (*domain.Manager, error) {

	manager := &domain.Manager{}

	// Query manager
	managerQuery := `SELECT manager_id, name, email, role, approved, created_at, updated_at
					FROM managers WHERE email = $1`
	err := mgr.db.QueryRow(ctx, managerQuery, email).Scan(
		&manager.ManagerID, &manager.Name, &manager.Email,
		&manager.Role, &manager.Approved, &manager.CreatedAt, &manager.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return manager, nil
}


func (mgr *PostgresManagerRepository) ApproveManager(ctx context.Context, email string) error {
	query := `UPDATE managers SET approved = $1, updated_at = $2 WHERE email = $3`
	_, err := mgr.db.Exec(ctx, query, "TRUE", time.Now(), email)
	return err

}