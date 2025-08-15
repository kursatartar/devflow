package services

import (
	"devflow/internal/interfaces"
	"devflow/internal/models"
	"fmt"
	"time"
)

type UserManager struct{}

func NewUserService() interfaces.UserService {
	return &UserManager{}
}

func (s *UserManager) CreateUser(id, username, email, passwordHash, role, firstName, lastName, avatarURL string) error {
	for _, u := range models.Users {
		if u.Email == email {
			return ErrEmailExists
		}
	}
	if _, exists := models.Users[id]; exists {
		return ErrUserAlreadyExists
	}

	profile := models.Profile{FirstName: firstName, LastName: lastName, AvatarURL: avatarURL}
	user, err := models.NewUser(id, username, email, passwordHash, role, profile)
	if err != nil {
		return fmt.Errorf("kullanıcı oluşturulamadı: %w", err)
	}

	isValid, err := user.IsEmailValid()
	if err != nil {
		return fmt.Errorf("email validasyon hatası: %w", err)
	}
	if !isValid {
		return ErrInvalidEmail
	}

	models.Users[id] = user
	return nil
}

func (s *UserManager) ListUsers() []*models.User {
	out := make([]*models.User, 0, len(models.Users))
	for _, u := range models.Users {
		out = append(out, u)
	}
	return out
}

func (s *UserManager) FilterUsersByRole(role string) []*models.User {
	var filtered []*models.User
	for _, u := range models.Users {
		if u.Role == role {
			filtered = append(filtered, u)
		}
	}
	return filtered
}

func (s *UserManager) UpdateUser(id, newUsername, newEmail, newPasswordHash, newRole, newFirstName, newLastName, newAvatarURL string) error {
	user, exists := models.Users[id]
	if !exists {
		return ErrUserNotFound
	}

	for _, u := range models.Users {
		if u.Email == newEmail && u.ID != id {
			return ErrEmailExists
		}
	}

	tempUser := *user
	tempUser.Email = newEmail
	isValid, err := tempUser.IsEmailValid()
	if err != nil {
		return err
	}
	if !isValid {
		return ErrInvalidEmail
	}

	user.Username = newUsername
	user.Email = newEmail
	user.PasswordHash = newPasswordHash
	user.Role = newRole
	user.Profile = models.Profile{
		FirstName: newFirstName,
		LastName:  newLastName,
		AvatarURL: newAvatarURL,
	}
	user.UpdatedAt = time.Now()

	return nil
}

func (s *UserManager) DeleteUser(id string) error {
	if _, exists := models.Users[id]; exists {
		delete(models.Users, id)
		return nil
	}
	return ErrUserNotFound
}
