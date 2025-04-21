package phone

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidPhone = errors.New("invalid phone number format")
)

type Phone struct {
	value string
}

func New(value string) (Phone, error) {
	regex := regexp.MustCompile(`^\(?\d{2}\)?\s?\d{4,5}-?\d{4}$`)
	if !regex.MatchString(value) {
		return Phone{}, ErrInvalidPhone
	}
	return Phone{value: value}, nil
}

func (p Phone) String() string {
	return p.value
}

func (p Phone) Equals(other Phone) bool {
	return p.value == other.value
}
