package domain

import "time"

// Auth represents authentication data
type Auth struct {
	AuthID       string      `json:"authId"`
	UserID       string      `json:"userId"`
	UserType     string      `json:"userType"` // "customer" | "manager"
	Email        string      `json:"email"`
	PasswordHash string      `json:"passwordHash"`
	SessionData  SessionData `json:"sessionData"`
	CreatedAt    time.Time   `json:"createdAt"`
	UpdatedAt    time.Time   `json:"updatedAt"`
}

// SessionData tracks session-related information
type SessionData struct {
	LastLogin     time.Time `json:"lastLogin"`
	CurrentToken  string    `json:"currentToken"`
	TokenExpiry   time.Time `json:"tokenExpiry"`
	LoginAttempts int       `json:"loginAttempts"`
}
