package requests

type CreateUserReq struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	AvatarURL    string `json:"avatar_url"`
}

type UpdateUserReq struct {
	Username     *string `json:"username"`
	Email        *string `json:"email"`
	PasswordHash *string `json:"password_hash"`
	Role         *string `json:"role"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	AvatarURL    *string `json:"avatar_url"`
}
