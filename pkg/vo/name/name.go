package name

import (
	"fmt"
	"strings"
)

type Name struct {
	value string
}

func New(value string) (Name, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return Name{}, fmt.Errorf("name is required")
	}
	if len(value) > 100 {
		return Name{}, fmt.Errorf("name must be at most 100 characters")
	}
	return Name{value: value}, nil
}

func (t Name) String() string {
	return t.value
}
