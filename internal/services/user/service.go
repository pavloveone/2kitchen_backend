package userservices

import (
	"2kitchen/internal/models"
	userrepositories "2kitchen/internal/repositories/user"
)

type UserService struct {
	repo *userrepositories.UserRepository
}

func NewUserRepository(repo *userrepositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) AllUsers() ([]models.UserResponse, error) {
	return s.repo.AllUsers()
}

func (s *UserService) UserById(id int) (models.UserResponse, error) {
	return s.repo.UserById(id)
}

func (s *UserService) AddUser(newUser models.CreateUserRequest) (int, error) {
	return s.repo.AddUser(newUser)
}

func (s *UserService) LogIn(loginUser models.LogInUser) (models.UserResponse, error) {
	return s.repo.LogIn(loginUser)
}
