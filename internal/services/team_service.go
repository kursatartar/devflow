package services

import (
	"devflow/internal/interfaces"
	"devflow/internal/models"
	"fmt"
	"time"
)

type TeamManager struct{}

func (s *TeamManager) CreateTeam(id, name, description, ownerID string, members []models.TeamMember, settings models.TeamSettings) error {
	if _, exists := models.Teams[id]; exists {
		return ErrTeamExists
	}
	team := models.NewTeam(id, name, description, ownerID, members, settings)
	models.Teams[id] = team
	fmt.Println("team created:", id)
	return nil
}

func (s *TeamManager) UpdateTeam(id, newName, newDescription string, newSettings models.TeamSettings) {
	if t, ok := models.Teams[id]; ok {
		if newName != "" {
			t.Name = newName
		}
		if newDescription != "" {
			t.Description = newDescription
		}
		t.Settings = newSettings
		t.UpdatedAt = time.Now()
		fmt.Println("team updated:", id)
		return
	}
	fmt.Println(ErrTeamNotFound.Error())
}

func (s *TeamManager) AddMember(teamID, userID, role string) {
	if t, ok := models.Teams[teamID]; ok {

		for i := range t.Members {
			if t.Members[i].UserID == userID {
				t.Members[i].Role = role
				t.Members[i].JoinedAt = time.Now()
				t.UpdatedAt = time.Now()
				fmt.Println("team member role updated:", teamID, userID)
				return
			}
		}
		t.Members = append(t.Members, models.TeamMember{UserID: userID, Role: role, JoinedAt: time.Now()})
		t.UpdatedAt = time.Now()
		fmt.Println("team member added:", teamID, userID)
		return
	}
	fmt.Println(ErrTeamNotFound.Error())
}

func (s *TeamManager) RemoveMember(teamID, userID string) {
	if t, ok := models.Teams[teamID]; ok {
		filtered := t.Members[:0]
		for _, m := range t.Members {
			if m.UserID != userID {
				filtered = append(filtered, m)
			}
		}
		t.Members = filtered
		t.UpdatedAt = time.Now()
		fmt.Println("team member removed:", teamID, userID)
		return
	}
	fmt.Println(ErrTeamNotFound.Error())
}

func (s *TeamManager) ChangeMemberRole(teamID, userID, newRole string) {
	if t, ok := models.Teams[teamID]; ok {
		for i := range t.Members {
			if t.Members[i].UserID == userID {
				t.Members[i].Role = newRole
				t.UpdatedAt = time.Now()
				fmt.Println("team member role changed:", teamID, userID)
				return
			}
		}
		fmt.Println("member not found in team")
		return
	}
	fmt.Println(ErrTeamNotFound.Error())
}

func (s *TeamManager) ListTeams() {
	for id, t := range models.Teams {
		fmt.Printf("- %s: %s (members:%d)\n", id, t.Name, len(t.Members))
	}
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

func (s *TeamManager) DeleteTeam(id string) {
	if _, ok := models.Teams[id]; ok {
		delete(models.Teams, id)
		fmt.Println("team deleted:", id)
		return
	}
	fmt.Println(ErrTeamNotFound.Error())
}
func NewTeamService() interfaces.TeamService {
	return &TeamManager{}
}
