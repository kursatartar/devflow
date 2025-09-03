package handlers

import (
	"devflow/internal/converters"
	"devflow/internal/requests"
	"devflow/internal/responses"
	"devflow/internal/services"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var body requests.CreateUserReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
	}

	if body.Username == "" {
		return responses.ValidationError(c, "username boş olamaz")
	}
	if body.Email == "" {
		return responses.ValidationError(c, "email boş olamaz")
	}
	if body.PasswordHash == "" {
		return responses.ValidationError(c, "hash boş olamaz")
	}
	if body.Role == "" {
		return responses.ValidationError(c, "role boş olamaz")
	}
	if body.FirstName == "" {
		return responses.ValidationError(c, "firstname boş olamaz")
	}
	if body.LastName == "" {
		return responses.ValidationError(c, "lastname boş olamaz")
	}
	if body.AvatarURL == "" {
		return responses.ValidationError(c, "avatarurl boş olamaz")
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
		switch {
		case errors.Is(err, services.ErrEmailExists):

			return responses.Conflict(c, err.Error())
		case errors.Is(err, services.ErrInvalidEmail):
			return responses.ValidationError(c, err.Error())
		default:
			return responses.Internal(c, err)
		}

	}

	return responses.Created(c, "user created successfully", converters.ToUserResponse(u))
}
func ListUsers(c *fiber.Ctx) error {
	users := userService.ListUsers()
	return responses.Success(c, "users fetched successfully", converters.ToUserListResponse(users))
}

func v(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var body requests.UpdateUserReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
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
		switch {
		case errors.Is(err, services.ErrUserNotFound):
			return responses.NotFound(c, fmt.Sprintf("user %s not found", id))
		case errors.Is(err, services.ErrEmailExists):
			return responses.Conflict(c, "email already exists")
		case errors.Is(err, services.ErrInvalidEmail):
			return responses.ValidationError(c, "invalid email format")
		default:
			return responses.Internal(c, err)
		}
	}

	u, err := userService.GetUser(id)
	if err != nil {
		return responses.Internal(c, err)
	}

	return responses.Success(c, "user updated successfully", converters.ToUserResponse(u))
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("_id")
	if err := userService.DeleteUser(id); err != nil {

		if errors.Is(err, services.ErrUserNotFound) {
			return responses.NotFound(c, fmt.Sprintf("user %s not found", id))
		}
		return responses.Internal(c, err)
	}
	return responses.Success(c, "user deleted successfully", nil)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("_id")
	u, err := userService.GetUser(id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return responses.NotFound(c, fmt.Sprintf("user %s not found", id))
		}
		return responses.Internal(c, err)
	}

	return responses.Success(c, "user fetched successfully", converters.ToUserResponse(u))
}
