package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/shopnest/user-service/internal/domain"
	"time"
)

type AuthRepository interface {
	CreateAccount(ctx context.Context, auth *domain.Auth) error
	GetAccount(ctx context.Context, authId string) (*domain.Auth, error)
	GetAccountByEmail(ctx context.Context, email string) (*domain.Auth, error)
	UpdateVerificationStatus(ctx context.Context, authId string, status bool) error
}

// PostgresUserRepository implements UserRepository with PostgreSQL
type PostgresAuthRepository struct {
	db *pgx.Conn
}

// NewPostgresUserRepository creates a new repository instance
func NewPostgresAuthRepository(db *pgx.Conn) *PostgresAuthRepository {
	return &PostgresAuthRepository{db: db}
}

func (r *PostgresAuthRepository) CreateAccount(ctx context.Context, auth *domain.Auth) error {
	authQuery := `
        INSERT INTO auth (auth_id, user_id, user_type, email, password_hash, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)`
	// userId is initially set as empty, verifield is defaulted to false
	_, err := r.db.Exec(ctx, authQuery, auth.AuthID, "" /*auth.UserID*/, auth.UserType, auth.Email,
		auth.PasswordHash, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresAuthRepository) GetAccount(ctx context.Context, authId string) (*domain.Auth, error) {
	auth := &domain.Auth{}

	authQuery := `SELECT auth_id, user_id, user_type, email, password_hash, verified, created_at, updated_at
                  FROM auth WHERE auth_id = $1`
	err := r.db.QueryRow(ctx, authQuery, authId).Scan(&auth.AuthID, &auth.UserID, &auth.UserType, &auth.Email,
		&auth.PasswordHash, auth.Verified, &auth.CreatedAt, &auth.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
func (r *PostgresAuthRepository) GetAccountByEmail(ctx context.Context, email string) (*domain.Auth, error) {
	auth := &domain.Auth{}

	authQuery := `SELECT auth_id, user_id, user_type, email, password_hash, verified, created_at, updated_at
                  FROM auth WHERE email = $1`
	err := r.db.QueryRow(ctx, authQuery, email).Scan(&auth.AuthID, &auth.UserID, &auth.UserType, &auth.Email,
		&auth.PasswordHash, auth.Verified, &auth.CreatedAt, &auth.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (r *PostgresAuthRepository) UpdateVerificationStatus(ctx context.Context, authId string, status bool) error {
	query := `UPDATE auth SET verified = $1, updated_at = $2 WHERE user_id = $3`
	_, err := r.db.Exec(ctx, query, status, time.Now(), authId)
	return err
}
