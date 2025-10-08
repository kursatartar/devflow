package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"devflow/internal/presentation/api/converters"
	"devflow/internal/presentation/api/requests"
	"devflow/internal/presentation/api/responses"
	"devflow/internal/services"
	"devflow/utils"
)

func CreateTeam(c *fiber.Ctx) error {
	var body requests.CreateTeamReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
	}
    if err := utils.Validate.Struct(body); err != nil {
        return responses.JSON(c, 400, "validation error", map[string]any{"errors": utils.BuildValidationCauses(err)})
    }

	members := converters.ToDomainTeamMembers(body.Members)
	settings := converters.ToDomainTeamSettings(body.Settings)

	t, err := teamService.CreateTeam(
		"",
		body.Name,
		body.Description,
		body.OwnerID,
		members,
		settings,
	)
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
	if id == "" {
		return responses.NotFound(c, "team not found")
	}

    var body requests.UpdateTeamReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
	}
	if len(c.Body()) == 0 {
		return responses.ValidationError(c, "request body required")
	}
    if err := validate.Struct(body); err != nil {
        return responses.JSON(c, 400, "validation error", map[string]any{"errors": buildValidationCauses(err)})
    }

	domSettings := converters.ToDomainTeamSettingsPtr(body.Settings)

	t, err := teamService.UpdateTeam(
		id,
		body.Name,
		body.Description,
		domSettings,
	)
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
    if err := utils.Validate.Struct(body); err != nil {
        return responses.JSON(c, 400, "validation error", map[string]any{"errors": utils.BuildValidationCauses(err)})
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
	if id == "" {
		return responses.NotFound(c, "team not found")
	}
	if err := teamService.DeleteTeam(id); err != nil {
		if errors.Is(err, services.ErrTeamNotFound) {
			return responses.NotFound(c, fmt.Sprintf("team %s not found", id))
		}
		return responses.Internal(c, err)
	}
	return responses.Success(c, "team deleted successfully", nil)
}
