package domain

import (
	"errors"
	"regexp"
)

type EmailAddress struct {
	Value string
}

func NewEmailAddress(email string) (*EmailAddress, error) {
	valid := regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,4}$")
	if valid.MatchString(email) == false {
		return nil, errors.New("email address is invalid")
	}

	return &EmailAddress{
		Value: email,
	}, nil
}
