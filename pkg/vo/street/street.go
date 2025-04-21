package street

import (
	"errors"
	"strings"
)

type Street struct {
	value string
}

func New(value string) (Street, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return Street{}, errors.New("street cannot be empty")
	}
	return Street{value: value}, nil
}

func (s Street) String() string {
	return s.value
}
