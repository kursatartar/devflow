package interfaces

import "devflow/internal/models"

type UserService interface {
	CreateUser(id, username, email, passwordHash, role, firstName, lastName, avatarURL string) (*models.User, error)
	ListUsers() []*models.User
	UpdateUser(id string, newUsername, newEmail, newPasswordHash, newRole, newFirstName, newLastName, newAvatarURL *string) error
	DeleteUser(id string) error
	FilterUsersByRole(role string) []*models.User
	GetUser(id string) (*models.User, error)
	Register(username, email, password, firstName, lastName, avatarURL string) (*models.User, error)
	Authenticate(identifier, password string) (*models.User, error)
}
