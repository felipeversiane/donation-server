package city

import (
	"time"

	"github.com/felipeversiane/donation-server/internal/core/vo/uuid"
)

type City struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	StateID   *uuid.UUID `json:"state_id,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func New(name string, stateID *uuid.UUID) (*City, error) {
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
