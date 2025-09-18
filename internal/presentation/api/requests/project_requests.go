package requests

type CreateProjectReq struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	OwnerID      string   `json:"owner_id"`
    TeamID       string   `json:"team_id"`
	Status       string   `json:"status"`
	TeamMembers  []string `json:"team_members"`
	IsPrivate    bool     `json:"is_private"`
	TaskWorkflow []string `json:"task_workflow"`
}

type UpdateProjectReq struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Status       string   `json:"status"`
    TeamID       string   `json:"team_id"`
	TeamMembers  []string `json:"team_members"`
	IsPrivate    bool     `json:"is_private"`
	TaskWorkflow []string `json:"task_workflow"`
}
