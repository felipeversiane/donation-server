package city

import (
	"time"

	"github.com/felipeversiane/donation-server/internal/core/vo/uuid"
	"github.com/felipeversiane/donation-server/pkg/field"
)

type City struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	StateID   *uuid.UUID `json:"state_id,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func New(name string, stateID *uuid.UUID) (*City, error) {
	if err := field.ValidateRequired(name, "name"); err != nil {
		return nil, err
	}
	if err := field.ValidateMinLength(name, 2, "name"); err != nil {
		return nil, err
	}
	if err := field.ValidateMaxLength(name, 100, "name"); err != nil {
		return nil, err
	}

	id, err := uuid.New()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &City{
		ID:        id,
		Name:      name,
		StateID:   stateID,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
