package handlers

import (
	"devflow/internal/converters"
	"devflow/internal/requests"
	"devflow/internal/responses"
	"devflow/internal/services"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	var body requests.CreateTaskReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
	}
	if body.Title == "" || body.ProjectID == "" || body.CreatedBy == "" || body.Status == "" || body.Priority == "" || body.DueDate == "" {
		return responses.ValidationError(c, "missing required fields")
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
		body.Estimated,
		body.Logged,
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

	var body requests.UpdateTaskReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
	}
	if len(c.Body()) == 0 {
		return responses.ValidationError(c, "request body required")
	}

	t, err := taskService.UpdateTask(
		id,
		body.Title,
		body.Description,
		body.Status,
		body.Priority,
		body.DueDate,
		body.Labels,
		body.Estimated,
		body.Logged,
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
