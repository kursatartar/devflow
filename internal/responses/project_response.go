package responses

import "devflow/internal/models"

func ProjectResource(p *models.Project) map[string]any {
	return map[string]any{
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
	}
}

func ProjectList(ps []*models.Project) []map[string]any {
	out := make([]map[string]any, 0, len(ps))
	for _, p := range ps {
		out = append(out, ProjectResource(p))
	}
	return out
}
