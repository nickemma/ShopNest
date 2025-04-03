package domain

import "time"

// Manager represents an admin/support in the system
type Manager struct {
	ManagerID string    `json:"managerId"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	roles     string    `json:"roles"` // "Admin" | "Support"
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
