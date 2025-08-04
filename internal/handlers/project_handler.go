package handlers

import (
	"devflow/internal/models"
	"fmt"
)

func CreateProject(
	id, name, description, ownerID, status string,
	teamMembers []string, isPrivate bool, taskWorkflow []string,
) {
	project := models.NewProject(id, name, description, ownerID, teamMembers, status, isPrivate, taskWorkflow)
	models.Projects[id] = project
	fmt.Println("project created:", project.Name)
}

func ListProjects() {
	fmt.Println("all projects:")
	for id, project := range models.Projects {
		fmt.Printf("- %s: %s (%s)\n", id, project.Name, project.Status)
	}
}

func UpdateProject(
	id, newName, newDescription, newStatus string,
	newTeam []string, newIsPrivate bool, newWorkflow []string,
) {
	if project, exists := models.Projects[id]; exists {
		project.Name = newName
		project.Description = newDescription
		project.Status = newStatus
		project.TeamMembers = newTeam
		project.Settings.IsPrivate = newIsPrivate
		project.Settings.TaskWorkflow = newWorkflow
		project.UpdatedAt = models.NewProject(id, "", "", "", nil, "", false, nil).UpdatedAt

		models.Projects[id] = project
		fmt.Println("project updated:", project.Name)
	} else {
		fmt.Println("project not found")
	}
}

func DeleteProject(id string) {
	if _, exists := models.Projects[id]; exists {
		delete(models.Projects, id)
		fmt.Println("project deleted:", id)
	} else {
		fmt.Println("project not found")
	}
}
