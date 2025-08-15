package handlers

import (
	"devflow/internal/services"
	"errors"
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

	if err := userService.CreateUser(
		"", body.Username, body.Email, body.PasswordHash, body.Role, body.FirstName, body.LastName, body.AvatarURL,
	); err != nil {
		status := fiber.StatusInternalServerError
		if errors.Is(err, services.ErrEmailExists) {
			status = fiber.StatusConflict
		} else if errors.Is(err, services.ErrInvalidEmail) {
			status = fiber.StatusBadRequest
		}
		return c.Status(status).JSON(fiber.Map{"success": false, "message": err.Error()})
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

func ListUsers(c *fiber.Ctx) error {
	users := userService.ListUsers()

	list := make([]fiber.Map, 0, len(users))
	for _, u := range users {
		list = append(list, fiber.Map{
			"username":  u.Username,
			"email":     u.Email,
			"role":      u.Role,
			"firstName": u.Profile.FirstName,
			"lastName":  u.Profile.LastName,
			"avatarURL": u.Profile.AvatarURL,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Users fetched successfully",
		"data":    list,
	})
}

type updateUserReq struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	AvatarURL    string `json:"avatar_url"`
}

func UpdateUser(u *fiber.Ctx) error {
	id := u.Params("id")

	var body updateUserReq
	if err := u.BodyParser(&body); err != nil {
		return u.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "invalid json",
		})
	}

	if err := userService.UpdateUser(
		id,
		body.Username,
		body.Email,
		body.PasswordHash,
		body.Role,
		body.FirstName,
		body.LastName,
		body.AvatarURL,
	); err != nil {
		status := fiber.StatusInternalServerError
		switch err {
		case services.ErrUserNotFound:
			status = fiber.StatusNotFound
		case services.ErrInvalidEmail:
			status = fiber.StatusBadRequest
		case services.ErrEmailExists:
			status = fiber.StatusConflict
		}
		return u.Status(status).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	updated := fiber.Map{
		"username":  body.Username,
		"email":     body.Email,
		"role":      body.Role,
		"firstName": body.FirstName,
		"lastName":  body.LastName,
		"avatarURL": body.AvatarURL,
	}

	return u.JSON(fiber.Map{
		"success": true,
		"message": "User updated successfully",
		"data":    updated,
	})
}

func ListUsersByRole(c *fiber.Ctx) error {
	role := c.Query("role")
	users := userService.FilterUsersByRole(role)

	list := make([]fiber.Map, 0, len(users))
	for _, u := range users {
		list = append(list, fiber.Map{
			"username":  u.Username,
			"email":     u.Email,
			"role":      u.Role,
			"firstName": u.Profile.FirstName,
			"lastName":  u.Profile.LastName,
			"avatarURL": u.Profile.AvatarURL,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Users fetched successfully",
		"data":    list,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := userService.DeleteUser(id); err != nil {
		status := fiber.StatusInternalServerError
		if errors.Is(err, services.ErrUserNotFound) {
			status = fiber.StatusNotFound
		}
		return c.Status(status).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "User deleted successfully",
	})
}
