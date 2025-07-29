package models

import (
	"time"
)

type User struct {
	ID           string  // `json:"_id"`
	Username     string  // `json:"username"`
	Email        string  // `json:"email"`
	PasswordHash string  // `json:"password_hash"`
	Role         string  // `json:"role"`
	Profile      Profile // `json:"profile"`
	CreatedAt    string  // `json:"created_at"`
	UpdatedAt    string  // `json:"updated_at"`
}

type Profile struct {
	FirstName string // `json:"first_name"`
	LastName  string // `json:"last_name"`
	AvatarURL string // `json:"avatar_url"`
}

var Users = map[string]*User{}

func NewUser(id, username, email, passwordHash, role string, profile Profile) *User {
	now := time.Now().UTC().Format(time.RFC3339)
	return &User{
		ID:           id,
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		Profile:      profile,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func (u *User) Update(newUsername, newEmail, newPasswordHash, newRole string, newProfile Profile) {
	u.Username = newUsername
	u.Email = newEmail
	u.PasswordHash = newPasswordHash
	u.Role = newRole
	u.Profile = newProfile
	u.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
}
