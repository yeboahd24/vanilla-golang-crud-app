package service

import (
	"context"
	"crud-app/model"
	"crud-app/repository"
	"time"
)

// Interface Segregation
type UserService interface {
	CreateUser(ctx context.Context, name, email string) (*model.User, error)
	GetUser(ctx context.Context, id int64) (*model.User, error)
	UpdateUser(ctx context.Context, id int64, name, email string) (*model.User, error)
	DeleteUser(ctx context.Context, id int64) error
	ListUsers(ctx context.Context) ([]model.User, error)
}

// Dependency Injection
type userService struct {
	repo repository.UserRepository
}

// Constructor
func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, name, email string) (*model.User, error) {
	now := time.Now()
	user := &model.User{
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUser(ctx context.Context, id int64) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *userService) UpdateUser(ctx context.Context, id int64, name, email string) (*model.User, error) {
	user := &model.User{
		ID:        id,
		Name:      name,
		Email:     email,
		UpdatedAt: time.Now(),
	}

	err := s.repo.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *userService) ListUsers(ctx context.Context) ([]model.User, error) {
	return s.repo.List(ctx)
}
