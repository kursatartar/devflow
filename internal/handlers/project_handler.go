package handlers

import (
	"devflow/internal/services"
	"fmt"
)

var projectService = services.NewProjectService()

func CreateProject(id, name, description, ownerID, status string, teamMembers []string, isPrivate bool, taskWorkflow []string) {
	err := projectService.CreateProject(id, name, description, ownerID, status, teamMembers, isPrivate, taskWorkflow)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("project created")
}

func ListProjects() {
	projectService.ListProjects()
}

func UpdateProject(id, newName, newDescription, newStatus string, newTeamMembers []string, newIsPrivate bool, newWorkflow []string) {
	projectService.UpdateProject(id, newName, newDescription, newStatus, newTeamMembers, newIsPrivate, newWorkflow)
}

func DeleteProject(id string) {
	projectService.DeleteProject(id)
}
