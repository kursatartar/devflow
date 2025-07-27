package handlers

import (
	"fmt"
	"github.com/kursatartar/devflowv2/models"
)

const (
	StatusPending = "pending"
	StatusActive  = "active"
	StatusDone    = "done"
)

func CreateUser(id, username, email, passwordHash, role string, profile models.Profile) {
	user := models.NewUser(id, username, email, passwordHash, role, profile)
	models.Users[id] = user
	fmt.Println("user created:", user)
}

func ListUsers() {
	fmt.Println("all users:")
	for id, user := range models.Users {
		fmt.Printf("- %s: %s %s (%s)\n", id, user.Profile.FirstName, user.Profile.LastName, user.Email)
	}
}

func UpdateUser(id, newUsername, newEmail, newPasswordHash, newRole string, newProfile models.Profile) {
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

func DeleteUser(id string) {
	if _, exists := models.Users[id]; exists {
		delete(models.Users, id)
		fmt.Println("user deleted:", id)
	} else {
		fmt.Println("user not found")
	}
}
