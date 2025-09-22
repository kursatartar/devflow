package handlers

import (
    "devflow/internal/presentation/api/converters"
    "devflow/internal/presentation/api/requests"
    "devflow/internal/presentation/api/responses"
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
    if err := validate.Struct(body); err != nil {
        return responses.JSON(c, 400, "validation error", map[string]any{"errors": buildValidationCauses(err)})
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
    if err := validate.Struct(body); err != nil {
        return responses.JSON(c, 400, "validation error", map[string]any{"errors": buildValidationCauses(err)})
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

func Register(c *fiber.Ctx) error {
	var body requests.RegisterReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
	}
	if err := validate.Struct(body); err != nil {
		return responses.JSON(c, 400, "validation error", map[string]any{"errors": buildValidationCauses(err)})
	}
	u, err := userService.Register(body.Username, body.Email, body.Password, body.FirstName, body.LastName, body.AvatarURL)
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

	// Generate token
	token, err := authService.GenerateToken(u.ID, u.Username, u.Role)
	if err != nil {
		return responses.Internal(c, err)
	}

	return responses.Created(c, "registered successfully", fiber.Map{
		"user":  converters.ToUserResponse(u),
		"token": token,
	})
}

func Login(c *fiber.Ctx) error {
	var body requests.LoginReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
	}
	if err := validate.Struct(body); err != nil {
		return responses.JSON(c, 400, "validation error", map[string]any{"errors": buildValidationCauses(err)})
	}
	u, err := userService.Authenticate(body.Identifier, body.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			return responses.Unauthorized(c, "invalid credentials")
		}
		return responses.Internal(c, err)
	}

	// Generate token
	token, err := authService.GenerateToken(u.ID, u.Username, u.Role)
	if err != nil {
		return responses.Internal(c, err)
	}

	return responses.Success(c, "login successful", fiber.Map{
		"user":  converters.ToUserResponse(u),
		"token": token,
	})
}
