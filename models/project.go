package models

import "time"

type Project struct {
	ID          string
	Name        string
	Description string
	OwnerID     string
	TeamMembers []string
	Status      string
	Settings    ProjectSettings
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProjectSettings struct {
	IsPrivate    bool
	TaskWorkflow []string
}

var Projects = map[string]Project{}

func NewProject(id, name, description, ownerID string, teamMembers []string, status string, isPrivate bool, taskWorkflow []string) Project {
	now := time.Now().UTC()
	return Project{
		ID:          id,
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
		TeamMembers: teamMembers,
		Status:      status,
		Settings: ProjectSettings{
			IsPrivate:    isPrivate,
			TaskWorkflow: taskWorkflow,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (p *Project) Update(newName, newDescription, newStatus string, newTeam []string, newIsPrivate bool, newWorkflow []string) {
	p.Name = newName
	p.Description = newDescription
	p.Status = newStatus
	p.TeamMembers = newTeam
	p.Settings.IsPrivate = newIsPrivate
	p.Settings.TaskWorkflow = newWorkflow
	p.UpdatedAt = time.Now().UTC()
}

// type Project struct {
//   ID          string   // `json:"id"`
//   Name        string   // `json:"name"`
//   Description string   // `json:"description"`
//   OwnerID     string   // `json:"owner_id"`
//   TeamMembers []string // `json:"team_members"`
//   Status      string   // `json:"status"`
//   Settings    Settings // `json:"settings"`
//   CreatedAt   string   // `json:"created_at"`
//   UpdatedAt   string   // `json:"updated_at"`
// }

// type Settings struct {
//   IsPrivate    bool     // `json:"is_private"`
//   TaskWorkflow []string // `json:"task_workflow"`
// }
