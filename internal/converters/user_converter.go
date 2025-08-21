package converters

import (
	"devflow/internal/models"
	"devflow/internal/responses"
)

func ToUserResponse(u *models.User) responses.UserResponse {
	return responses.UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
		FirstName: u.Profile.FirstName,
		LastName:  u.Profile.LastName,
		AvatarURL: u.Profile.AvatarURL,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToUserListResponse(users []*models.User) responses.UserListResponse {
	items := make([]responses.UserResponse, 0, len(users))
	for _, u := range users {
		items = append(items, ToUserResponse(u))
	}
	var out responses.UserListResponse
	out.Users = items
	out.Metadata.Total = int64(len(items))
	return out
}
