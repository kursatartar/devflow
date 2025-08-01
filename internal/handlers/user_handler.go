package handlers

import (
	"devflow/internal/models"
	"devflow/internal/services"
)

const (
	StatusPending = "pending"
	StatusActive  = "active"
	StatusDone    = "done"
)

var userService = services.NewUserService()

func CreateUser(id, username, email, passwordHash, role string, profile models.Profile) {
	userService.CreateUser(id, username, email, passwordHash, role, profile)
}

func ListUsers() {
	userService.ListUsers()
}

func UpdateUser(id, newUsername, newEmail, newPasswordHash, newRole string, newProfile models.Profile) {
	userService.UpdateUser(id, newUsername, newEmail, newPasswordHash, newRole, newProfile)
}

func DeleteUser(id string) {
	userService.DeleteUser(id)
}
