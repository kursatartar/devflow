package handlers

import (
	"devflow/internal/requests"
	"errors"

	"devflow/internal/services"
	"github.com/gofiber/fiber/v2"
)

var projectService = services.NewProjectService()

func CreateProject(c *fiber.Ctx) error {
	var body requests.CreateProjectReq
	if err := c.BodyParser(&body); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"success": false, "message": "invalid json"})
	}

	if body.Name == "" || body.OwnerID == "" || body.Status == "" {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"success": false, "message": "missing required fields"})
	}

	p, err := projectService.CreateProject(
		"",
		body.Name,
		body.Description,
		body.OwnerID,
		body.Status,
		body.TeamMembers,
		body.IsPrivate,
		body.TaskWorkflow,
	)
	if err != nil {

		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	return c.
		Status(fiber.StatusCreated).
		JSON(fiber.Map{
			"success": true,
			"message": "Project created successfully",
			"data": fiber.Map{
				"id":           p.ID,
				"name":         p.Name,
				"description":  p.Description,
				"ownerId":      p.OwnerID,
				"status":       p.Status,
				"teamMembers":  p.TeamMembers,
				"isPrivate":    p.Settings.IsPrivate,
				"taskWorkflow": p.Settings.TaskWorkflow,
				"createdAt":    p.CreatedAt,
				"updatedAt":    p.UpdatedAt,
			},
		})
}

func ListProjects(c *fiber.Ctx) error {
	ps := projectService.ListProjects()

	out := make([]fiber.Map, 0, len(ps))
	for _, p := range ps {
		out = append(out, fiber.Map{
			"id":           p.ID,
			"name":         p.Name,
			"description":  p.Description,
			"ownerId":      p.OwnerID,
			"status":       p.Status,
			"teamMembers":  p.TeamMembers,
			"isPrivate":    p.Settings.IsPrivate,
			"taskWorkflow": p.Settings.TaskWorkflow,
			"createdAt":    p.CreatedAt,
			"updatedAt":    p.UpdatedAt,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Projects fetched successfully",
		"data":    out,
	})
}

func UpdateProject(c *fiber.Ctx) error {
	id := c.Params("id")

	if len(c.Body()) == 0 {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"success": false, "message": "request body required"})
	}

	var body requests.UpdateProjectReq
	if err := c.BodyParser(&body); err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"success": false, "message": "invalid json"})
	}

	p, err := projectService.UpdateProject(
		id,
		body.Name,
		body.Description,
		body.Status,
		body.TeamMembers,
		body.IsPrivate,
		body.TaskWorkflow,
	)
	if err != nil {
		status := fiber.StatusInternalServerError
		if errors.Is(err, services.ErrProjectNotFound) {
			status = fiber.StatusNotFound
		}
		return c.Status(status).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Project updated successfully",
		"data": fiber.Map{
			"id":           p.ID,
			"name":         p.Name,
			"description":  p.Description,
			"ownerId":      p.OwnerID,
			"status":       p.Status,
			"teamMembers":  p.TeamMembers,
			"isPrivate":    p.Settings.IsPrivate,
			"taskWorkflow": p.Settings.TaskWorkflow,
			"createdAt":    p.CreatedAt,
			"updatedAt":    p.UpdatedAt,
		},
	})
}

func DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := projectService.DeleteProject(id); err != nil {
		status := fiber.StatusInternalServerError
		if errors.Is(err, services.ErrProjectNotFound) {
			status = fiber.StatusNotFound
		}
		return c.Status(status).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Project deleted successfully",
	})
}

func GetProject(c *fiber.Ctx) error {
	id := c.Params("id")
	p, err := projectService.GetProject(id)
	if err != nil {
		status := fiber.StatusInternalServerError
		if errors.Is(err, services.ErrProjectNotFound) {
			status = fiber.StatusNotFound
		}
		return c.Status(status).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Project fetched successfully",
		"data": fiber.Map{
			"id":           p.ID,
			"name":         p.Name,
			"description":  p.Description,
			"ownerId":      p.OwnerID,
			"status":       p.Status,
			"teamMembers":  p.TeamMembers,
			"isPrivate":    p.Settings.IsPrivate,
			"taskWorkflow": p.Settings.TaskWorkflow,
			"createdAt":    p.CreatedAt,
			"updatedAt":    p.UpdatedAt,
		},
	})
}
