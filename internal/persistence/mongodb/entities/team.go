package entities

import (
    "devflow/internal/models"
    "time"
)

type TeamMemberEntity struct {
    UserID   string    `bson:"user_id"`
    Role     string    `bson:"role"`
    JoinedAt time.Time `bson:"joined_at"`
}

type TeamSettingsEntity struct {
    IsPrivate         bool `bson:"is_private"`
    AllowMemberInvite bool `bson:"allow_member_invite"`
}

type TeamEntity struct {
    ID          string              `bson:"_id"`
    Name        string              `bson:"name"`
    Description string              `bson:"description"`
    OwnerID     string              `bson:"owner_id"`
    Members     []TeamMemberEntity  `bson:"members"`
    Settings    TeamSettingsEntity  `bson:"settings"`
    CreatedAt   time.Time           `bson:"created_at"`
    UpdatedAt   time.Time           `bson:"updated_at"`
}

func TeamFromModel(m *models.Team) *TeamEntity {
    if m == nil {
        return nil
    }
    members := make([]TeamMemberEntity, 0, len(m.Members))
    for _, mm := range m.Members {
        members = append(members, TeamMemberEntity{UserID: mm.UserID, Role: mm.Role, JoinedAt: mm.JoinedAt})
    }
    return &TeamEntity{
        ID:          m.ID,
        Name:        m.Name,
        Description: m.Description,
        OwnerID:     m.OwnerID,
        Members:     members,
        Settings: TeamSettingsEntity{
            IsPrivate:         m.Settings.IsPrivate,
            AllowMemberInvite: m.Settings.AllowMemberInvite,
        },
        CreatedAt: m.CreatedAt,
        UpdatedAt: m.UpdatedAt,
    }
}

func (e *TeamEntity) ToModel() *models.Team {
    if e == nil {
        return nil
    }
    members := make([]models.TeamMember, 0, len(e.Members))
    for _, mm := range e.Members {
        members = append(members, models.TeamMember{UserID: mm.UserID, Role: mm.Role, JoinedAt: mm.JoinedAt})
    }
    return &models.Team{
        ID:          e.ID,
        Name:        e.Name,
        Description: e.Description,
        OwnerID:     e.OwnerID,
        Members:     members,
        Settings: models.TeamSettings{
            IsPrivate:         e.Settings.IsPrivate,
            AllowMemberInvite: e.Settings.AllowMemberInvite,
        },
        CreatedAt: e.CreatedAt,
        UpdatedAt: e.UpdatedAt,
    }
}


