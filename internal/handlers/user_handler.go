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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "invalid json"})
	}

	if body.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "username boş olamaz"})
	}
	if body.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "email boş olamaz"})
	}
	if body.PasswordHash == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "password_hash boş olamaz"})
	}
	if body.Role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "role boş olamaz"})
	}
	if body.FirstName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "firstname boş olama"})
	}
	if body.LastName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "lastname boş olamaz"})
	}
	if body.AvatarURL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "avatarurl boş olamaz"})
	}

	u, err := userService.CreateUser(
		"",
		body.Username,
		body.Email,
		body.PasswordHash,
		body.Role,
		body.FirstName,
		body.LastName,
		body.AvatarURL,
	)
	if err != nil {
		status := fiber.StatusInternalServerError
		switch {
		case errors.Is(err, services.ErrEmailExists):
			status = fiber.StatusConflict
		case errors.Is(err, services.ErrInvalidEmail):
			status = fiber.StatusBadRequest
		}
		return c.Status(status).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
		"data": fiber.Map{
			"username":  u.Username,
			"email":     u.Email,
			"role":      u.Role,
			"firstName": u.Profile.FirstName,
			"lastName":  u.Profile.LastName,
			"avatarURL": u.Profile.AvatarURL,
		},
	})
}
func ListUsers(c *fiber.Ctx) error {
	users := userService.ListUsers()

	list := make([]fiber.Map, 0, len(users))
	for _, u := range users {
		list = append(list, fiber.Map{
			"id":        u.ID,
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
	Username     *string `json:"username"`
	Email        *string `json:"email"`
	PasswordHash *string `json:"password_hash"`
	Role         *string `json:"role"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	AvatarURL    *string `json:"avatar_url"`
}

func v(s *string) string {
	if s == nil {
		return ""
	}
	return *s
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
	if len(u.Body()) == 0 {
		return u.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "request body required",
		})
	}

	if body.Username == nil {
		return u.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "username değeri girilmeli"})
	}
	if body.Email == nil {
		return u.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "email değeri girilmeli"})
	}
	if body.PasswordHash == nil {
		return u.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "password_hash değeri girilmeli"})
	}
	if body.Role == nil {
		return u.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "role değeri girilmeli"})
	}
	if body.FirstName == nil {
		return u.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "firstname değeri girilmeli"})
	}
	if body.LastName == nil {
		return u.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "lastname değeri girilmeli"})
	}
	if body.AvatarURL == nil {
		return u.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "avatarurl değeri girilmeli"})
	}

	if err := userService.UpdateUser(
		id,
		v(body.Username),
		v(body.Email),
		v(body.PasswordHash),
		v(body.Role),
		v(body.FirstName),
		v(body.LastName),
		v(body.AvatarURL),
	); err != nil {
		status := fiber.StatusInternalServerError
		switch {
		case errors.Is(err, services.ErrUserNotFound):
			status = fiber.StatusNotFound
		case errors.Is(err, services.ErrInvalidEmail):
			status = fiber.StatusBadRequest
		case errors.Is(err, services.ErrEmailExists):
			status = fiber.StatusConflict
		}
		return u.Status(status).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return u.JSON(fiber.Map{
		"success": true,
		"message": "User updated successfully",
		"data": fiber.Map{
			"username":  v(body.Username),
			"email":     v(body.Email),
			"role":      v(body.Role),
			"firstName": v(body.FirstName),
			"lastName":  v(body.LastName),
			"avatarURL": v(body.AvatarURL),
		},
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
