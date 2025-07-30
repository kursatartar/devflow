package models

import "time"

type ProjectSettings struct {
	IsPrivate    bool
	TaskWorkflow []string
}

type Project struct {
	ID          string
	Name        string
	Description string
	OwnerID     string
	TeamMembers []string
	Status      string
	Settings    ProjectSettings
	CreatedAt   string
	UpdatedAt   string
}

var Projects = map[string]Project{}

func NewProject(id, name, description, ownerID string, teamMembers []string, status string, isPrivate bool, taskWorkflow []string) Project {
	return Project{
		ID:          id,
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
		TeamMembers: teamMembers,
		Status:      status,
		Settings: ProjectSettings{
			IsPrivate:    isPrivate,
			TaskWorkflow: taskWorkflow,
		},
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
}

// type Project struct {
//   ID          string   // `json:"id"`
//   Name        string   // `json:"name"`
//   Description string   // `json:"description"`
//   OwnerID     string   // `json:"owner_id"`
//   TeamMembers []string // `json:"team_members"`
//   Status      string   // `json:"status"`
//   Settings    Settings // `json:"settings"`
//   CreatedAt   string   // `json:"created_at"`
//   UpdatedAt   string   // `json:"updated_at"`
// }

// type Settings struct {
//   IsPrivate    bool     // `json:"is_private"`
//   TaskWorkflow []string // `json:"task_workflow"`
// }
