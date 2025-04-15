package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/shopnest/user-service/internal/domain"
)

type ManagerRepository interface {
	RegisterManager(ctx context.Context, authId string, manager *domain.RegisterManager) error
	GetManagerByEmail(ctx context.Context, email string) (*domain.Manager, error)
	GetManager(ctx context.Context, managerId string) (*domain.Manager, error)
	ApproveManager(ctx context.Context, email string) error
}

type PostgresManagerRepository struct {
	db *pgx.Conn
}

func NewPostgresManagerRepository(db *pgx.Conn) *PostgresManagerRepository {
	return &PostgresManagerRepository{db: db}
}

func (mgr *PostgresManagerRepository) RegisterManager(ctx context.Context, authId string, manager *domain.RegisterManager) error {
	tx, err := mgr.db.Begin(ctx)
	if err != nil {
		return err
	}
	// rolling back our transaction in case of failed transactions or error
	defer tx.Rollback(ctx)

	mgrQuery := `
		INSERT INTO managers (manager_id, name, email, role, approved)
		VALUES ($1, $2, $3, $4, $5)`

	_, err = tx.Exec(ctx, mgrQuery, manager.ManagerID, manager.Name, manager.Email, manager.Role /*manager.Approved*/, false)
	if err != nil {
		return err
	}

	// Update auth to only change the user_id
	authQuery := `
		UPDATE auth
		SET user_id = $1,
			updated_at = $2
		WHERE auth_id = $3`

	_, err = tx.Exec(ctx, authQuery, manager.ManagerID, time.Now(), authId)
	if err != nil {
		return err
	}

	// Commiting the transaction
	return tx.Commit(ctx)
}

func (mgr *PostgresManagerRepository) GetManager(ctx context.Context, managerId string) (*domain.Manager, error) {

	manager := &domain.Manager{}

	// Query manager
	managerQuery := `SELECT manager_id, name, email, role, approved, created_at, updated_at
					FROM managers WHERE manager_id = $1`
	err := mgr.db.QueryRow(ctx, managerQuery, managerId).Scan(
		&manager.ManagerID, &manager.Name, &manager.Email,
		&manager.Role, &manager.Approved, &manager.CreatedAt, &manager.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return manager, nil
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

func (mgr *PostgresManagerRepository) ApproveManager(ctx context.Context, managerID string) error {
	query := `UPDATE managers SET approved = $1, updated_at = $2 WHERE manager_id = $3`
	_, err := mgr.db.Exec(ctx, query, "TRUE", time.Now(), managerID)
	return err

}
