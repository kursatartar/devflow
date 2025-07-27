package handlers

import (
	"fmt"
	"time"

	"github.com/kursatartar/devflowv2/models"
)

func CreateUser(id, username, email, passwordHash, role, firstName, lastName, avatarURL string) {
	user := models.NewUser(id, username, email, passwordHash, role, firstName, lastName, avatarURL)
	models.Users[id] = user
	fmt.Println("user created:", user)
}

func ListUsers() {
	fmt.Println("all users:")
	for id, user := range models.Users {
		fmt.Printf("- id: %s, name: %s, e-mail: %s\n", id, user.Username, user.Email)
	}
}

func UpdateUser(id, newUsername, newEmail, newPasswordHash, newRole, newFirstName, newLastName, newAvatarURL string) {
	if user, exists := models.Users[id]; exists {
		user.Username = newUsername
		user.Email = newEmail
		user.PasswordHash = newPasswordHash
		user.Role = newRole
		user.Profile.FirstName = newFirstName
		user.Profile.LastName = newLastName
		user.Profile.AvatarURL = newAvatarURL
		user.UpdatedAt = time.Now().Format(time.RFC3339)

		models.Users[id] = user
		fmt.Println("user updated:", user.Username)
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
