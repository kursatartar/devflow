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

type TaskRepository struct {
    col *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) *TaskRepository {
    return &TaskRepository{col: db.Collection("tasks")}
}

func (r *TaskRepository) Create(ctx context.Context, t *models.Task) (string, error) {
    if t == nil {
        return "", errors.New("nil task")
    }
    now := time.Now()
    if t.CreatedAt.IsZero() {
        t.CreatedAt = now
    }
    t.UpdatedAt = now
    _, err := r.col.InsertOne(ctx, entities.TaskFromModel(t))
    if err != nil {
        return "", err
    }
    return t.ID, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id string) (*models.Task, error) {
    var out entities.TaskEntity
    err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&out)
    if errors.Is(err, mongo.ErrNoDocuments) {
        return nil, nil
    }
    return out.ToModel(), err
}

func (r *TaskRepository) List(ctx context.Context) ([]*models.Task, error) {
    cur, err := r.col.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cur.Close(ctx)
    var out []*models.Task
    for cur.Next(ctx) {
        var e entities.TaskEntity
        if err := cur.Decode(&e); err != nil {
            return nil, err
        }
        out = append(out, e.ToModel())
    }
    return out, cur.Err()
}

func (r *TaskRepository) FilterByProject(ctx context.Context, projectID string) ([]*models.Task, error) {
    cur, err := r.col.Find(ctx, bson.M{"project_id": projectID})
    if err != nil {
        return nil, err
    }
    defer cur.Close(ctx)
    var out []*models.Task
    for cur.Next(ctx) {
        var e entities.TaskEntity
        if err := cur.Decode(&e); err != nil {
            return nil, err
        }
        out = append(out, e.ToModel())
    }
    return out, cur.Err()
}

func (r *TaskRepository) UpdateFields(ctx context.Context, id string, title, description, status, priority, dueDate *string, labels *[]string, estimated, logged *float64) error {
    set := bson.M{"updated_at": time.Now()}
    if title != nil {
        set["title"] = *title
    }
    if description != nil {
        set["description"] = *description
    }
    if status != nil {
        set["status"] = *status
    }
    if priority != nil {
        set["priority"] = *priority
    }
    if dueDate != nil {
        set["due_date"] = *dueDate
    }
    if labels != nil {
        set["labels"] = *labels
    }
    if estimated != nil {
        set["time_tracking.estimated_hours"] = *estimated
    }
    if logged != nil {
        set["time_tracking.logged_hours"] = *logged
    }
    _, err := r.col.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": set})
    return err
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
    _, err := r.col.DeleteOne(ctx, bson.M{"_id": id})
    return err
}

var _ interfaces.TaskRepository = (*TaskRepository)(nil)


