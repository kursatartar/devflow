package entities

import (
    "devflow/internal/models"
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type TimeTrackingEntity struct {
    EstimatedHours float64 `bson:"estimated_hours"`
    LoggedHours    float64 `bson:"logged_hours"`
}

type TaskEntity struct {
    ID           string             `bson:"_id"`
    Title        string             `bson:"title"`
    Description  string             `bson:"description"`
    ProjectID    primitive.ObjectID `bson:"project_id"`
    AssignedTo   primitive.ObjectID `bson:"assigned_to"`
    CreatedBy    primitive.ObjectID `bson:"created_by"`
    Status       string             `bson:"status"`
    Priority     string             `bson:"priority"`
    Labels       []string           `bson:"labels"`
    DueDate      string             `bson:"due_date"`
    TimeTracking TimeTrackingEntity `bson:"time_tracking"`
    CreatedAt    time.Time          `bson:"created_at"`
    UpdatedAt    time.Time          `bson:"updated_at"`
}

func FromDomainTask(m *models.Task) *TaskEntity {
    if m == nil {
        return nil
    }
    var projectID, assignedTo, createdBy primitive.ObjectID
    if m.ProjectID != "" {
        if oid, err := primitive.ObjectIDFromHex(m.ProjectID); err == nil {
            projectID = oid
        }
    }
    if m.AssignedTo != "" {
        if oid, err := primitive.ObjectIDFromHex(m.AssignedTo); err == nil {
            assignedTo = oid
        }
    }
    if m.CreatedBy != "" {
        if oid, err := primitive.ObjectIDFromHex(m.CreatedBy); err == nil {
            createdBy = oid
        }
    }
    return &TaskEntity{
        ID:          m.ID,
        Title:       m.Title,
        Description: m.Description,
        ProjectID:   projectID,
        AssignedTo:  assignedTo,
        CreatedBy:   createdBy,
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

func (e *TaskEntity) ToDomainTask() *models.Task {
    if e == nil {
        return nil
    }
    return &models.Task{
        ID:          e.ID,
        Title:       e.Title,
        Description: e.Description,
        ProjectID:   e.ProjectID.Hex(),
        AssignedTo:  e.AssignedTo.Hex(),
        CreatedBy:   e.CreatedBy.Hex(),
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


