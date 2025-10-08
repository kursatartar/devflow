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

type ProjectRepository struct {
    col *mongo.Collection
}

func NewProjectRepository(db *mongo.Database) *ProjectRepository {
    return &ProjectRepository{col: db.Collection("projects")}
}

func (r *ProjectRepository) Create(ctx context.Context, p *models.Project) (string, error) {
    if p == nil {
        return "", errors.New("nil project")
    }
    now := time.Now()
    if p.CreatedAt.IsZero() {
        p.CreatedAt = now
    }
    p.UpdatedAt = now
    e := entities.FromDomainProject(p)
    _, err := r.col.InsertOne(ctx, e)
    if err != nil {
        return "", err
    }
    return p.ID, nil
}

func (r *ProjectRepository) GetByID(ctx context.Context, id string) (*models.Project, error) {
    var out entities.ProjectEntity
    err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&out)
    if errors.Is(err, mongo.ErrNoDocuments) {
        return nil, nil
    }
    return out.ToDomainProject(), err
}

func (r *ProjectRepository) List(ctx context.Context) ([]*models.Project, error) {
    cur, err := r.col.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cur.Close(ctx)
    var out []*models.Project
    for cur.Next(ctx) {
        var e entities.ProjectEntity
        if err := cur.Decode(&e); err != nil {
            return nil, err
        }
        out = append(out, e.ToDomainProject())
    }
    return out, cur.Err()
}

func (r *ProjectRepository) FilterByOwner(ctx context.Context, ownerID string) ([]*models.Project, error) {
    cur, err := r.col.Find(ctx, bson.M{"owner_id": ownerID})
    if err != nil {
        return nil, err
    }
    defer cur.Close(ctx)
    var out []*models.Project
    for cur.Next(ctx) {
        var e entities.ProjectEntity
        if err := cur.Decode(&e); err != nil {
            return nil, err
        }
        out = append(out, e.ToDomainProject())
    }
    return out, cur.Err()
}

func (r *ProjectRepository) UpdateFields(ctx context.Context, id string, name, description, status *string, isPrivate *bool, taskWorkflow *[]string, ownerID, teamID *string) error {
    set := bson.M{"updated_at": time.Now()}
    if name != nil {
        set["name"] = *name
    }
    if description != nil {
        set["description"] = *description
    }
    if status != nil {
        set["status"] = *status
    }
    if isPrivate != nil {
        set["settings.is_private"] = *isPrivate
    }
    if taskWorkflow != nil {
        set["settings.task_workflow"] = *taskWorkflow
    }
    if ownerID != nil {
        set["owner_id"] = *ownerID
    }
    if teamID != nil {
        set["team_id"] = *teamID
    }
    _, err := r.col.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": set})
    return err
}

func (r *ProjectRepository) Delete(ctx context.Context, id string) error {
    _, err := r.col.DeleteOne(ctx, bson.M{"_id": id})
    return err
}

var _ interfaces.ProjectRepository = (*ProjectRepository)(nil)


