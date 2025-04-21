package state

import (
	"errors"
	"strings"
)

type State struct {
	value string
}

func New(value string) (State, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return State{}, errors.New("state cannot be empty")
	}
	return State{value: value}, nil
}

func (s State) String() string {
	return s.value
}
