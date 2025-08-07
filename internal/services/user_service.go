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
	if _, exists := models.Users[id]; exists {
		return ErrUserAlreadyExists
	}
	// handler'da fonksiyonun aldığı değerler kısmını model'den çekmek yerine
	profile := models.Profile{FirstName: firstName, LastName: lastName, AvatarURL: avatarURL}
	user := models.NewUser(id, username, email, passwordHash, role, profile)

	if !user.IsEmailValid() {
		return ErrInvalidEmail
	}

	models.Users[id] = user
	return nil
}

func (s *UserManager) ListUsers() {
	fmt.Println("all users:")
	for _, user := range models.Users {
		fmt.Println(user)
	}
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
	models.Users[id] = user

	return nil
}

func (s *UserManager) DeleteUser(id string) {
	if _, exists := models.Users[id]; exists {
		delete(models.Users, id)
		fmt.Println("user deleted:", id)
	} else {
		fmt.Println("user not found")
	}
}
