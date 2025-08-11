package models

import (
	"strings"
	"time"
)

type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AvatarURL string `json:"avatar_url"`
}

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Role         string    `json:"role"`
	Profile      Profile   `json:"profile"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

var Users = map[string]*User{}

func NewUser(id, username, email, passwordHash, role string, profile Profile) *User {
	return &User{
		ID:           id,
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		Profile:      profile,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

}

func (u *User) IsEmailValid() bool {
	return strings.Contains(u.Email, "@")
}

func (u *User) GetID() string {
	return u.ID
}
