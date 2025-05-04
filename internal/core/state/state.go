package state

import (
	"time"

	"github.com/felipeversiane/donation-server/internal/core/vo/uuid"
)

type State struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	UF        string     `json:"uf"`
	CountryID *uuid.UUID `json:"country_id,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func New(name, uf string, countryID *uuid.UUID) (*State, error) {
	id, err := uuid.New()
	if err != nil {
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
