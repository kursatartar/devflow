package entities

import (
    "devflow/internal/models"
    "time"
)

type ProfileEntity struct {
    FirstName string `bson:"first_name"`
    LastName  string `bson:"last_name"`
    AvatarURL string `bson:"avatar_url"`
}

type UserEntity struct {
    ID           string        `bson:"_id"`
    Username     string        `bson:"username"`
    Email        string        `bson:"email"`
    PasswordHash string        `bson:"password_hash"`
    Role         string        `bson:"role"`
    Profile      ProfileEntity `bson:"profile"`
    CreatedAt    time.Time     `bson:"created_at"`
    UpdatedAt    time.Time     `bson:"updated_at"`
}

func FromDomainUser(m *models.User) *UserEntity {
    if m == nil {
        return nil
    }
    return &UserEntity{
        ID:       m.ID,
        Username: m.Username,
        Email:    m.Email,
        PasswordHash: m.PasswordHash,
        Role:     m.Role,
        Profile: ProfileEntity{
            FirstName: m.Profile.FirstName,
            LastName:  m.Profile.LastName,
            AvatarURL: m.Profile.AvatarURL,
        },
        CreatedAt: m.CreatedAt,
        UpdatedAt: m.UpdatedAt,
    }
}

func (e *UserEntity) ToDomainUser() *models.User {
    if e == nil {
        return nil
    }
    return &models.User{
        ID:       e.ID,
        Username: e.Username,
        Email:    e.Email,
        PasswordHash: e.PasswordHash,
        Role:     e.Role,
        Profile: models.Profile{
            FirstName: e.Profile.FirstName,
            LastName:  e.Profile.LastName,
            AvatarURL: e.Profile.AvatarURL,
        },
        CreatedAt: e.CreatedAt,
        UpdatedAt: e.UpdatedAt,
    }
}


