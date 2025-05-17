package file

import (
	"time"

	"github.com/felipeversiane/donation-server/internal/core/vo/filetype"
	"github.com/felipeversiane/donation-server/internal/core/vo/uuid"
	"github.com/felipeversiane/donation-server/pkg/field"
)

type File struct {
	ID        uuid.UUID         `json:"id"`
	Name      string            `json:"name"`
	URL       string            `json:"url"`
	Type      filetype.FileType `json:"type"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func New(name, url, fileTypeStr string) (*File, error) {

	id, err := uuid.New()
	if err != nil {
		return nil, err
	}

	if err := field.ValidateRequired(name, "name"); err != nil {
		return nil, err
	}
	if err := field.ValidateRequired(url, "url"); err != nil {
		return nil, err
	}
	if err := field.ValidateMaxLength(name, 100, "name"); err != nil {
		return nil, err
	}

	fileTypeVO, err := filetype.New(fileTypeStr)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &File{
		ID:        id,
		Name:      name,
		URL:       url,
		Type:      fileTypeVO,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
