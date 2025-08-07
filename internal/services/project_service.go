package services

import (
	"devflow/internal/interfaces"
	"devflow/internal/models"
	"fmt"
	"time"
)

type ProjectManager struct{}

func NewProjectService() interfaces.ProjectService {
	return &ProjectManager{}
}

func (p *ProjectManager) CreateProject(id, name, description, ownerID, status string, teamMembers []string, isPrivate bool, taskWorkflow []string) error {
	if _, exists := models.Projects[id]; exists {
		return ErrProjectExists
	}

	project := models.NewProject(id, name, description, ownerID, teamMembers, status, isPrivate, taskWorkflow)
	models.Projects[id] = project
	return nil
}

func (p *ProjectManager) UpdateProject(
	id, name, description, status string,
	teamMembers []string,
	isPrivate bool,
	taskWorkflow []string,
) {
	project, exists := models.Projects[id]
	if !exists {
		fmt.Println("project not found:", id)
		return
	}

	project.Name = name
	project.Description = description
	project.Status = status
	project.TeamMembers = teamMembers
	project.Settings.IsPrivate = isPrivate
	project.Settings.TaskWorkflow = taskWorkflow
	project.UpdatedAt = time.Now()

	fmt.Println("project updated:", project.Name)
}

func (p *ProjectManager) DeleteProject(id string) {
	if _, exists := models.Projects[id]; exists {
		delete(models.Projects, id)
		fmt.Println("project deleted:", id)
	} else {
		fmt.Println("project not found:", id)
	}
}

func (p *ProjectManager) ListProjects() {
	fmt.Println("all projects:")
	for id, project := range models.Projects {
		fmt.Printf("- %s: %s (%s)\n", id, project.Name, project.Status)
	}
}

func (p *ProjectManager) FilterProjectsByOwner(ownerID string) []*models.Project {
	var filtered []*models.Project
	for _, project := range models.Projects {
		if project.OwnerID == ownerID {
			filtered = append(filtered, project)
		}
	}
	return filtered
}
