package handlers

import (
	"devflow/internal/requests"
	"errors"
	"fmt"

	"devflow/internal/models"
	"devflow/internal/services"
	"github.com/gofiber/fiber/v2"
)

var teamService = services.NewTeamService()

func teamResource(t *models.Team) fiber.Map {
	ms := make([]fiber.Map, 0, len(t.Members))
	for _, m := range t.Members {
		ms = append(ms, fiber.Map{
			"userId":   m.UserID,
			"role":     m.Role,
			"joinedAt": m.JoinedAt,
		})
	}
	return fiber.Map{
		"id":          t.ID,
		"name":        t.Name,
		"description": t.Description,
		"ownerId":     t.OwnerID,
		"members":     ms,
		"settings": fiber.Map{
			"isPrivate":         t.Settings.IsPrivate,
			"allowMemberInvite": t.Settings.AllowMemberInvite,
		},
		"createdAt": t.CreatedAt,
		"updatedAt": t.UpdatedAt,
	}
}

func CreateTeam(c *fiber.Ctx) error {
	var body requests.CreateTeamReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "invalid json"})
	}
	if body.Name == "" || body.OwnerID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "missing required fields"})
	}
	members := make([]models.TeamMember, 0, len(body.Members))
	for _, m := range body.Members {
		members = append(members, models.TeamMember{UserID: m.UserID, Role: m.Role})
	}
	t, err := teamService.CreateTeam("", body.Name, body.Description, body.OwnerID, members, body.Settings)
	if err != nil {
		code := fiber.StatusInternalServerError
		if errors.Is(err, services.ErrTeamExists) {
			code = fiber.StatusConflict
		}
		return c.Status(code).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Team created successfully",
		"data":    teamResource(t),
	})
}

func ListTeams(c *fiber.Ctx) error {
	teams := teamService.ListTeams()
	out := make([]fiber.Map, 0, len(teams))
	for _, t := range teams {
		out = append(out, teamResource(t))
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Teams fetched successfully",
		"data":    out,
	})
}

func GetTeam(c *fiber.Ctx) error {
	id := c.Params("id")
	t, err := teamService.GetTeam(id)
	if err != nil {
		code := fiber.StatusInternalServerError
		if errors.Is(err, services.ErrTeamNotFound) {
			code = fiber.StatusNotFound
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": fmt.Sprintf("team %s not found", id)})
		}
		return c.Status(code).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Team fetched successfully",
		"data":    teamResource(t),
	})
}

func UpdateTeam(c *fiber.Ctx) error {
	id := c.Params("id")
	var body requests.UpdateTeamReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "invalid json"})
	}
	if len(c.Body()) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "request body required"})
	}
	t, err := teamService.UpdateTeam(id, body.Name, body.Description, body.Settings)
	if err != nil {
		code := fiber.StatusInternalServerError
		if errors.Is(err, services.ErrTeamNotFound) {
			code = fiber.StatusNotFound
			return c.Status(code).JSON(fiber.Map{"success": false, "message": fmt.Sprintf("team %s not found", id)})
		}
		return c.Status(code).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Team updated successfully",
		"data":    teamResource(t),
	})
}

func AddTeamMember(c *fiber.Ctx) error {
	id := c.Params("id")
	var body requests.AddMemberReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "invalid json"})
	}
	if body.UserID == "" || body.Role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "missing required fields"})
	}
	t, err := teamService.AddMember(id, body.UserID, body.Role)
	if err != nil {
		code := fiber.StatusInternalServerError
		if errors.Is(err, services.ErrTeamNotFound) {
			code = fiber.StatusNotFound
		}
		return c.Status(code).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Team member added successfully",
		"data":    teamResource(t),
	})
}

func DeleteTeam(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := teamService.DeleteTeam(id); err != nil {
		code := fiber.StatusInternalServerError
		if errors.Is(err, services.ErrTeamNotFound) {
			code = fiber.StatusNotFound
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": fmt.Sprintf("team %s not found", id),
			})
		}
		return c.Status(code).JSON(fiber.Map{"success": false, "message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("team %s deleted successfully", id),
	})
}
