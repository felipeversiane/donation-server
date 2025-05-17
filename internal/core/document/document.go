package document

import (
	"errors"
	"regexp"
	"time"

	"github.com/felipeversiane/donation-server/internal/core/vo/documenttype"
	"github.com/felipeversiane/donation-server/internal/core/vo/uuid"
	"github.com/felipeversiane/donation-server/pkg/field"
	"github.com/felipeversiane/donation-server/pkg/str"
)

type Document struct {
	ID        uuid.UUID                 `json:"id"`
	Value     string                    `json:"value"`
	Type      documenttype.DocumentType `json:"type"`
	UserID    uuid.UUID                 `json:"user_id"`
	CreatedAt time.Time                 `json:"created_at"`
	UpdatedAt time.Time                 `json:"updated_at"`
}

func New(value string, docType string, userID uuid.UUID) (*Document, error) {
	if err := field.ValidateRequired(value, "document"); err != nil {
		return nil, err
	}

	cleanedValue := str.CleanString(value)

	docTypeVO, err := documenttype.New(docType)
	if err != nil {
		return nil, err
	}

	if err := validateValueByType(cleanedValue, docTypeVO); err != nil {
		return nil, err
	}

	id, err := uuid.New()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Document{
		ID:        id,
		Value:     cleanedValue,
		Type:      docTypeVO,
		UserID:    userID,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func validateValueByType(value string, docType documenttype.DocumentType) error {
	switch docType {
	case documenttype.CPF:
		if !regexp.MustCompile(`^\d{11}$`).MatchString(value) {
			return errors.New("invalid CPF format")
		}
	case documenttype.CNPJ:
		if !regexp.MustCompile(`^\d{14}$`).MatchString(value) {
			return errors.New("invalid CNPJ format")
		}
	default:
		return errors.New("unsupported document type")
	}
	return nil
}
