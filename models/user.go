package models

import "time"

type Profile struct {
	FirstName string
	LastName  string
	AvatarURL string
}

type User struct {
	ID           string
	Username     string
	Email        string
	PasswordHash string
	Role         string
	Profile      Profile
	CreatedAt    string
	UpdatedAt    string
}

var Users = map[string]User{}

func NewUser(id, username, email, passwordHash, role, firstName, lastName, avatarURL string) User {
	return User{
		ID:           id,
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		Profile: Profile{
			FirstName: firstName,
			LastName:  lastName,
			AvatarURL: avatarURL,
		},
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
}
