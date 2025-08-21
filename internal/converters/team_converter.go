package converters

import (
	"devflow/internal/models"
	"devflow/internal/responses"
)

func ToTeamResponse(t *models.Team) responses.TeamResponse {
	members := make([]responses.TeamMemberResponse, 0, len(t.Members))
	for _, m := range t.Members {
		members = append(members, responses.TeamMemberResponse{
			UserID:   m.UserID,
			Role:     m.Role,
			JoinedAt: m.JoinedAt,
		})
	}
	return responses.TeamResponse{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		OwnerID:     t.OwnerID,
		Members:     members,
		Settings: responses.TeamSettingsResponse{
			IsPrivate:         t.Settings.IsPrivate,
			AllowMemberInvite: t.Settings.AllowMemberInvite,
		},
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func ToTeamListResponse(ts []*models.Team) responses.TeamListResponse {
	items := make([]responses.TeamResponse, 0, len(ts))
	for _, t := range ts {
		items = append(items, ToTeamResponse(t))
	}
	var out responses.TeamListResponse
	out.Teams = items
	out.Metadata.Total = int64(len(items))
	return out
}
