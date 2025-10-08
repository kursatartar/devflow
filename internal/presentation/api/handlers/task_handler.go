package handlers

import (
	"devflow/internal/presentation/api/converters"
	"devflow/internal/presentation/api/requests"
	"devflow/internal/presentation/api/responses"
	"devflow/internal/services"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	var body requests.CreateTaskReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
	}
    if err := validate.Struct(body); err != nil {
        return responses.JSON(c, 400, "validation error", map[string]any{"errors": buildValidationCauses(err)})
    }

    t, err := taskService.CreateTask(
        "",
		body.Title,
		body.Description,
		body.ProjectID,
		body.AssignedTo,
		body.CreatedBy,
		body.Status,
		body.Priority,
		body.DueDate,
		body.Labels,
        body.TimeTracking.Estimated,
        body.TimeTracking.Logged,
	)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrTaskExists):
			return responses.Conflict(c, err.Error())
		case errors.Is(err, services.ErrInvalidDueDate):
			return responses.ValidationError(c, err.Error())
		default:
			return responses.Internal(c, err)
		}
	}

	return responses.Created(c, "task created successfully", converters.ToTaskResponse(t))
}

func ListTasks(c *fiber.Ctx) error {
	ts := taskService.ListTasks()
	return responses.Success(c, "tasks fetched successfully", converters.ToTaskListResponse(ts))
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return responses.NotFound(c, "task not found")
	}

	var body requests.UpdateTaskReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
	}
	if len(c.Body()) == 0 {
		return responses.ValidationError(c, "request body required")
	}
    if err := validate.Struct(body); err != nil {
        return responses.JSON(c, 400, "validation error", map[string]any{"errors": buildValidationCauses(err)})
    }

	t, err := taskService.UpdateTask(
		id,
		body.Title,
		body.Description,
		body.Status,
		body.Priority,
		body.DueDate,
        body.Labels,
        func() *float64 { if body.TimeTracking!=nil { return body.TimeTracking.Estimated }; return nil }(),
        func() *float64 { if body.TimeTracking!=nil { return body.TimeTracking.Logged }; return nil }(),
	)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrTaskNotFound):
			return responses.NotFound(c, "task not found")
		case errors.Is(err, services.ErrInvalidDueDate):
			return responses.ValidationError(c, err.Error())
		default:
			return responses.Internal(c, err)
		}
	}

	return responses.Success(c, "task updated successfully", converters.ToTaskResponse(t))
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return responses.NotFound(c, "task not found")
	}
	if err := taskService.DeleteTask(id); err != nil {
		if errors.Is(err, services.ErrTaskNotFound) {
			return responses.NotFound(c, "task not found")
		}
		return responses.Internal(c, err)
	}
	return responses.Success(c, "task deleted successfully", nil)
}

func GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	t, err := taskService.GetTask(id)
	if err != nil {
		if errors.Is(err, services.ErrTaskNotFound) {
			return responses.NotFound(c, "task not found")
		}
		return responses.Internal(c, err)
	}

	return responses.Success(c, "task fetched successfully", converters.ToTaskResponse(t))
}
