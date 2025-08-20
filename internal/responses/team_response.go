package responses

import "devflow/internal/models"

func TeamResource(t *models.Team) map[string]any {
	ms := make([]map[string]any, 0, len(t.Members))
	for _, m := range t.Members {
		ms = append(ms, map[string]any{
			"userId":   m.UserID,
			"role":     m.Role,
			"joinedAt": m.JoinedAt,
		})
	}
	return map[string]any{
		"id":          t.ID,
		"name":        t.Name,
		"description": t.Description,
		"ownerId":     t.OwnerID,
		"members":     ms,
		"settings": map[string]any{
			"isPrivate":         t.Settings.IsPrivate,
			"allowMemberInvite": t.Settings.AllowMemberInvite,
		},
		"createdAt": t.CreatedAt,
		"updatedAt": t.UpdatedAt,
	}
}

func TeamList(ts []*models.Team) []map[string]any {
	out := make([]map[string]any, 0, len(ts))
	for _, t := range ts {
		out = append(out, TeamResource(t))
	}
	return out
}
