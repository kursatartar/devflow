package requests

type CreateTaskReq struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ProjectID   string   `json:"project_id"`
	AssignedTo  string   `json:"assigned_to"`
	CreatedBy   string   `json:"created_by"`
	Status      string   `json:"status"`
	Priority    string   `json:"priority"`
	Labels      []string `json:"labels"`
	DueDate     string   `json:"due_date"`
}

type TimeTracking struct {
	Estimated float64 `json:"estimated_hours"`
	Logged    float64 `json:"logged_hours"`
}

type UpdateTaskReq struct {
	Title       *string   `json:"title"`
	Description *string   `json:"description"`
	Status      *string   `json:"status"`
	Priority    *string   `json:"priority"`
	Labels      *[]string `json:"labels"`
	DueDate     *string   `json:"due_date"`
	Estimated   *float64  `json:"estimated_hours"`
	Logged      *float64  `json:"logged_hours"`
}
