package email

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidEmail = errors.New("invalid email format")
)

type Email struct {
	value string
}

func New(value string) (Email, error) {
	if value == "" {
		return Email{}, errors.New("email is required")
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !regex.MatchString(value) {
		return Email{}, ErrInvalidEmail
	}
	return Email{value: value}, nil
}

func (e Email) String() string {
	return e.value
}

func (e Email) Equals(other Email) bool {
	return e.value == other.value
}
