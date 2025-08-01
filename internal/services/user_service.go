package services

import (
	"devflow/internal/models"
	"fmt"
)

type UserService interface {
	CreateUser(id, username, email, passwordHash, role string, profile models.Profile)
	ListUsers()
	UpdateUser(id, newUsername, newEmail, newPasswordHash, newRole string, newProfile models.Profile)
	DeleteUser(id string)
	FilterUsersByRole(role string) []models.User
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) CreateUser(id, username, email, passwordHash, role string, profile models.Profile) {
	user := models.NewUser(id, username, email, passwordHash, role, profile)
	models.Users[id] = user
	fmt.Println("user created:", user)
}

func (s *userService) ListUsers() {
	fmt.Println("all users:")
	for id, user := range models.Users {
		fmt.Printf("- %s: %s %s (%s)\n", id, user.Profile.FirstName, user.Profile.LastName, user.Email)
	}
}

func (s *userService) FilterUsersByRole(role string) []models.User {
	var filtered []models.User
	for _, u := range models.Users {
		if u.Role == role {
			filtered = append(filtered, u)
		}
	}
	return filtered
}

func (s *userService) UpdateUser(id, newUsername, newEmail, newPasswordHash, newRole string, newProfile models.Profile) {
	if user, exists := models.Users[id]; exists {
		user.Username = newUsername
		user.Email = newEmail
		user.PasswordHash = newPasswordHash
		user.Role = newRole
		user.Profile = newProfile
		user.UpdatedAt = user.UpdatedAt.UTC()
		models.Users[id] = user
		fmt.Println("user updated:", user)
	} else {
		fmt.Println("user not found")
	}
}

func (s *userService) DeleteUser(id string) {
	if _, exists := models.Users[id]; exists {
		delete(models.Users, id)
		fmt.Println("user deleted:", id)
	} else {
		fmt.Println("user not found")
	}

}
