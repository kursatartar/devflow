package requests

type CreateTaskReq struct {
    Title       string   `json:"title" validate:"required,min=2"`
    Description string   `json:"description" validate:"omitempty,max=2048"`
    ProjectID   string   `json:"project_id" validate:"required,len=24,hexadecimal"`
    AssignedTo  string   `json:"assigned_to" validate:"omitempty,len=24,hexadecimal"`
    CreatedBy   string   `json:"created_by" validate:"required,len=24,hexadecimal"`
    Status      string   `json:"status" validate:"required,oneof=todo in_progress done"`
    Priority    string   `json:"priority" validate:"required,oneof=low medium high"`
    Labels      []string `json:"labels"`
    DueDate     string   `json:"due_date" validate:"required"`
    TimeTracking struct {
        Estimated float64 `json:"estimated_hours" validate:"omitempty,gte=0"`
        Logged    float64 `json:"logged_hours" validate:"omitempty,gte=0"`
    } `json:"time_tracking"`
}

type UpdateTaskReq struct {
    Title       *string   `json:"title" validate:"omitempty,min=2"`
    Description *string   `json:"description" validate:"omitempty,max=2048"`
    Status      *string   `json:"status" validate:"omitempty,oneof=todo in_progress done"`
    Priority    *string   `json:"priority" validate:"omitempty,oneof=low medium high"`
    Labels      *[]string `json:"labels"`
    DueDate     *string   `json:"due_date"`
    TimeTracking *struct {
        Estimated *float64 `json:"estimated_hours" validate:"omitempty,gte=0"`
        Logged    *float64 `json:"logged_hours" validate:"omitempty,gte=0"`
    } `json:"time_tracking"`
}
