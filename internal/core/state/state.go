package state

import (
	"time"

	"github.com/felipeversiane/donation-server/pkg/field"
	"github.com/felipeversiane/donation-server/pkg/uuid"
)

type State struct {
	ID        uuid.UUID
	Name      string
	UF        string
	CountryID *uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(name, uf string, countryID *uuid.UUID) (*State, error) {
	id, err := uuid.New()
	if err != nil {
		return nil, err
	}
	if err := field.ValidateRequired(name, "name"); err != nil {
		return nil, err
	}
	if err := field.ValidateRequired(uf, "uf"); err != nil {
		return nil, err
	}
	if err := field.ValidateMaxLength(name, 100, "name"); err != nil {
		return nil, err
	}
	if err := field.ValidateMaxLength(uf, 2, "uf"); err != nil {
		return nil, err
	}
	if err := field.ValidateMinLength(uf, 2, "uf"); err != nil {
		return nil, err
	}

	now := time.Now()

	return &State{
		ID:        id,
		Name:      name,
		UF:        uf,
		CountryID: countryID,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
