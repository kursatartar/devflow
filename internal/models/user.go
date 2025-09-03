package models

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Profile struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name"  bson:"last_name"`
	AvatarURL string `json:"avatar_url" bson:"avatar_url"`
}

type User struct {
	ID           string    `json:"id"            bson:"_id"`
	Username     string    `json:"username"      bson:"username"`
	Email        string    `json:"email"         bson:"email"`
	PasswordHash string    `json:"-"             bson:"password_hash"`
	Role         string    `json:"role"          bson:"role"`
	Profile      Profile   `json:"profile"       bson:"profile"`
	CreatedAt    time.Time `json:"created_at"   bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"   bson:"updated_at"`
}

var Users = map[string]*User{}

func NewUser(id, username, email, passwordHash, role string, profile Profile) (*User, error) {
	if id == "" {
		id = uuid.NewString()
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
