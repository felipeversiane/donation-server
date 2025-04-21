package number

import (
	"errors"
	"strings"
)

type Number struct {
	value string
}

func New(value string) (Number, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return Number{}, errors.New("number cannot be empty")
	}
	return Number{value: value}, nil
}

func (n Number) String() string {
	return n.value
}
