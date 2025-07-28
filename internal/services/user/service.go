package userservices

import (
	"2kitchen/internal/models"
	userrepositories "2kitchen/internal/repositories/user"
	"context"
)

type UserService struct {
	repo *userrepositories.UserRepository
}

func NewUserRepository(repo *userrepositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AllUsers(ctx context.Context) ([]models.UserResponse, error) {
	return s.repo.AllUsers(ctx)
}

func (s *UserService) UserById(ctx context.Context, id int) (models.UserResponse, error) {
	return s.repo.UserById(ctx, id)
}

func (s *UserService) AddUser(ctx context.Context, newUser models.CreateUserRequest) (int, error) {
	return s.repo.AddUser(ctx, newUser)
}

func (s *UserService) LogIn(ctx context.Context, loginUser models.LogInUser) (models.LoginResponse, error) {
	return s.repo.LogIn(ctx, loginUser)
}
