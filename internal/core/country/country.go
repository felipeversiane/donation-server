package country

import (
	"time"

	"github.com/felipeversiane/donation-server/pkg/vo/uuid"
)

type Country struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(name string) (*Country, error) {
	id, err := uuid.New()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Country{
		ID:        id,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
