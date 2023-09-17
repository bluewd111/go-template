package domain

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID    string
	Name  string
	Email EmailAddress
	Age   int
}

func NewUser(name, email string, age int) (*User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}
	emailAddress, err := NewEmailAddress(email)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:    uuid.NewString(),
		Name:  name,
		Email: *emailAddress,
		Age:   age,
	}, nil
}

func (u *User) UpdateName(name string) {
	u.Name = name
}

func (u *User) UpdateEmail(email string) error {
	emailAddress, err := NewEmailAddress(email)
	if err != nil {
		return err
	}
	u.Email = *emailAddress
	return nil
}

func (u *User) UpdateAge(age int) {
	u.Age = age
}
