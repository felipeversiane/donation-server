package file

import (
	"time"

	"github.com/felipeversiane/donation-server/internal/core/vo/filetype"
	"github.com/felipeversiane/donation-server/pkg/field"
	"github.com/felipeversiane/donation-server/pkg/uuid"
)

type File struct {
	ID        uuid.UUID
	Name      string
	URL       string
	Type      filetype.FileType
	CreatedAt time.Time
	UpdatedAt time.Time
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
