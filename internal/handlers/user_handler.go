package handlers

import (
	"devflow/internal/requests"
	"devflow/internal/responses"
	"devflow/internal/services"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

var userService = services.NewUserService()

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

	return responses.Created(c, "user created successfully", responses.UserResource(u))
}
func ListUsers(c *fiber.Ctx) error {
	users := userService.ListUsers()
	return responses.Success(c, "user fetched successfully", responses.UserList(users))
}

func v(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func UpdateUser(u *fiber.Ctx) error {
	id := u.Params("id")

	var body requests.UpdateUserReq
	if err := u.BodyParser(&body); err != nil {
		return responses.ValidationError(u, "invalid json")
	}
	if len(u.Body()) == 0 {
		return responses.ValidationError(u, "request body girilmeli")
	}

	if body.Username == nil {
		return responses.ValidationError(u, "username değeri girilmeli")
	}
	if body.Email == nil {
		return responses.ValidationError(u, "email değeri girilmeli")
	}
	if body.PasswordHash == nil {
		return responses.ValidationError(u, "hash değeri girilmeli")
	}
	if body.Role == nil {
		return responses.ValidationError(u, "role değeri girilmeli")
	}
	if body.FirstName == nil {
		return responses.ValidationError(u, "firstname değeri girilmeli")
	}
	if body.LastName == nil {
		return responses.ValidationError(u, "lastname değeri girilmeli")
	}
	if body.AvatarURL == nil {
		return responses.ValidationError(u, "avatarurl değeri girilmeli")
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
		switch {
		case errors.Is(err, services.ErrUserNotFound):
			return responses.NotFound(u, fmt.Sprintf("user %s not found", id))
		case errors.Is(err, services.ErrInvalidEmail):
			return responses.ValidationError(u, err.Error())
		case errors.Is(err, services.ErrEmailExists):
			return responses.Conflict(u, err.Error())
		default:
			return responses.Internal(u, err)

		}

	}
	return responses.Success(u, "user updated succesfully", map[string]any{
		"username":  v(body.Username),
		"email":     v(body.Email),
		"password":  v(body.PasswordHash),
		"role":      v(body.Role),
		"firstName": v(body.FirstName),
		"lastName":  v(body.LastName),
		"avatarUrl": v(body.AvatarURL),
	})

}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := userService.DeleteUser(id); err != nil {

		if errors.Is(err, services.ErrUserNotFound) {
			return responses.NotFound(c, fmt.Sprintf("user %s not found", id))
		}
		return responses.Internal(c, err)
	}
	return responses.Success(c, "user deleted successfully", nil)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	u, err := userService.GetUser(id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return responses.NotFound(c, fmt.Sprintf("user %s not found", id))
		}
		return responses.Internal(c, err)
	}

	return responses.Success(c, "user fetched successfully", responses.UserResource(u))
}
