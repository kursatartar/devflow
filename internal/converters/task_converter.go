package converters

import (
	"devflow/internal/models"
	"devflow/internal/responses"
)

func ToTaskResponse(t *models.Task) responses.TaskResponse {
	return responses.TaskResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		ProjectID:   t.ProjectID,
		AssignedTo:  t.AssignedTo,
		CreatedBy:   t.CreatedBy,
		Status:      t.Status,
		Priority:    t.Priority,
		Labels:      t.Labels,
		DueDate:     t.DueDate,
		TimeTracking: responses.TaskTimeTrackingResponse{
			EstimatedHours: t.TimeTracking.EstimatedHours,
			LoggedHours:    t.TimeTracking.LoggedHours,
		},
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func ToTaskListResponse(ts []*models.Task) responses.TaskListResponse {
	items := make([]responses.TaskResponse, 0, len(ts))
	for _, t := range ts {
		items = append(items, ToTaskResponse(t))
	}
	var out responses.TaskListResponse
	out.Tasks = items
	out.Metadata.Total = int64(len(items))
	return out
}
