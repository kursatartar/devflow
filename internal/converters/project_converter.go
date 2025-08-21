package converters

import (
	"devflow/internal/models"
	"devflow/internal/responses"
)

func ToProjectResponse(p *models.Project) responses.ProjectResponse {
	return responses.ProjectResponse{
		ID:           p.ID,
		Name:         p.Name,
		Description:  p.Description,
		OwnerID:      p.OwnerID,
		TeamMembers:  p.TeamMembers,
		Status:       p.Status,
		IsPrivate:    p.Settings.IsPrivate,
		TaskWorkflow: p.Settings.TaskWorkflow,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func ToProjectListResponse(ps []*models.Project) responses.ProjectListResponse {
	items := make([]responses.ProjectResponse, 0, len(ps))
	for _, p := range ps {
		items = append(items, ToProjectResponse(p))
	}
	var out responses.ProjectListResponse
	out.Projects = items
	out.Metadata.Total = int64(len(items))
	return out
}
