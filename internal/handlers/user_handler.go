package handlers

import (
	"devflow/internal/models"
	"devflow/internal/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var userService = services.NewUserService()

type createUserReq struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	AvatarURL    string `json:"avatar_url"`
}

func CreateUser(c *fiber.Ctx) error {
	var body createUserReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "invalid json",
		})
	}

	for _, u := range models.Users {
		if u.Email == body.Email {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"success": false,
				"message": "email already exists",
			})
		}
	}
	
	if err := userService.CreateUser(
		"",
		body.Username,
		body.Email,
		body.PasswordHash,
		body.Role,
		body.FirstName,
		body.LastName,
		body.AvatarURL,
	); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	createdUser := fiber.Map{
		"username":  body.Username,
		"email":     body.Email,
		"role":      body.Role,
		"firstName": body.FirstName,
		"lastName":  body.LastName,
		"avatarURL": body.AvatarURL,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
		"data":    createdUser,
	})
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
