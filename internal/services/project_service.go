package services

import (
	"context"
	"devflow/internal/interfaces"
	"devflow/internal/models"
)

type ProjectManager struct {
	repo interfaces.ProjectRepository
}

func NewProjectService(repo interfaces.ProjectRepository) *ProjectManager {
    return &ProjectManager{repo: repo}
}

func (p *ProjectManager) CreateProject(id, name, description, ownerID, teamID, status string, teamMembers []string, isPrivate bool, taskWorkflow []string) (*models.Project, error) {
    pr := models.NewProject(id, name, description, ownerID, teamID, teamMembers, status, isPrivate, taskWorkflow)
	_, err := p.repo.Create(context.Background(), pr)
	if err != nil {
		return nil, err
	}
	return pr, nil
}

func (p *ProjectManager) ListProjects() []*models.Project {
	out, _ := p.repo.List(context.Background())
	return out
}

func (p *ProjectManager) GetProject(id string) (*models.Project, error) {
	out, err := p.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrProjectNotFound
	}
	return out, nil
}

func (p *ProjectManager) FilterProjectsByOwner(ownerID string) []*models.Project {
	out, _ := p.repo.FilterByOwner(context.Background(), ownerID)
	return out
}

func (p *ProjectManager) UpdateProject(id string, name string, description string, status string, teamID string, teamMembers []string, isPrivate bool, taskWorkflow []string) (*models.Project, error) {
	namePtr := &name
	descPtr := &description
	statusPtr := &status
	privatePtr := &isPrivate
	workflowPtr := &taskWorkflow
	var ownerPtr *string = nil
    teamIDPtr := &teamID

    if err := p.repo.UpdateFields(context.Background(), id, namePtr, descPtr, statusPtr, privatePtr, workflowPtr, ownerPtr, teamIDPtr); err != nil {
		return nil, err
	}
	out, err := p.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrProjectNotFound
	}
	return out, nil
}

func (p *ProjectManager) DeleteProject(id string) error {
	return p.repo.Delete(context.Background(), id)
}

var _ interfaces.ProjectService = (*ProjectManager)(nil)
