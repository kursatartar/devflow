package responses

import "time"

type TeamMemberResponse struct {
	UserID   string    `json:"user_id"`
	Role     string    `json:"role"`
	JoinedAt time.Time `json:"joined_at"`
}

type TeamSettingsResponse struct {
	IsPrivate         bool `json:"is_private"`
	AllowMemberInvite bool `json:"allow_member_invite"`
}

type TeamResponse struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	OwnerID     string               `json:"owner_id"`
	Members     []TeamMemberResponse `json:"members"`
	Settings    TeamSettingsResponse `json:"settings"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type TeamListResponse struct {
	Teams    []TeamResponse `json:"teams"`
	Metadata struct {
		Total int64 `json:"total"`
	} `json:"metadata"`
}
