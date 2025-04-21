package country

import (
	"errors"
	"strings"
)

type Country string

func New(value string) (Country, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", errors.New("country cannot be empty")
	}
	return Country(value), nil
}

func (c Country) String() string {
	return string(c)
}
