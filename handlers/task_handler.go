package handlers

import (
	"fmt"
	"github.com/kursatartar/devflowv2/models"
)

func CreateTask(id, title, status, projectID string) {

	if _, ok := models.Projects[projectID]; !ok {
		fmt.Println("cannot create task — project not found:", projectID)
		return
	}
	task := models.Task{
		ID:        id,
		Title:     title,
		Status:    status,
		ProjectID: projectID,
	}
	models.Tasks[id] = task
	fmt.Println("task created:", task)
}

func ListTasks() {
	fmt.Println("all tasks:")
	for _, task := range models.Tasks {
		fmt.Printf("- id: %s, title: %s, status: %s, projectID: %s\n",
			task.ID, task.Title, task.Status, task.ProjectID)
	}
}

func UpdateTask(id, newTitle, newStatus, newProjectID string) {
	if _, ok := models.Tasks[id]; ok {
		task := models.Tasks[id]
		task.Title = newTitle
		task.Status = newStatus
		task.ProjectID = newProjectID
		models.Tasks[id] = task
		fmt.Println("task updated:", task)
	} else {
		fmt.Println("task not found")
	}
}

func DeleteTask(id string) {
	if _, ok := models.Tasks[id]; ok {
		delete(models.Tasks, id)
		fmt.Println("task deleted:", id)
	} else {
		fmt.Println("tssk not found")
	}
}
