package entities

import (
    "devflow/internal/models"
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectSettingsEntity struct {
    IsPrivate    bool     `bson:"is_private"`
    TaskWorkflow []string `bson:"task_workflow"`
}

type ProjectEntity struct {
    ID          string                 `bson:"_id"`
    Name        string                 `bson:"name"`
    Description string                 `bson:"description"`
    OwnerID     primitive.ObjectID     `bson:"owner_id"`
    TeamID      primitive.ObjectID     `bson:"team_id"`
    Status      string                 `bson:"status"`
    Settings    ProjectSettingsEntity  `bson:"settings"`
    CreatedAt   time.Time              `bson:"created_at"`
    UpdatedAt   time.Time              `bson:"updated_at"`
}

func FromDomainProject(m *models.Project) *ProjectEntity {
    if m == nil {
        return nil
    }
    var ownerID, teamID primitive.ObjectID
    if m.OwnerID != "" {
        if oid, err := primitive.ObjectIDFromHex(m.OwnerID); err == nil {
            ownerID = oid
        }
    }
    if m.TeamID != "" {
        if oid, err := primitive.ObjectIDFromHex(m.TeamID); err == nil {
            teamID = oid
        }
    }
    return &ProjectEntity{
        ID:          m.ID,
        Name:        m.Name,
        Description: m.Description,
        OwnerID:     ownerID,
        TeamID:      teamID,
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
        OwnerID:     e.OwnerID.Hex(),
        TeamID:      e.TeamID.Hex(),
        Status:      e.Status,
        Settings: models.ProjectSettings{
            IsPrivate:    e.Settings.IsPrivate,
            TaskWorkflow: e.Settings.TaskWorkflow,
        },
        CreatedAt: e.CreatedAt,
        UpdatedAt: e.UpdatedAt,
    }
}


