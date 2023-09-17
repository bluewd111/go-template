package service

import (
	"github.com/bluewd111/go-template/app/user/command"
	"github.com/bluewd111/go-template/app/user/domain"
	"github.com/bluewd111/go-template/app/user/query"
	"github.com/bluewd111/go-template/app/user/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(cmd command.CreateUserCommand) (*domain.User, error) {
	user, err := domain.NewUser(cmd.Name, cmd.Email, cmd.Age)
	if err != nil {
		return nil, err
	}
	return user, s.repo.Save(user)
}

func (s *UserService) UpdateUser(cmd command.UpdateUserCommand) (*domain.User, error) {
	user, err := s.repo.FindById(cmd.ID)
	if err != nil {
		return nil, err
	}
	user.UpdateName(cmd.Name)
	return user, s.repo.Save(user)
}

func (s *UserService) GetUser(query query.GetUserQuery) (*domain.User, error) {
	return s.repo.FindById(query.ID)
}

func (s *UserService) GetUsers() []*domain.User {
	return s.repo.FindAll()
}
