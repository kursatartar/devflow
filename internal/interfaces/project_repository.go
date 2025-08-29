package interfaces

import (
	"context"
	"devflow/internal/models"
)

type ProjectRepository interface {
	Create(ctx context.Context, p *models.Project) (string, error)
	GetByID(ctx context.Context, id string) (*models.Project, error)
	List(ctx context.Context) ([]*models.Project, error)
	FilterByOwner(ctx context.Context, ownerID string) ([]*models.Project, error)
	UpdateFields(ctx context.Context, id string, name, description, status *string, teamMembers *[]string, isPrivate *bool, taskWorkflow *[]string, ownerID *string) error
	Delete(ctx context.Context, id string) error
}
