package models

import "time"

type TimeTracking struct {
	EstimatedHours float64 `json:"estimated_hours" bson:"estimated_hours"`
	LoggedHours    float64 `json:"logged_hours"    bson:"logged_hours"`
}

type Task struct {
	ID           string       `json:"id"            bson:"_id,omitempty"`
	Title        string       `json:"title"         bson:"title"`
	Description  string       `json:"description"   bson:"description"`
	ProjectID    string       `json:"project_id"    bson:"project_id"`
	AssignedTo   string       `json:"assigned_to"   bson:"assigned_to"`
	CreatedBy    string       `json:"created_by"    bson:"created_by"`
	Status       string       `json:"status"        bson:"status"`
	Priority     string       `json:"priority"      bson:"priority"`
	Labels       []string     `json:"labels"        bson:"labels"`
	DueDate      string       `json:"due_date"      bson:"due_date"`
	TimeTracking TimeTracking `json:"time_tracking" bson:"time_tracking"`
	CreatedAt    time.Time    `json:"created_at"    bson:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"    bson:"updated_at"`
}

func NewTask(id, title, description, projectID, assignedTo, createdBy, status, priority, dueDate string, labels []string, estimated, logged float64) *Task {
	return &Task{
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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *Task) GetID() string {
	return u.ID
}
