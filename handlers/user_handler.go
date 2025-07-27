package handlers

import (
	"fmt"
	"github.com/kursatartar/devflowv2/models"
)

func CreateUser(id, name, email string) {
	user := models.User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	models.Users[id] = user
	fmt.Println("user created:", user)
}

func ListUsers() {
	fmt.Println("all users:")
	for id, user := range models.Users {
		fmt.Printf("- id: %s, name: %s, e-mail: %s\n", id, user.Name, user.Email)
	}
}

func UpdateUser(id, newName, newEmail string) {
	if _, exists := models.Users[id]; exists {
		user := models.Users[id]
		user.Name = newName
		user.Email = newEmail
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
