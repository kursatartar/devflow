package interfaces

import (
	"context"
	"devflow/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, u *models.User) (string, error)

	GetByID(ctx context.Context, id string) (*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	List(ctx context.Context) ([]*models.User, error)
	FilterByRole(ctx context.Context, role string) ([]*models.User, error)

	UpdateProfile(ctx context.Context, id string, p models.Profile) error
	UpdateCore(ctx context.Context, id, username, email, passwordHash, role string) error

	Delete(ctx context.Context, id string) error
}
