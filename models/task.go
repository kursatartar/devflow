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
	CreatedAt    string
	UpdatedAt    string
}

var Tasks = map[string]Task{}

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
