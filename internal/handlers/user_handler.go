package handlers

import (
	"devflow/internal/services"
	"fmt"
)

var userService = services.NewUserService()

func CreateUser(id, username, email, passwordHash, role, firstName, lastName, avatarURL string) {
	err := userService.CreateUser(id, username, email, passwordHash, role, firstName, lastName, avatarURL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user created")
}
func ListUsers() {
	userService.ListUsers()
}
func UpdateUser(id, newUsername, newEmail, newPasswordHash, newRole, newFirstName, newLastName, newAvatarURL string) {
	err := userService.UpdateUser(id, newUsername, newEmail, newPasswordHash, newRole, newFirstName, newLastName, newAvatarURL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user updated")
}
func ListUsersByRole(role string) {
	userService.FilterUsersByRole(role)
}
func DeleteUser(id string) {
	userService.DeleteUser(id)
}
