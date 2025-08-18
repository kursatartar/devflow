package services

import (
	"devflow/internal/interfaces"
	"devflow/internal/models"
	"time"

	"github.com/google/uuid"
)

type TeamManager struct{}

func NewTeamService() interfaces.TeamService {
	return &TeamManager{}
}

func (s *TeamManager) CreateTeam(id, name, description, ownerID string, members []models.TeamMember, settings models.TeamSettings) (*models.Team, error) {
	if id == "" {
		id = uuid.NewString()
	}
	if _, exists := models.Teams[id]; exists {
		return nil, ErrTeamExists
	}
	now := time.Now()
	out := make([]models.TeamMember, 0, len(members))
	for _, m := range members {
		if m.JoinedAt.IsZero() {
			m.JoinedAt = now
		}
		out = append(out, m)
	}
	team := models.NewTeam(id, name, description, ownerID, out, settings)
	models.Teams[id] = team
	return team, nil
}

func (s *TeamManager) GetTeam(id string) (*models.Team, error) {
	t, ok := models.Teams[id]
	if !ok {
		return nil, ErrTeamNotFound
	}
	return t, nil
}

func (s *TeamManager) UpdateTeam(id string, name, description *string, settings *models.TeamSettings) (*models.Team, error) {
	t, ok := models.Teams[id]
	if !ok {
		return nil, ErrTeamNotFound
	}
	if name != nil {
		t.Name = *name
	}
	if description != nil {
		t.Description = *description
	}
	if settings != nil {
		t.Settings = *settings
	}
	t.UpdatedAt = time.Now()
	return t, nil
}

func (s *TeamManager) AddMember(teamID, userID, role string) (*models.Team, error) {
	t, ok := models.Teams[teamID]
	if !ok {
		return nil, ErrTeamNotFound
	}
	for i := range t.Members {
		if t.Members[i].UserID == userID {
			t.Members[i].Role = role
			t.Members[i].JoinedAt = time.Now()
			t.UpdatedAt = time.Now()
			return t, nil
		}
	}
	t.Members = append(t.Members, models.TeamMember{UserID: userID, Role: role, JoinedAt: time.Now()})
	t.UpdatedAt = time.Now()
	return t, nil
}

func (s *TeamManager) RemoveMember(teamID, userID string) (*models.Team, error) {
	t, ok := models.Teams[teamID]
	if !ok {
		return nil, ErrTeamNotFound
	}
	dst := t.Members[:0]
	for _, m := range t.Members {
		if m.UserID != userID {
			dst = append(dst, m)
		}
	}
	t.Members = dst
	t.UpdatedAt = time.Now()
	return t, nil
}

func (s *TeamManager) ChangeMemberRole(teamID, userID, role string) (*models.Team, error) {
	t, ok := models.Teams[teamID]
	if !ok {
		return nil, ErrTeamNotFound
	}
	for i := range t.Members {
		if t.Members[i].UserID == userID {
			t.Members[i].Role = role
			t.UpdatedAt = time.Now()
			return t, nil
		}
	}
	t.UpdatedAt = time.Now()
	return t, nil
}

func (s *TeamManager) ListTeams() []*models.Team {
	out := make([]*models.Team, 0, len(models.Teams))
	for _, t := range models.Teams {
		out = append(out, t)
	}
	return out
}

func (s *TeamManager) FilterTeamsByOwner(ownerID string) []*models.Team {
	var out []*models.Team
	for _, t := range models.Teams {
		if t.OwnerID == ownerID {
			out = append(out, t)
		}
	}
	return out
}

func (s *TeamManager) DeleteTeam(id string) error {
	if _, ok := models.Teams[id]; !ok {
		return ErrTeamNotFound
	}
	delete(models.Teams, id)
	return nil
}
