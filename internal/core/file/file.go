package file

import (
	"time"

	"github.com/felipeversiane/donation-server/pkg/vo/uuid"
)

type File struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(name, url, fileType string) (*File, error) {
	id, err := uuid.New()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &File{
		ID:        id,
		Name:      name,
		URL:       url,
		Type:      fileType,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
