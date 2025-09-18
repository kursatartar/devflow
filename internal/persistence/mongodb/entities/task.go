package entities

import (
    "devflow/internal/models"
    "time"
)

type TimeTrackingEntity struct {
    EstimatedHours float64 `bson:"estimated_hours"`
    LoggedHours    float64 `bson:"logged_hours"`
}

type TaskEntity struct {
    ID           string             `bson:"id"`
    Title        string             `bson:"title"`
    Description  string             `bson:"description"`
    ProjectID    string             `bson:"project_id"`
    AssignedTo   string             `bson:"assigned_to"`
    CreatedBy    string             `bson:"created_by"`
    Status       string             `bson:"status"`
    Priority     string             `bson:"priority"`
    Labels       []string           `bson:"labels"`
    DueDate      string             `bson:"due_date"`
    TimeTracking TimeTrackingEntity `bson:"time_tracking"`
    CreatedAt    time.Time          `bson:"created_at"`
    UpdatedAt    time.Time          `bson:"updated_at"`
}

func TaskFromModel(m *models.Task) *TaskEntity {
    if m == nil {
        return nil
    }
    return &TaskEntity{
        ID:          m.ID,
        Title:       m.Title,
        Description: m.Description,
        ProjectID:   m.ProjectID,
        AssignedTo:  m.AssignedTo,
        CreatedBy:   m.CreatedBy,
        Status:      m.Status,
        Priority:    m.Priority,
        Labels:      m.Labels,
        DueDate:     m.DueDate,
        TimeTracking: TimeTrackingEntity{
            EstimatedHours: m.TimeTracking.EstimatedHours,
            LoggedHours:    m.TimeTracking.LoggedHours,
        },
        CreatedAt: m.CreatedAt,
        UpdatedAt: m.UpdatedAt,
    }
}

func (e *TaskEntity) ToModel() *models.Task {
    if e == nil {
        return nil
    }
    return &models.Task{
        ID:          e.ID,
        Title:       e.Title,
        Description: e.Description,
        ProjectID:   e.ProjectID,
        AssignedTo:  e.AssignedTo,
        CreatedBy:   e.CreatedBy,
        Status:      e.Status,
        Priority:    e.Priority,
        Labels:      e.Labels,
        DueDate:     e.DueDate,
        TimeTracking: models.TimeTracking{
            EstimatedHours: e.TimeTracking.EstimatedHours,
            LoggedHours:    e.TimeTracking.LoggedHours,
        },
        CreatedAt: e.CreatedAt,
        UpdatedAt: e.UpdatedAt,
    }
}


