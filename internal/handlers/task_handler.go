package handlers

import (
	"devflow/internal/requests"
	"errors"

	"devflow/internal/services"
	"github.com/gofiber/fiber/v2"
)

var taskService = services.NewTaskService()

func taskResource(t interface {
	GetID() string
}) fiber.Map {
	return fiber.Map{}
}

func mapTask(t interface {
	GetID() string
}) fiber.Map {
	type tt interface {
		GetID() string
	}
	_ = tt(t)
	return fiber.Map{}
}

func CreateTask(c *fiber.Ctx) error {
	var body requests.CreateTaskReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "invalid json"})
	}
	if body.Title == "" || body.ProjectID == "" || body.CreatedBy == "" || body.Status == "" || body.Priority == "" || body.DueDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "missing required fields"})
	}

	t, err := taskService.CreateTask("", body.Title, body.Description, body.ProjectID, body.AssignedTo, body.CreatedBy, body.Status, body.Priority, body.DueDate, body.Labels, body.Estimated, body.Logged)
	if err != nil {
		code := fiber.StatusInternalServerError
		switch {
		case errors.Is(err, services.ErrTaskExists):
			code = fiber.StatusConflict
		case errors.Is(err, services.ErrInvalidDueDate):
			code = fiber.StatusBadRequest
		}
		return c.Status(code).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Task created successfully",
		"data": fiber.Map{
			"id":          t.ID,
			"title":       t.Title,
			"description": t.Description,
			"projectId":   t.ProjectID,
			"assignedTo":  t.AssignedTo,
			"createdBy":   t.CreatedBy,
			"status":      t.Status,
			"priority":    t.Priority,
			"labels":      t.Labels,
			"dueDate":     t.DueDate,
			"timeTracking": fiber.Map{
				"estimated_hours": t.TimeTracking.EstimatedHours,
				"logged_hours":    t.TimeTracking.LoggedHours,
			},
			"createdAt": t.CreatedAt,
			"updatedAt": t.UpdatedAt,
		},
	})
}

func ListTasks(c *fiber.Ctx) error {
	ts := taskService.ListTasks()
	list := make([]fiber.Map, 0, len(ts))
	for _, t := range ts {
		list = append(list, fiber.Map{
			"id":          t.ID,
			"title":       t.Title,
			"description": t.Description,
			"projectId":   t.ProjectID,
			"assignedTo":  t.AssignedTo,
			"createdBy":   t.CreatedBy,
			"status":      t.Status,
			"priority":    t.Priority,
			"labels":      t.Labels,
			"dueDate":     t.DueDate,
			"timeTracking": fiber.Map{
				"estimated_hours": t.TimeTracking.EstimatedHours,
				"logged_hours":    t.TimeTracking.LoggedHours,
			},
			"createdAt": t.CreatedAt,
			"updatedAt": t.UpdatedAt,
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Tasks fetched successfully",
		"data":    list,
	})
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")

	var body requests.UpdateTaskReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "invalid json"})
	}
	if len(c.Body()) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "request body required"})
	}

	t, err := taskService.UpdateTask(id, body.Title, body.Description, body.Status, body.Priority, body.DueDate, body.Labels, body.Estimated, body.Logged)
	if err != nil {
		code := fiber.StatusInternalServerError
		switch {
		case errors.Is(err, services.ErrTaskNotFound):
			code = fiber.StatusNotFound
		case errors.Is(err, services.ErrInvalidDueDate):
			code = fiber.StatusBadRequest
		}
		return c.Status(code).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Task updated successfully",
		"data": fiber.Map{
			"id":          t.ID,
			"title":       t.Title,
			"description": t.Description,
			"projectId":   t.ProjectID,
			"assignedTo":  t.AssignedTo,
			"createdBy":   t.CreatedBy,
			"status":      t.Status,
			"priority":    t.Priority,
			"labels":      t.Labels,
			"dueDate":     t.DueDate,
			"timeTracking": fiber.Map{
				"estimated_hours": t.TimeTracking.EstimatedHours,
				"logged_hours":    t.TimeTracking.LoggedHours,
			},
			"createdAt": t.CreatedAt,
			"updatedAt": t.UpdatedAt,
		},
	})
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := taskService.DeleteTask(id); err != nil {
		code := fiber.StatusInternalServerError
		if errors.Is(err, services.ErrTaskNotFound) {
			code = fiber.StatusNotFound
		}
		return c.Status(code).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Task deleted successfully",
	})
}

func GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	t, err := taskService.GetTask(id)
	if err != nil {
		status := fiber.StatusInternalServerError
		if errors.Is(err, services.ErrTaskNotFound) {
			status = fiber.StatusNotFound
		}
		return c.Status(status).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Task fetched successfully",
		"data": fiber.Map{
			"ID":          t.ID,
			"Title":       t.Title,
			"Description": t.Description,
			"ProjectId":   t.ProjectID,
			"AssignedTo":  t.AssignedTo,
			"CreatedBy":   t.CreatedBy,
			"Status":      t.Status,
			"Priority":    t.Priority,
			"Labels":      t.Labels,
			"DueDate":     t.DueDate,
			"timeTracking": fiber.Map{
				"estimated_hours": t.TimeTracking.EstimatedHours,
				"logged_hours":    t.TimeTracking.LoggedHours,
			},
		},
	})
}
