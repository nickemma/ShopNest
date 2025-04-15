package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/shopnest/user-service/internal/domain"
)

// UserRepository defines the interface for user data access
type CustomerRepository interface {
	CreateCustomer(ctx context.Context, authId string, customer *domain.Customer) error
	GetCustomerByEmail(ctx context.Context, email string) (*domain.Customer, error)
	GetCustomer(ctx context.Context, customerId string) (*domain.Customer, error)
	UpdateCustomerStatus(ctx context.Context, customerId, status string) error
}

// PostgresUserRepository implements UserRepository with PostgreSQL
type PostgresCustomerRepository struct {
	db *pgx.Conn
}

// NewPostgresUserRepository creates a new repository instance
func NewPostgresCustomerRepository(db *pgx.Conn) *PostgresCustomerRepository {
	return &PostgresCustomerRepository{db: db}
}

// CreateUser inserts a new user and auth record into the database
func (r *PostgresCustomerRepository) CreateCustomer(ctx context.Context, authId string, customer *domain.Customer) error {
	// Beginning the sql transaction and commiting it to the database
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	// rolling back our transaction in case of failed transactions or error
	defer tx.Rollback(ctx)

	// Insert the user to database
	userQuery := `
        INSERT INTO customers (customer_id, name, email, phone, address, status, preferences, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err = tx.Exec(ctx, userQuery, customer.CustomerID, customer.Name, customer.Email, customer.Phone,
		customer.Address, "inactive", customer.Preferences, time.Now(), time.Now())
	if err != nil {
		return err
	}

	// Update auth to only change the user_id
	authQuery := `
		UPDATE auth
		SET user_id = $1,
			updated_at = $2
		WHERE auth_id = $3`

	_, err = tx.Exec(ctx, authQuery, customer.CustomerID, time.Now(), authId)
	if err != nil {
		return err
	}

	// Commiting the transaction
	return tx.Commit(ctx)
}

func (r *PostgresCustomerRepository) GetCustomer(ctx context.Context, customerId string) (*domain.Customer, error) {
	customer := &domain.Customer{}

	// Query user
	userQuery := `SELECT customer_id, name, email, phone, address, status, preferences, created_at, updated_at
                  FROM customers WHERE customer_id = $1`
	err := r.db.QueryRow(ctx, userQuery, customerId).Scan(&customer.CustomerID, &customer.Name, &customer.Email, &customer.Phone,
		&customer.Address, &customer.Status, &customer.Preferences, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// GetUserByEmail retrieves a user and their auth data by email
func (r *PostgresCustomerRepository) GetCustomerByEmail(ctx context.Context, email string) (*domain.Customer, error) {
	customer := &domain.Customer{}

	// Query user
	userQuery := `SELECT customer_id, name, email, phone, address, status, preferences, created_at, updated_at
                  FROM customers WHERE email = $1`
	err := r.db.QueryRow(ctx, userQuery, email).Scan(&customer.CustomerID, &customer.Name, &customer.Email, &customer.Phone,
		&customer.Address, &customer.Status, &customer.Preferences, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// UpdateUserStatus updates the user's status (e.g., to "active" after verification)
func (r *PostgresCustomerRepository) UpdateCustomerStatus(ctx context.Context, userID, status string) error {
	query := `UPDATE customers SET status = $1, updated_at = $2 WHERE customer_id = $3`
	_, err := r.db.Exec(ctx, query, status, time.Now(), userID)
	return err
}
