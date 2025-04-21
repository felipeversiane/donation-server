package neighborhood

import (
	"errors"
	"strings"
)

type Neighborhood struct {
	value string
}

func New(value string) (Neighborhood, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return Neighborhood{}, errors.New("neighborhood cannot be empty")
	}
	return Neighborhood{value: value}, nil
}

func (n Neighborhood) String() string {
	return n.value
}
