package repository

import (
    "context"
    "database/sql"
    "crud-app/db/sqlc"
    "crud-app/model"
    "time"
)

type UserRepository interface {
    Create(ctx context.Context, user *model.User) error
    GetByID(ctx context.Context, id int64) (*model.User, error)
    Update(ctx context.Context, user *model.User) error
    Delete(ctx context.Context, id int64) error
    List(ctx context.Context) ([]model.User, error)
}

type userRepository struct {
    queries *db.Queries
}

func NewUserRepository(dbConn *sql.DB) UserRepository {
    return &userRepository{
        queries: db.New(dbConn),
    }
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
    now := time.Now()
    params := db.CreateUserParams{
        Name:      user.Name,
        Email:     user.Email,
        CreatedAt: now,
        UpdatedAt: now,
    }
    
    result, err := r.queries.CreateUser(ctx, params)
    if err != nil {
        return err
    }
    
    user.ID = result.ID
    user.CreatedAt = result.CreatedAt
    user.UpdatedAt = result.UpdatedAt
    return nil
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
    result, err := r.queries.GetUser(ctx, id)
    if err != nil {
        return nil, err
    }
    
    return &model.User{
        ID:        result.ID,
        Name:      result.Name,
        Email:     result.Email,
        CreatedAt: result.CreatedAt,
        UpdatedAt: result.UpdatedAt,
    }, nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
    params := db.UpdateUserParams{
        ID:        user.ID,
        Name:      user.Name,
        Email:     user.Email,
        UpdatedAt: time.Now(),
    }
    
    result, err := r.queries.UpdateUser(ctx, params)
    if err != nil {
        return err
    }
    
    user.UpdatedAt = result.UpdatedAt
    return nil
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
    return r.queries.DeleteUser(ctx, id)
}

func (r *userRepository) List(ctx context.Context) ([]model.User, error) {
    results, err := r.queries.ListUsers(ctx)
    if err != nil {
        return nil, err
    }
    
    users := make([]model.User, len(results))
    for i, result := range results {
        users[i] = model.User{
            ID:        result.ID,
            Name:      result.Name,
            Email:     result.Email,
            CreatedAt: result.CreatedAt,
            UpdatedAt: result.UpdatedAt,
        }
    }
    return users, nil
}
