package domain

import "time"

// Manager represents an admin/support in the system
type Manager struct {
	RegisterManager
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}


type RegisterManager struct {
	ManagerID string    `json:"managerId"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role     string    `json:"role"` // "Admin" | "Support"
}