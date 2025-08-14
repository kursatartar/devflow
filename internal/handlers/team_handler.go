package handlers

import (
	"devflow/internal/models"
	"devflow/internal/services"
)

var teamService = services.NewTeamService()

func CreateTeam(id, name, description, ownerID string, members []models.TeamMember, settings models.TeamSettings) error {
	return teamService.CreateTeam(id, name, description, ownerID, members, settings)
}
func UpdateTeam(id, newName, newDescription string, newSettings models.TeamSettings) {
	teamService.UpdateTeam(id, newName, newDescription, newSettings)
}
func AddTeamMember(teamID, userID, role string) {
	teamService.AddMember(teamID, userID, role)
}
func RemoveTeamMember(teamID, userID string) {
	teamService.RemoveMember(teamID, userID)
}
func ChangeTeamMemberRole(teamID, userID, role string) {
	teamService.ChangeMemberRole(teamID, userID, role)
}
func ListTeams() {
	teamService.ListTeams()
}
func FilterTeamsByOwner(ownerID string) {
	teamService.FilterTeamsByOwner(ownerID)
}
func DeleteTeam(id string) {
	teamService.DeleteTeam(id)
}
