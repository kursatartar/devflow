package models

import "time"

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
	DueDate      time.Time
	TimeTracking TimeTracking
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type TimeTracking struct {
	EstimatedHours float64
	LoggedHours    float64
}

var Tasks = map[string]Task{}

func NewTask(
	id, title, description, projectID, assignedTo, createdBy, status, priority string,
	dueDate time.Time, labels []string, estimated, logged float64,
) Task {
	now := time.Now().UTC()
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
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (t *Task) Update(
	newTitle, newDescription, newStatus, newPriority string,
	newDueDate time.Time, newLabels []string,
	newEstimated, newLogged float64,
) {
	t.Title = newTitle
	t.Description = newDescription
	t.Status = newStatus
	t.Priority = newPriority
	t.DueDate = newDueDate
	t.Labels = newLabels
	t.TimeTracking.EstimatedHours = newEstimated
	t.TimeTracking.LoggedHours = newLogged
	t.UpdatedAt = time.Now().UTC()
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
