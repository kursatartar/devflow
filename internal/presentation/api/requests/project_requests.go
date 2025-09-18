package requests

type CreateProjectReq struct {
    Name         string   `json:"name" validate:"required,min=2"`
    Description  string   `json:"description" validate:"omitempty,max=1024"`
    OwnerID      string   `json:"owner_id" validate:"required,len=24,hexadecimal"`
    TeamID       string   `json:"team_id" validate:"required,len=24,hexadecimal"`
    Status       string   `json:"status" validate:"required,oneof=active archived on_hold"`
	TeamMembers  []string `json:"team_members"`
    IsPrivate    bool     `json:"is_private" validate:"boolean"`
	TaskWorkflow []string `json:"task_workflow"`
}

type UpdateProjectReq struct {
    Name         string   `json:"name" validate:"omitempty,min=2"`
    Description  string   `json:"description" validate:"omitempty,max=1024"`
    Status       string   `json:"status" validate:"omitempty,oneof=active archived on_hold"`
    TeamID       string   `json:"team_id" validate:"omitempty,len=24,hexadecimal"`
	TeamMembers  []string `json:"team_members"`
    IsPrivate    bool     `json:"is_private" validate:"boolean"`
	TaskWorkflow []string `json:"task_workflow"`
}
