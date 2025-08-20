package responses

import "devflow/internal/models"

func TaskResource(t *models.Task) map[string]any {
	return map[string]any{
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
		"timeTracking": map[string]any{
			"estimated_hours": t.TimeTracking.EstimatedHours,
			"logged_hours":    t.TimeTracking.LoggedHours,
		},
		"createdAt": t.CreatedAt,
		"updatedAt": t.UpdatedAt,
	}
}

func TaskList(ts []*models.Task) []map[string]any {
	out := make([]map[string]any, 0, len(ts))
	for _, t := range ts {
		out = append(out, TaskResource(t))
	}
	return out
}
