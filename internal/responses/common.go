package responses

import (
	"github.com/gofiber/fiber/v2"
)

type Envelope map[string]any

func write(c *fiber.Ctx,
	code int,
	message string,
	data any) error {
	payload := Envelope{
		"success": code >= 200 && code < 300,
		"message": message,
	}
	if data != nil {
		payload["data"] = data
	}
	return c.Status(code).JSON(payload)
}

func Success(c *fiber.Ctx, message string, data any) error {
	return write(c, fiber.StatusOK, message, data)
}

func Created(c *fiber.Ctx, message string, data any) error {
	return write(c, fiber.StatusCreated, message, data)
}

func ValidationError(c *fiber.Ctx, message string) error {
	return write(c, fiber.StatusBadRequest, message, nil)
}

func NotFound(c *fiber.Ctx, message string) error {
	return write(c, fiber.StatusNotFound, message, nil)
}

func Conflict(c *fiber.Ctx, message string) error {
	return write(c, fiber.StatusConflict, message, nil)
}

func Unauthorized(c *fiber.Ctx, message string) error {
	return write(c, fiber.StatusUnauthorized, message, nil)
}

func Forbidden(c *fiber.Ctx, message string) error {
	return write(c, fiber.StatusForbidden, message, nil)
}

func Internal(c *fiber.Ctx, err error) error {
	msg := "internal server error"
	if err != nil && err.Error() != "" {
		msg = err.Error()
	}
	return write(c, fiber.StatusInternalServerError, msg, nil)
}

func JSON(c *fiber.Ctx, code int, message string, data any) error {
	return write(c, code, message, data)
}
