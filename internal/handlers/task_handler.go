package handlers

import (
	"devflow/internal/services"
	"fmt"
)

const (
	StatusPending = "pending"
	StatusActive  = "active"
	StatusDone    = "done"
)

var taskService = services.NewTaskService()

func CreateTask(id, title, description, projectID, assignedTo, createdBy, status, priority, dueDate string, labels []string, estimated, logged float64) {
	err := taskService.CreateTask(id, title, description, projectID, assignedTo, createdBy, status, priority, dueDate, labels, estimated, logged)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("task created")

}

func ListTasks() {
	taskService.ListTasks()
}

func UpdateTask(id, newTitle, newDescription, newStatus, newPriority, newDueDate string, newLabels []string, newEstimated, newLogged float64) {
	taskService.UpdateTask(id, newTitle, newDescription, newStatus, newPriority, newDueDate, newLabels, newEstimated, newLogged)
}

func DeleteTask(id string) {
	taskService.DeleteTask(id)
}
