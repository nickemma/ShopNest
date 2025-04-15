package domain

import "time"

// User represents a customer in the system
type Customer struct {
	CustomerID  string      `json:"customerId"`
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	Phone       string      `json:"phone"`
	Address     Address     `json:"address"`
	Status      string      `json:"status"` // "active" | "inactive" | "suspended"
	Preferences Preferences `json:"preferences"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

// Address represents a user's address
type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

// Preferences represents user preferences
type Preferences struct {
	Currency string `json:"currency"`
	Language string `json:"language"`
}
