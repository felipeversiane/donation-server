package uuid

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type UUID struct {
	value uuid.UUID
}

func New() (UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return UUID{}, errors.Wrap(err, "failed to generate UUID v7")
	}

	return UUID{value: id}, nil
}

func FromString(input string) (UUID, error) {
	id, err := uuid.Parse(input)
	if err != nil {
		return UUID{}, errors.Wrap(err, "invalid UUID format")
	}

	return UUID{value: id}, nil
}

func (u UUID) String() string {
	return u.value.String()
}

func (u UUID) UUID() uuid.UUID {
	return u.value
}
