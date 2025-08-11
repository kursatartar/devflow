package services

import (
	"devflow/internal/interfaces"
	"devflow/internal/models"
	"fmt"
	"time"
)

type TaskManager struct{}

func (t TaskManager) CreateTask(id, title, description, projectID, assignedTo, createdBy, status, priority, dueDate string, labels []string, estimated, logged float64) error {
	if _, exists := models.Tasks[id]; exists {
		return ErrTaskExists
	}

	_, err := time.Parse(time.RFC3339, dueDate)
	if err != nil {
		return ErrInvalidDueDate
	}

	task := models.NewTask(id, title, description, projectID, assignedTo, createdBy, status, priority, dueDate, labels, estimated, logged)
	models.Tasks[id] = task

	return nil
}

func (t TaskManager) UpdateTask(id, newTitle, newDescription, newStatus, newPriority, newDueDate string, newLabels []string, newEstimated, newLogged float64) {
	if task, exists := models.Tasks[id]; exists {
		task.Title = newTitle
		task.Description = newDescription
		task.Status = newStatus
		task.Priority = newPriority
		task.DueDate = newDueDate
		task.Labels = newLabels
		models.Tasks[id] = task
		fmt.Println("task updated:", task)
	}
}

func (t TaskManager) DeleteTask(id string) {
	if task, exists := models.Tasks[id]; exists {
		delete(models.Tasks, id)
		fmt.Println("task deleted:", task)
	}
}

func (t TaskManager) ListTasks() {
	fmt.Println("task list")
	for _, task := range models.Tasks {
		fmt.Println(task)
	}
}

func (t TaskManager) FilterTasksByProject(projectID string) []*models.Task {
	var tasks []*models.Task
	for _, task := range models.Tasks {
		if task.ProjectID == projectID {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func NewTaskService() interfaces.TaskService {
	return &TaskManager{}
}
