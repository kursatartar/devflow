package models

import (
	"errors"
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

func NewUser(id, username, email, passwordHash, role string, profile Profile) (*User, error) {
	if id == "" {
		return nil, errors.New("kullanıcı ID'si boş olamaz")
	}
	if username == "" {
		return nil, errors.New("kullanıcı adı boş olamaz")
	}
	if email == "" {
		return nil, errors.New("email adresi boş olamaz")
	}
	if passwordHash == "" {
		return nil, errors.New("şifre hash'i boş olamaz")
	}
	if role == "" {
		return nil, errors.New("kullanıcı rolü boş olamaz")
	}

	return &User{
		ID:           id,
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		Profile:      profile,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}

func (u *User) IsEmailValid() (bool, error) {
	if u.Email == "" {
		return false, errors.New("email adresi boş olamaz")
	}
	return strings.Contains(u.Email, "@"), nil
}

func (u *User) GetID() string {
	return u.ID
}
