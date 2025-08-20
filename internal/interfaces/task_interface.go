package interfaces

import "devflow/internal/models"

type TaskService interface {
	CreateTask(id, title, description, projectID, assignedTo, createdBy, status, priority, dueDate string, labels []string, estimated, logged float64) (*models.Task, error)
	UpdateTask(id string, title, description, status, priority, dueDate *string, labels *[]string, estimated, logged *float64) (*models.Task, error)
	DeleteTask(id string) error
	ListTasks() []*models.Task
	FilterTasksByProject(projectID string) []*models.Task
	GetTask(id string) (*models.Task, error)
}
