package interfaces

import "devflow/internal/models"

type ProjectService interface {
	CreateProject(id, name, description, ownerID, status string, teamMembers []string, isPrivate bool, taskWorkflow []string) (*models.Project, error)
	UpdateProject(id, name, description, status string, teamMembers []string, isPrivate bool, taskWorkflow []string) (*models.Project, error)
	DeleteProject(id string) error
	ListProjects() []*models.Project
	FilterProjectsByOwner(ownerID string) []*models.Project
	GetProject(id string) (*models.Project, error)
}
