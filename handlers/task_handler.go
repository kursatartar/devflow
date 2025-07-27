package handlers

import (
	"fmt"
	"github.com/kursatartar/devflowv2/models"
)

func CreateTask(
	id, title, description, projectID, assignedTo, createdBy,
	status, priority, dueDate string,
	labels []string, estimated, logged float64,
) {
	task := models.NewTask(id, title, description, projectID, assignedTo, createdBy, status, priority, dueDate, labels, estimated, logged)
	models.Tasks[id] = task
	fmt.Println("task created:", task.Title)
}

func ListTasks() {
	fmt.Println("all tasks:")
	for id, task := range models.Tasks {
		fmt.Printf("- %s: %s (%s)\n", id, task.Title, task.Status)
	}
}

func UpdateTask(
	id, newTitle, newDescription, newStatus,
	newPriority, newDueDate string,
	newLabels []string, newEstimated, newLogged float64,
) {
	if task, exists := models.Tasks[id]; exists {
		task.Title = newTitle
		task.Description = newDescription
		task.Status = newStatus
		task.Priority = newPriority
		task.DueDate = newDueDate
		task.Labels = newLabels
		task.TimeTracking.EstimatedHours = newEstimated
		task.TimeTracking.LoggedHours = newLogged
		task.UpdatedAt = models.NewTask(id, "", "", "", "", "", "", "", "", nil, 0, 0).UpdatedAt

		models.Tasks[id] = task
		fmt.Println("task updated:", task.Title)
	} else {
		fmt.Println("task not found!")
	}
}

func DeleteTask(id string) {
	if _, exists := models.Tasks[id]; exists {
		delete(models.Tasks, id)
		fmt.Println("task deleted:", id)
	} else {
		fmt.Println("task not found")
	}
}
