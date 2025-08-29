package services

import (
	"context"
	"devflow/internal/interfaces"
	"devflow/internal/models"
)

type TeamManager struct {
	repo interfaces.TeamRepository
}

func NewTeamService(repo interfaces.TeamRepository) interfaces.TeamService {
	return &TeamManager{repo}
}

func (t *TeamManager) CreateTeam(
	id, name, description, ownerID string,
	members []models.TeamMember,
	settings models.TeamSettings,
) (*models.Team, error) {
	team := models.NewTeam(id, name, description, ownerID, members, settings)
	_, err := t.repo.Create(context.Background(), team)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func (t *TeamManager) ListTeams() []*models.Team {
	out, _ := t.repo.List(context.Background())
	return out
}

func (t *TeamManager) GetTeam(id string) (*models.Team, error) {
	out, err := t.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrTeamNotFound
	}
	return out, nil
}

func (t *TeamManager) UpdateTeam(
	id string,
	name, description *string,
	settings *models.TeamSettings,
) (*models.Team, error) {
	if err := t.repo.UpdateFields(context.Background(), id, name, description, settings); err != nil {
		return nil, err
	}
	out, err := t.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrTeamNotFound
	}
	return out, nil
}

func (t *TeamManager) DeleteTeam(id string) error {
	return t.repo.Delete(context.Background(), id)
}

func (t *TeamManager) AddMember(teamID, userID, role string) (*models.Team, error) {
	if err := t.repo.AddMember(context.Background(), teamID, userID, role); err != nil {
		return nil, err
	}
	out, err := t.repo.GetByID(context.Background(), teamID)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrTeamNotFound
	}
	return out, nil
}

func (t *TeamManager) RemoveMember(teamID, userID string) (*models.Team, error) {
	if err := t.repo.RemoveMember(context.Background(), teamID, userID); err != nil {
		return nil, err
	}
	out, err := t.repo.GetByID(context.Background(), teamID)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrTeamNotFound
	}
	return out, nil
}

func (t *TeamManager) ChangeMemberRole(teamID, userID, role string) (*models.Team, error) {
	if err := t.repo.ChangeMemberRole(context.Background(), teamID, userID, role); err != nil {
		return nil, err
	}
	out, err := t.repo.GetByID(context.Background(), teamID)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrTeamNotFound
	}
	return out, nil
}

func (t *TeamManager) FilterTeamsByOwner(ownerID string) []*models.Team {
	out, _ := t.repo.FilterByOwner(context.Background(), ownerID)
	return out
}

var _ interfaces.TeamService = (*TeamManager)(nil)
