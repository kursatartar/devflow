package requests

type TeamSettingsReq struct {
	IsPrivate         bool `json:"is_private"`
	AllowMemberInvite bool `json:"allow_member_invite"`
}

type CreateTeamMemberReq struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type CreateTeamReq struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	OwnerID     string                `json:"owner_id"`
	Members     []CreateTeamMemberReq `json:"members"`
	Settings    TeamSettingsReq       `json:"settings"`
}

type UpdateTeamReq struct {
	Name        *string          `json:"name"`
	Description *string          `json:"description"`
	Settings    *TeamSettingsReq `json:"settings"`
}

type AddMemberReq struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}
