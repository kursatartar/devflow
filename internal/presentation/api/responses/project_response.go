package responses

import "time"

type ProjectResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	OwnerID      string    `json:"owner_id"`
    TeamID       string    `json:"team_id"`
	TeamMembers  []string  `json:"team_members"`
	Status       string    `json:"status"`
	IsPrivate    bool      `json:"is_private"`
	TaskWorkflow []string  `json:"task_workflow"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
    Team         *TeamResponse `json:"team,omitempty"`
}

type ProjectListResponse struct {
	Projects []ProjectResponse `json:"projects"`
	Metadata struct {
		Total int64 `json:"total"`
	} `json:"metadata"`
}
