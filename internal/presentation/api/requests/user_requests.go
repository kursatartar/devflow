package requests

type CreateUserReq struct {
    Username     string `json:"username" validate:"required,min=3,max=32"`
    Email        string `json:"email" validate:"required,email"`
    PasswordHash string `json:"password_hash" validate:"required,min=6"`
    Role         string `json:"role" validate:"required,oneof=admin user manager"`
    FirstName    string `json:"first_name" validate:"required,min=2"`
    LastName     string `json:"last_name" validate:"required,min=2"`
    AvatarURL    string `json:"avatar_url" validate:"required,url"`
}

type UpdateUserReq struct {
    Username     *string `json:"username" validate:"omitempty,min=3,max=32"`
    Email        *string `json:"email" validate:"omitempty,email"`
    PasswordHash *string `json:"password_hash" validate:"omitempty,min=6"`
    Role         *string `json:"role" validate:"omitempty,oneof=admin user manager"`
    FirstName    *string `json:"first_name" validate:"omitempty,min=2"`
    LastName     *string `json:"last_name" validate:"omitempty,min=2"`
    AvatarURL    *string `json:"avatar_url" validate:"omitempty,url"`
}
