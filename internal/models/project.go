package models

import "time"

type ProjectSettings struct {
	IsPrivate    bool     `json:"is_private"    bson:"is_private"`
	TaskWorkflow []string `json:"task_workflow" bson:"task_workflow"`
}

type Project struct {
	ID          string          `json:"id"           bson:"_id,omitempty"`
	Name        string          `json:"name"         bson:"name"`
	Description string          `json:"description"  bson:"description"`
	OwnerID     string          `json:"owner_id"     bson:"owner_id"`
    TeamID      string          `json:"team_id"      bson:"team_id"`
	TeamMembers []string        `json:"team_members" bson:"team_members"`
	Status      string          `json:"status"       bson:"status"`
	Settings    ProjectSettings `json:"settings"     bson:"settings"`
	CreatedAt   time.Time       `json:"created_at"   bson:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"   bson:"updated_at"`
}

func NewProject(id, name, description, ownerID, teamID string, teamMembers []string, status string, isPrivate bool, taskWorkflow []string) *Project {
	return &Project{
		ID:          id,
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
        TeamID:      teamID,
		TeamMembers: teamMembers,
		Status:      status,
		Settings: ProjectSettings{
			IsPrivate:    isPrivate,
			TaskWorkflow: taskWorkflow,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *Project) GetID() string {
	return u.ID
}
