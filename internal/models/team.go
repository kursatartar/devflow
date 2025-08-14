package models

import "time"

var Teams = map[string]*Team{}

type TeamMember struct {
	UserID   string    `json:"user_id"`
	Role     string    `json:"role"`
	JoinedAt time.Time `json:"joined_at"`
}

type TeamSettings struct {
	IsPrivate         bool `json:"is_private"`
	AllowMemberInvite bool `json:"allow_member_invite"`
}

type Team struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	OwnerID     string       `json:"owner_id"`
	Members     []TeamMember `json:"members"`
	Settings    TeamSettings `json:"settings"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func NewTeam(id, name, description, ownerID string, members []TeamMember, settings TeamSettings) *Team {
	return &Team{
		ID:          id,
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
		Members:     members,
		Settings:    settings,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Team) GetID() string { return t.ID }
