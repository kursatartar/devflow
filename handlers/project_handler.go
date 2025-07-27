package handlers

import (
	"fmt"
	"github.com/kursatartar/devflowv2/models"
)

func CreateProject(id, name, description string) {
	project := models.Project{
		ID:          id,
		Name:        name,
		Description: description,
	}
	models.Projects[id] = project
	fmt.Println("project created:", project)
}

func ListProjects() {
	fmt.Println("all projects:")
	for _, project := range models.Projects {
		fmt.Printf("- id: %s, name: %s, description: %s\n", project.ID, project.Name, project.Description)
	}
}

func UpdateProject(id, newName, newDescription string) {
	if _, exists := models.Projects[id]; exists {
		project := models.Projects[id]
		project.Name = newName
		project.Description = newDescription
		models.Projects[id] = project
		fmt.Println("poject updated:", project)
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
