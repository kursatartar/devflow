package requests

import "devflow/internal/models"

type CreateTeamReq struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	OwnerID     string                `json:"owner_id"`
	Members     []CreateTeamMemberReq `json:"members"`
	Settings    models.TeamSettings   `json:"settings"`
}

type CreateTeamMemberReq struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type UpdateTeamReq struct {
	Name        *string              `json:"name"`
	Description *string              `json:"description"`
	Settings    *models.TeamSettings `json:"settings"`
}

type AddMemberReq struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}
