package entities

import (
    "devflow/internal/models"
    "time"
)

type ProjectSettingsEntity struct {
    IsPrivate    bool     `bson:"is_private"`
    TaskWorkflow []string `bson:"task_workflow"`
}

type ProjectEntity struct {
    ID          string                 `bson:"_id"`
    Name        string                 `bson:"name"`
    Description string                 `bson:"description"`
    OwnerID     string                 `bson:"owner_id"`
    TeamID      string                 `bson:"team_id"`
    Status      string                 `bson:"status"`
    Settings    ProjectSettingsEntity  `bson:"settings"`
    CreatedAt   time.Time              `bson:"created_at"`
    UpdatedAt   time.Time              `bson:"updated_at"`
}

func FromDomainProject(m *models.Project) *ProjectEntity {
    if m == nil {
        return nil
    }
    return &ProjectEntity{
        ID:          m.ID,
        Name:        m.Name,
        Description: m.Description,
        OwnerID:     m.OwnerID,
        TeamID:      m.TeamID,
        Status:      m.Status,
        Settings: ProjectSettingsEntity{
            IsPrivate:    m.Settings.IsPrivate,
            TaskWorkflow: m.Settings.TaskWorkflow,
        },
        CreatedAt: m.CreatedAt,
        UpdatedAt: m.UpdatedAt,
    }
}

func (e *ProjectEntity) ToDomainProject() *models.Project {
    if e == nil {
        return nil
    }
    return &models.Project{
        ID:          e.ID,
        Name:        e.Name,
        Description: e.Description,
        OwnerID:     e.OwnerID,
        TeamID:      e.TeamID,
        Status:      e.Status,
        Settings: models.ProjectSettings{
            IsPrivate:    e.Settings.IsPrivate,
            TaskWorkflow: e.Settings.TaskWorkflow,
        },
        CreatedAt: e.CreatedAt,
        UpdatedAt: e.UpdatedAt,
    }
}


