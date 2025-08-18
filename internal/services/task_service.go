package services

import (
	"devflow/internal/interfaces"
	"devflow/internal/models"
	"time"

	"github.com/google/uuid"
)

type TaskManager struct{}

func NewTaskService() interfaces.TaskService {
	return &TaskManager{}
}

func (t TaskManager) CreateTask(id, title, description, projectID, assignedTo, createdBy, status, priority, dueDate string, labels []string, estimated, logged float64) (*models.Task, error) {
	if id == "" {
		id = uuid.NewString()
	}
	if _, exists := models.Tasks[id]; exists {
		return nil, ErrTaskExists
	}
	if _, err := time.Parse(time.RFC3339, dueDate); err != nil {
		return nil, ErrInvalidDueDate
	}
	task := models.NewTask(id, title, description, projectID, assignedTo, createdBy, status, priority, dueDate, labels, estimated, logged)
	models.Tasks[id] = task
	return task, nil
}

func (t TaskManager) UpdateTask(id string, title, description, status, priority, dueDate *string, labels *[]string, estimated, logged *float64) (*models.Task, error) {
	task, ok := models.Tasks[id]
	if !ok {
		return nil, ErrTaskNotFound
	}
	if title != nil {
		task.Title = *title
	}
	if description != nil {
		task.Description = *description
	}
	if status != nil {
		task.Status = *status
	}
	if priority != nil {
		task.Priority = *priority
	}
	if dueDate != nil {
		if _, err := time.Parse(time.RFC3339, *dueDate); err != nil {
			return nil, ErrInvalidDueDate
		}
		task.DueDate = *dueDate
	}
	if labels != nil {
		task.Labels = *labels
	}
	if estimated != nil {
		task.TimeTracking.EstimatedHours = *estimated
	}
	if logged != nil {
		task.TimeTracking.LoggedHours = *logged
	}
	task.UpdatedAt = time.Now()
	return task, nil
}

func (t TaskManager) DeleteTask(id string) error {
	if _, exists := models.Tasks[id]; !exists {
		return ErrTaskNotFound
	}
	delete(models.Tasks, id)
	return nil
}

func (t TaskManager) ListTasks() []*models.Task {
	out := make([]*models.Task, 0, len(models.Tasks))
	for _, task := range models.Tasks {
		out = append(out, task)
	}
	return out
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
