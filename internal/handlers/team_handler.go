package handlers

import (
	"devflow/internal/converters"
	"devflow/internal/models"
	"devflow/internal/requests"
	"devflow/internal/responses"
	"devflow/internal/services"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateTeam(c *fiber.Ctx) error {
	var body requests.CreateTeamReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
	}
	if body.Name == "" || body.OwnerID == "" {
		return responses.ValidationError(c, "missing required fields")
	}
	members := make([]models.TeamMember, 0, len(body.Members))
	for _, m := range body.Members {
		members = append(members, models.TeamMember{UserID: m.UserID, Role: m.Role})
	}
	t, err := teamService.CreateTeam("", body.Name, body.Description, body.OwnerID, members, body.Settings)
	if err != nil {
		if errors.Is(err, services.ErrTeamExists) {
			return responses.Conflict(c, err.Error())
		}
		return responses.Internal(c, err)
	}
	return responses.Created(c, "team created successfully", converters.ToTeamResponse(t))
}

func ListTeams(c *fiber.Ctx) error {
	teams := teamService.ListTeams()
	return responses.Success(c, "teams fetched successfully", converters.ToTeamListResponse(teams))
}

func GetTeam(c *fiber.Ctx) error {
	id := c.Params("id")
	t, err := teamService.GetTeam(id)
	if err != nil {
		if errors.Is(err, services.ErrTeamNotFound) {
			return responses.NotFound(c, fmt.Sprintf("team %s not found", id))
		}
		return responses.Internal(c, err)
	}
	return responses.Success(c, "team fetched successfully", converters.ToTeamResponse(t))
}

func UpdateTeam(c *fiber.Ctx) error {
	id := c.Params("id")
	var body requests.UpdateTeamReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
	}
	if len(c.Body()) == 0 {
		return responses.ValidationError(c, "request body required")
	}
	t, err := teamService.UpdateTeam(id, body.Name, body.Description, body.Settings)
	if err != nil {
		if errors.Is(err, services.ErrTeamNotFound) {
			return responses.NotFound(c, fmt.Sprintf("team %s not found", id))
		}
		return responses.Internal(c, err)
	}
	return responses.Success(c, "team updated successfully", converters.ToTeamResponse(t))
}

func AddTeamMember(c *fiber.Ctx) error {
	id := c.Params("id")
	var body requests.AddMemberReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
	}
	if body.UserID == "" || body.Role == "" {
		return responses.ValidationError(c, "missing required fields")
	}
	t, err := teamService.AddMember(id, body.UserID, body.Role)
	if err != nil {
		if errors.Is(err, services.ErrTeamNotFound) {
			return responses.NotFound(c, "team not found")
		}
		return responses.Internal(c, err)
	}
	return responses.Success(c, "team member added successfully", converters.ToTeamResponse(t))
}

func DeleteTeam(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := teamService.DeleteTeam(id); err != nil {
		if errors.Is(err, services.ErrTeamNotFound) {
			return responses.NotFound(c, fmt.Sprintf("team %s not found", id))
		}
		return responses.Internal(c, err)
	}
	return responses.Success(c, "team deleted successfully", nil)
}
