package requests

type TeamSettingsReq struct {
    IsPrivate         bool `json:"is_private" validate:"boolean"`
    AllowMemberInvite bool `json:"allow_member_invite" validate:"boolean"`
}

type CreateTeamMemberReq struct {
    UserID string `json:"user_id" validate:"required,len=24,hexadecimal"`
    Role   string `json:"role" validate:"required,oneof=admin member viewer"`
}

type CreateTeamReq struct {
    Name        string                `json:"name" validate:"required,min=2"`
    Description string                `json:"description" validate:"omitempty,max=1024"`
    OwnerID     string                `json:"owner_id" validate:"required,len=24,hexadecimal"`
	Members     []CreateTeamMemberReq `json:"members"`
	Settings    TeamSettingsReq       `json:"settings"`
}

type UpdateTeamReq struct {
    Name        *string          `json:"name" validate:"omitempty,min=2"`
    Description *string          `json:"description" validate:"omitempty,max=1024"`
	Settings    *TeamSettingsReq `json:"settings"`
}

type AddMemberReq struct {
    UserID string `json:"user_id" validate:"required,len=24,hexadecimal"`
    Role   string `json:"role" validate:"required,oneof=admin member viewer"`
}
