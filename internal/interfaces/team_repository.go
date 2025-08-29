package interfaces

import (
	"context"
	"devflow/internal/models"
)

type TeamRepository interface {
	Create(ctx context.Context, t *models.Team) (string, error)
	GetByID(ctx context.Context, id string) (*models.Team, error)
	List(ctx context.Context) ([]*models.Team, error)
	UpdateFields(ctx context.Context, id string, name, description *string, settings *models.TeamSettings) error
	AddMember(ctx context.Context, teamID, userID, role string) error
	RemoveMember(ctx context.Context, teamID, userID string) error
	ChangeMemberRole(ctx context.Context, teamID, userID, role string) error
	Delete(ctx context.Context, id string) error
	FilterByOwner(ctx context.Context, ownerID string) ([]*models.Team, error)
}
