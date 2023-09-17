package repository

import (
	"github.com/bluewd111/go-template/app/functions"
	"github.com/bluewd111/go-template/app/user/domain"
)

type UserRepository interface {
	FindAll() []*domain.User
	FindById(id string) (*domain.User, error)
	Save(user *domain.User) error
}

type InMemoryUserRepository struct {
	data map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		data: make(map[string]*domain.User),
	}
}

func (r *InMemoryUserRepository) FindAll() []*domain.User {
	return functions.MapValues[string, domain.User](r.data)
}

func (r *InMemoryUserRepository) FindById(id string) (*domain.User, error) {
	user, ok := r.data[id]
	if !ok {
		return nil, nil
	}
	return user, nil
}

func (r *InMemoryUserRepository) Save(user *domain.User) error {
	r.data[user.ID] = user
	return nil
}
