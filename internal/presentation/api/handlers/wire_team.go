package handlers

import "devflow/internal/interfaces"

var teamService interfaces.TeamService

func InitTeamService(s interfaces.TeamService) {
	teamService = s
}
