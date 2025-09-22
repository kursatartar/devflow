package services

import (
	"context"
	"devflow/internal/interfaces"
	"devflow/internal/models"
	"time"
)

type TaskManager struct {
	repo interfaces.TaskRepository
}

func NewTaskService(repo interfaces.TaskRepository) *TaskManager {
    return &TaskManager{repo}
}

func (t *TaskManager) CreateTask(id, title, description, projectID, assignedTo, createdBy, status, priority, dueDate string, labels []string, estimated, logged float64) (*models.Task, error) {
	if _, err := time.Parse(time.RFC3339, dueDate); err != nil {
		return nil, ErrInvalidDueDate
	}
	task := models.NewTask(id, title, description, projectID, assignedTo, createdBy, status, priority, dueDate, labels, estimated, logged)
	_, err := t.repo.Create(context.Background(), task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskManager) UpdateTask(id string, title, description, status, priority, dueDate *string, labels *[]string, estimated, logged *float64) (*models.Task, error) {
	if dueDate != nil {
		if _, err := time.Parse(time.RFC3339, *dueDate); err != nil {
			return nil, ErrInvalidDueDate
		}
	}
	if err := t.repo.UpdateFields(context.Background(), id, title, description, status, priority, dueDate, labels, estimated, logged); err != nil {
		return nil, err
	}
	out, err := t.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrTaskNotFound
	}
	return out, nil
}

func (t *TaskManager) DeleteTask(id string) error {
	return t.repo.Delete(context.Background(), id)
}

func (t *TaskManager) ListTasks() []*models.Task {
	out, _ := t.repo.List(context.Background())
	return out
}

func (t *TaskManager) FilterTasksByProject(projectID string) []*models.Task {
	out, _ := t.repo.FilterByProject(context.Background(), projectID)
	return out
}

func (t *TaskManager) GetTask(id string) (*models.Task, error) {
	out, err := t.repo.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrTaskNotFound
	}
	return out, nil
}
