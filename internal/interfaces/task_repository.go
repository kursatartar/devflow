package interfaces

import (
	"context"
	"devflow/internal/models"
)

type TaskRepository interface {
	Create(ctx context.Context, t *models.Task) (string, error)
	GetByID(ctx context.Context, id string) (*models.Task, error)
	List(ctx context.Context) ([]*models.Task, error)
	FilterByProject(ctx context.Context, projectID string) ([]*models.Task, error)
	UpdateFields(ctx context.Context, id string, title, description, status, priority, dueDate *string, labels *[]string, estimated, logged *float64) error
	Delete(ctx context.Context, id string) error
}
