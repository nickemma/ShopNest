package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/shopnest/user-service/internal/domain"
)

// UserRepository defines the interface for user data access
type UserRepository interface {
	CreateUser(ctx context.Context, authId string, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUser(ctx context.Context, userId string) (*domain.User, error)
	UpdateUserStatus(ctx context.Context, userID, status string) error
}

// PostgresUserRepository implements UserRepository with PostgreSQL
type PostgresUserRepository struct {
	db *pgx.Conn
}

// NewPostgresUserRepository creates a new repository instance
func NewPostgresCustomerRepository(db *pgx.Conn) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

// CreateUser inserts a new user and auth record into the database
func (r *PostgresUserRepository) CreateUser(ctx context.Context, authId string, user *domain.User) error {
	// Beginning the sql transaction and commiting it to the database
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	// rolling back our transaction in case of failed transactions or error
	defer tx.Rollback(ctx)

	// Insert the user to database
	userQuery := `
        INSERT INTO users (user_id, name, email, phone, address, status, preferences, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err = tx.Exec(ctx, userQuery, user.UserID, user.Name, user.Email, user.Phone,
		user.Address, "inactive", user.Preferences, time.Now(), time.Now())
	if err != nil {
		return err
	}

	// Update auth to only change the user_id
	authQuery := `
		UPDATE auth
		SET user_id = $1,
			updated_at = $2
		WHERE auth_id = $3`

	_, err = tx.Exec(ctx, authQuery, user.UserID, time.Now(), authId)
	if err != nil {
		return err
	}

	// Commiting the transaction
	return tx.Commit(ctx)
}

func (r *PostgresUserRepository) GetUser(ctx context.Context, userId string) (*domain.User, error) {
	user := &domain.User{}

	// Query user
	userQuery := `SELECT user_id, name, email, phone, address, status, preferences, created_at, updated_at
                  FROM users WHERE user_id = $1`
	err := r.db.QueryRow(ctx, userQuery, userId).Scan(&user.UserID, &user.Name, &user.Email, &user.Phone,
		&user.Address, &user.Status, &user.Preferences, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByEmail retrieves a user and their auth data by email
func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}

	// Query user
	userQuery := `SELECT user_id, name, email, phone, address, status, preferences, created_at, updated_at
                  FROM users WHERE email = $1`
	err := r.db.QueryRow(ctx, userQuery, email).Scan(&user.UserID, &user.Name, &user.Email, &user.Phone,
		&user.Address, &user.Status, &user.Preferences, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUserStatus updates the user's status (e.g., to "active" after verification)
func (r *PostgresUserRepository) UpdateUserStatus(ctx context.Context, userID, status string) error {
	query := `UPDATE users SET status = $1, updated_at = $2 WHERE user_id = $3`
	_, err := r.db.Exec(ctx, query, status, time.Now(), userID)
	return err
}
