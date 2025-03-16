package models

// User represents a user in the system.
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewUser creates a new User.
func NewUser(id, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}

// Users is a collection of User objects.
type Users []*User

// NewUsers creates a new Users collection.
func NewUsers(users ...*User) Users {
	return users
}
