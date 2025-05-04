package document

import (
	"time"

	"github.com/felipeversiane/donation-server/pkg/vo/uuid"
)

type Document struct {
	ID        uuid.UUID `json:"id"`
	Value     string    `json:"value"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(value string, userID uuid.UUID) (*Document, error) {
	id, err := uuid.New()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Document{
		ID:        id,
		Value:     value,
		UserID:    userID,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
