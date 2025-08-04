package models

import "time"

type TimeTracking struct {
	EstimatedHours float64
	LoggedHours    float64
}

type Task struct {
	ID           string
	Title        string
	Description  string
	ProjectID    string
	AssignedTo   string
	CreatedBy    string
	Status       string
	Priority     string
	Labels       []string
	DueDate      string
	TimeTracking TimeTracking
	CreatedAt    string // time.time
	UpdatedAt    string
}

// pointerı tek bir değerde deneyelim

var Tasks = map[string]*Task{}

func NewTask(id, title, description, projectID, assignedTo, createdBy, status, priority, dueDate string, labels []string, estimated, logged float64) Task {
	return Task{
		ID:          id,
		Title:       title,
		Description: description,
		ProjectID:   projectID,
		AssignedTo:  assignedTo,
		CreatedBy:   createdBy,
		Status:      status,
		Priority:    priority,
		Labels:      labels,
		DueDate:     dueDate,
		TimeTracking: TimeTracking{
			EstimatedHours: estimated,
			LoggedHours:    logged,
		},
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
}

// type Task struct {
//   ID           string       // `json:"id"`
//   Title        string       // `json:"title"`
//   Description  string       // `json:"description"`
//   ProjectID    string       // `json:"project_id"`
//   AssignedTo   string       // `json:"assigned_to"`
//   CreatedBy    string       // `json:"created_by"`
//   Status       string       // `json:"status"`
//   Priority     string       // `json:"priority"`
//   Labels       []string     // `json:"labels"`
//   DueDate      string       // `json:"due_date"`
//   TimeTracking TimeTracking // `json:"time_tracking"`
//   CreatedAt    string       // `json:"created_at"`
//   UpdatedAt    string       // `json:"updated_at"`
// }

// type TimeTracking struct {
//   EstimatedHours float64 // `json:"estimated_hours"`
//   LoggedHours    float64 // `json:"logged_hours"`
// }
