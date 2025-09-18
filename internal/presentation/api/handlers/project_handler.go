package handlers

import (
	"devflow/internal/presentation/api/converters"
	"devflow/internal/presentation/api/requests"
	"devflow/internal/presentation/api/responses"
	"devflow/internal/services"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func CreateProject(c *fiber.Ctx) error {
	var body requests.CreateProjectReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
	}
	if body.Name == "" || body.OwnerID == "" || body.TeamID == "" || body.Status == "" {
		return responses.ValidationError(c, "missing required fields")
	}
	p, err := projectService.CreateProject(
		"",
		body.Name,
		body.Description,
		body.OwnerID,
		body.TeamID,
		body.Status,
		body.TeamMembers,
		body.IsPrivate,
		body.TaskWorkflow,
	)
	if err != nil {
		return responses.ValidationError(c, err.Error())
	}
	return responses.Created(c, "project created successfully", converters.ToProjectResponse(p))
}

func ListProjects(c *fiber.Ctx) error {
	ps := projectService.ListProjects()
	// Liste uç noktasında şimdilik team detayını iliştirmiyoruz; sadece project alanları döner
	return responses.Success(c, "projects fetched successfully", converters.ToProjectListResponse(ps))
}

func UpdateProject(c *fiber.Ctx) error {
	id := c.Params("id")
	if len(c.Body()) == 0 {
		return responses.ValidationError(c, "request body required")
	}
	var body requests.UpdateProjectReq
	if err := c.BodyParser(&body); err != nil {
		return responses.ValidationError(c, "invalid json")
	}
	p, err := projectService.UpdateProject(
		id,
		body.Name,
		body.Description,
		body.Status,
		body.TeamID,
		body.TeamMembers,
		body.IsPrivate,
		body.TaskWorkflow,
	)
	if err != nil {
		if errors.Is(err, services.ErrProjectNotFound) {
			return responses.NotFound(c, "project not found")
		}
		return responses.Internal(c, err)
	}
	return responses.Success(c, "project updated successfully", converters.ToProjectResponse(p))
}

func DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := projectService.DeleteProject(id); err != nil {
		if errors.Is(err, services.ErrProjectNotFound) {
			return responses.NotFound(c, "project not found")
		}
		return responses.Internal(c, err)
	}
	return responses.Success(c, "project deleted successfully", nil)
}

func GetProject(c *fiber.Ctx) error {
	id := c.Params("id")
	p, err := projectService.GetProject(id)
	if err != nil {
		if errors.Is(err, services.ErrProjectNotFound) {
			return responses.NotFound(c, "project not found")
		}
		return responses.Internal(c, err)
	}
	var teamResp *responses.TeamResponse
	if p.TeamID != "" {
		t, terr := teamService.GetTeam(p.TeamID)
		if terr == nil && t != nil {
			tr := converters.ToTeamResponse(t)
			teamResp = &tr
		}
	}
	resp := converters.ToProjectResponse(p)
	resp.Team = teamResp
	return responses.Success(c, "project fetched successfully", resp)
}
