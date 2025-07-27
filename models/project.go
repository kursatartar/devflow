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
