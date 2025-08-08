package models

import "time"

type ProjectSettings struct {
	IsPrivate    bool
	TaskWorkflow []string
}

type Project struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	OwnerID     string          `json:"owner_id"`
	TeamMembers []string        `json:"team_members"`
	Status      string          `json:"status"`
	Settings    ProjectSettings `json:"settings"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

var Projects = map[string]*Project{}

func NewProject(id, name, description, ownerID string, teamMembers []string, status string, isPrivate bool, taskWorkflow []string) *Project {
	return &Project{
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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
