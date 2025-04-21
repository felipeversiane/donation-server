package city

import (
	"errors"
	"strings"
)

type City struct {
	value string
}

func New(value string) (City, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return City{}, errors.New("city cannot be empty")
	}
	return City{value: value}, nil
}

func (c City) String() string {
	return c.value
}
