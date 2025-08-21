package responses

import "time"

type TaskTimeTrackingResponse struct {
	EstimatedHours float64 `json:"estimated_hours"`
	LoggedHours    float64 `json:"logged_hours"`
}

type TaskResponse struct {
	ID           string                   `json:"id"`
	Title        string                   `json:"title"`
	Description  string                   `json:"description"`
	ProjectID    string                   `json:"project_id"`
	AssignedTo   string                   `json:"assigned_to"`
	CreatedBy    string                   `json:"created_by"`
	Status       string                   `json:"status"`
	Priority     string                   `json:"priority"`
	Labels       []string                 `json:"labels"`
	DueDate      string                   `json:"due_date"`
	TimeTracking TaskTimeTrackingResponse `json:"time_tracking"`
	CreatedAt    time.Time                `json:"created_at"`
	UpdatedAt    time.Time                `json:"updated_at"`
}

type TaskListResponse struct {
	Tasks    []TaskResponse `json:"tasks"`
	Metadata struct {
		Total int64 `json:"total"`
	} `json:"metadata"`
}
