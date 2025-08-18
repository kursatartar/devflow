package services

import (
	"devflow/internal/models"
	"errors"
	"time"
)

type ProjectManager struct{}

func NewProjectService() *ProjectManager {
	return &ProjectManager{}
}

func (s *ProjectManager) CreateProject(id, name, description, ownerID, status string, teamMembers []string, isPrivate bool, taskWorkflow []string) (*models.Project, error) {
	if _, exists := models.Projects[id]; exists {
		return nil, errors.New("project already exists")
	}
	project := models.NewProject(id, name, description, ownerID, teamMembers, status, isPrivate, taskWorkflow)
	models.Projects[id] = project
	return project, nil
}

func (s *ProjectManager) ListProjects() []*models.Project {
	out := make([]*models.Project, 0, len(models.Projects))
	for _, p := range models.Projects {
		out = append(out, p)
	}
	return out
}

func (s *ProjectManager) UpdateProject(id, name, description, status string, teamMembers []string, isPrivate bool, taskWorkflow []string) (*models.Project, error) {
	project, ok := models.Projects[id]
	if !ok {
		return nil, ErrProjectNotFound
	}
	project.Name = name
	project.Description = description
	project.Status = status
	project.TeamMembers = teamMembers
	project.Settings.IsPrivate = isPrivate
	project.Settings.TaskWorkflow = taskWorkflow
	project.UpdatedAt = time.Now()
	return project, nil
}

func (s *ProjectManager) DeleteProject(id string) error {
	if _, ok := models.Projects[id]; !ok {
		return ErrProjectNotFound
	}
	delete(models.Projects, id)
	return nil
}

func (s *ProjectManager) FilterProjectsByOwner(ownerID string) []*models.Project {
	var filtered []*models.Project
	for _, p := range models.Projects {
		if p.OwnerID == ownerID {
			filtered = append(filtered, p)
		}
	}
	return filtered
}
