package mongodb

import (
	"context"
	"errors"
	"time"

	"devflow/internal/interfaces"
	"devflow/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	col *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{col: db.Collection("users")}
}

func (r *UserRepository) Create(ctx context.Context, u *models.User) (string, error) {
	if u == nil {
		return "", errors.New("nil user")
	}
	now := time.Now()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = now
	}
	u.UpdatedAt = now
	_, err := r.col.InsertOne(ctx, u)
	if err != nil {
		return "", err
	}
	return u.ID, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	var out models.User
	err := r.col.FindOne(ctx, bson.M{"id": id}).Decode(&out)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	return &out, err
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var out models.User
	err := r.col.FindOne(ctx, bson.M{"username": username}).Decode(&out)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	return &out, err
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var out models.User
	err := r.col.FindOne(ctx, bson.M{"email": email}).Decode(&out)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	return &out, err
}

func (r *UserRepository) List(ctx context.Context) ([]*models.User, error) {
	cur, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var out []*models.User
	for cur.Next(ctx) {
		var u models.User
		if err := cur.Decode(&u); err != nil {
			return nil, err
		}
		out = append(out, &u)
	}
	return out, cur.Err()
}

func (r *UserRepository) FilterByRole(ctx context.Context, role string) ([]*models.User, error) {
	cur, err := r.col.Find(ctx, bson.M{"role": role})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var out []*models.User
	for cur.Next(ctx) {
		var u models.User
		if err := cur.Decode(&u); err != nil {
			return nil, err
		}
		out = append(out, &u)
	}
	return out, cur.Err()
}

func (r *UserRepository) UpdateProfile(ctx context.Context, id string, p models.Profile) error {
	_, err := r.col.UpdateOne(ctx,
		bson.M{"id": id},
		bson.M{"$set": bson.M{
			"profile.first_name": p.FirstName,
			"profile.last_name":  p.LastName,
			"profile.avatar_url": p.AvatarURL,
			"updated_at":         time.Now(),
		}},
	)
	return err
}

func (r *UserRepository) UpdateCore(ctx context.Context, id, username, email, passwordHash, role string) error {
	set := bson.M{"updated_at": time.Now()}
	if username != "" {
		set["username"] = username
	}
	if email != "" {
		set["email"] = email
	}
	if passwordHash != "" {
		set["password_hash"] = passwordHash
	}
	if role != "" {
		set["role"] = role
	}
	_, err := r.col.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": set})
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	_, err := r.col.DeleteOne(ctx, bson.M{"id": id})
	return err
}

var _ interfaces.UserRepository = (*UserRepository)(nil)
