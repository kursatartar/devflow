package repositories

import (
    "context"
    "errors"
    "time"

    "devflow/internal/interfaces"
    "devflow/internal/models"
    "devflow/internal/persistence/mongodb/entities"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type TeamRepository struct {
    col *mongo.Collection
}

func NewTeamRepository(db *mongo.Database) *TeamRepository {
    return &TeamRepository{col: db.Collection("teams")}
}

func (r *TeamRepository) Create(ctx context.Context, t *models.Team) (string, error) {
    if t == nil {
        return "", errors.New("nil team")
    }
    now := time.Now()
    if t.CreatedAt.IsZero() {
        t.CreatedAt = now
    }
    t.UpdatedAt = now
    _, err := r.col.InsertOne(ctx, entities.TeamFromModel(t))
    if err != nil {
        return "", err
    }
    return t.ID, nil
}

func (r *TeamRepository) GetByID(ctx context.Context, id string) (*models.Team, error) {
    var out entities.TeamEntity
    err := r.col.FindOne(ctx, bson.M{"id": id}).Decode(&out)
    if errors.Is(err, mongo.ErrNoDocuments) {
        return nil, nil
    }
    return out.ToModel(), err
}

func (r *TeamRepository) List(ctx context.Context) ([]*models.Team, error) {
    cur, err := r.col.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cur.Close(ctx)
    var out []*models.Team
    for cur.Next(ctx) {
        var e entities.TeamEntity
        if err := cur.Decode(&e); err != nil {
            return nil, err
        }
        out = append(out, e.ToModel())
    }
    return out, cur.Err()
}

func (r *TeamRepository) UpdateFields(ctx context.Context, id string, name, description *string, settings *models.TeamSettings) error {
    set := bson.M{"updated_at": time.Now()}
    if name != nil {
        set["name"] = *name
    }
    if description != nil {
        set["description"] = *description
    }
    if settings != nil {
        set["settings"] = bson.M{"is_private": settings.IsPrivate, "allow_member_invite": settings.AllowMemberInvite}
    }
    _, err := r.col.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": set})
    return err
}

func (r *TeamRepository) AddMember(ctx context.Context, teamID, userID, role string) error {
    _, err := r.col.UpdateOne(ctx, bson.M{"id": teamID},
        bson.M{"$push": bson.M{"members": bson.M{"user_id": userID, "role": role, "joined_at": time.Now()}}, "$set": bson.M{"updated_at": time.Now()}},
    )
    return err
}

func (r *TeamRepository) RemoveMember(ctx context.Context, teamID, userID string) error {
    _, err := r.col.UpdateOne(ctx, bson.M{"id": teamID},
        bson.M{"$pull": bson.M{"members": bson.M{"user_id": userID}}, "$set": bson.M{"updated_at": time.Now()}},
    )
    return err
}

func (r *TeamRepository) ChangeMemberRole(ctx context.Context, teamID, userID, role string) error {
    _, err := r.col.UpdateOne(ctx, bson.M{"id": teamID, "members.user_id": userID},
        bson.M{"$set": bson.M{"members.$.role": role, "updated_at": time.Now()}},
    )
    return err
}

func (r *TeamRepository) Delete(ctx context.Context, id string) error {
    _, err := r.col.DeleteOne(ctx, bson.M{"id": id})
    return err
}

func (r *TeamRepository) FilterByOwner(ctx context.Context, ownerID string) ([]*models.Team, error) {
    cur, err := r.col.Find(ctx, bson.M{"owner_id": ownerID})
    if err != nil {
        return nil, err
    }
    defer cur.Close(ctx)
    var out []*models.Team
    for cur.Next(ctx) {
        var e entities.TeamEntity
        if err := cur.Decode(&e); err != nil {
            return nil, err
        }
        out = append(out, e.ToModel())
    }
    return out, cur.Err()
}

var _ interfaces.TeamRepository = (*TeamRepository)(nil)


