package handlers

import (
	"devflow/internal/services"
	"errors"
	"github.com/gofiber/fiber/v2"
)

var projectService = services.NewProjectService()

func CreateProject(c *fiber.Ctx) error {
	var body struct {
		ID           string   `json:"id"`
		Name         string   `json:"name"`
		Description  string   `json:"description"`
		OwnerID      string   `json:"owner_id"`
		Status       string   `json:"status"`
		TeamMembers  []string `json:"team_members"`
		IsPrivate    bool     `json:"is_private"`
		TaskWorkflow []string `json:"task_workflow"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "invalid json"})
	}

	project, err := projectService.CreateProject(body.ID, body.Name, body.Description, body.OwnerID, body.Status, body.TeamMembers, body.IsPrivate, body.TaskWorkflow)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Project created successfully", "data": project})
}

func ListProjects(c *fiber.Ctx) error {
	projects := projectService.ListProjects()
	return c.JSON(fiber.Map{"success": true, "message": "Projects fetched successfully", "data": projects})
}

func UpdateProject(c *fiber.Ctx) error {
	id := c.Params("id")
	var body struct {
		Name         string   `json:"name"`
		Description  string   `json:"description"`
		Status       string   `json:"status"`
		TeamMembers  []string `json:"team_members"`
		IsPrivate    bool     `json:"is_private"`
		TaskWorkflow []string `json:"task_workflow"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "invalid json"})
	}

	project, err := projectService.UpdateProject(id, body.Name, body.Description, body.Status, body.TeamMembers, body.IsPrivate, body.TaskWorkflow)
	if err != nil {
		status := fiber.StatusInternalServerError
		if err == services.ErrProjectNotFound {
			status = fiber.StatusNotFound
		}
		return c.Status(status).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Project updated successfully", "data": project})
}

func DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := projectService.DeleteProject(id); err != nil {
		status := fiber.StatusInternalServerError
		if err == services.ErrProjectNotFound {
			status = fiber.StatusNotFound
		}
		return c.Status(status).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "message": "Project deleted successfully"})
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
			"Id":          p.ID,
			"Name":        p.Name,
			"Description": p.Description,
			"Status":      p.Status,
			"TeamMembers": p.TeamMembers,
		},
	})
}
