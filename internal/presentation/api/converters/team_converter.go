package converters

import (
	"devflow/internal/models"
	"devflow/internal/presentation/api/requests"
	"devflow/internal/presentation/api/responses"
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

func ToDomainTeamSettings(req requests.TeamSettingsReq) models.TeamSettings {
	return models.TeamSettings{
		IsPrivate:         req.IsPrivate,
		AllowMemberInvite: req.AllowMemberInvite,
	}
}

func ToDomainTeamSettingsPtr(req *requests.TeamSettingsReq) *models.TeamSettings {
	if req == nil {
		return nil
	}
	m := ToDomainTeamSettings(*req)
	return &m
}

func ToDomainTeamMembers(reqs []requests.CreateTeamMemberReq) []models.TeamMember {
	out := make([]models.TeamMember, 0, len(reqs))
	for _, m := range reqs {
		out = append(out, models.TeamMember{
			UserID: m.UserID,
			Role:   m.Role,
		})
	}
	return out
}
