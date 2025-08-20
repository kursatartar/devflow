package interfaces

import "devflow/internal/models"

type TeamService interface {
	CreateTeam(id, name, description, ownerID string, members []models.TeamMember, settings models.TeamSettings) (*models.Team, error)
	GetTeam(id string) (*models.Team, error)
	UpdateTeam(id string, name, description *string, settings *models.TeamSettings) (*models.Team, error)
	AddMember(teamID, userID, role string) (*models.Team, error)
	RemoveMember(teamID, userID string) (*models.Team, error)
	ChangeMemberRole(teamID, userID, role string) (*models.Team, error)
	ListTeams() []*models.Team
	FilterTeamsByOwner(ownerID string) []*models.Team
	DeleteTeam(id string) error
}
