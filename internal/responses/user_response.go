package responses

import "devflow/internal/models"

func UserResource(u *models.User) map[string]any {
	return map[string]any{
		"id":        u.ID,
		"username":  u.Username,
		"email":     u.Email,
		"role":      u.Role,
		"firstName": u.Profile.FirstName,
		"lastName":  u.Profile.LastName,
		"avatarURL": u.Profile.AvatarURL,
		"createdAt": u.CreatedAt,
		"updatedAt": u.UpdatedAt,
	}
}

func UserList(users []*models.User) []map[string]any {
	out := make([]map[string]any, 0, len(users))
	for _, u := range users {
		out = append(out, UserResource(u))
	}
	return out
}
