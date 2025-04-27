package document

import (
	"time"

	"github.com/felipeversiane/donation-server/pkg/helpers"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Type string

const (
	CPF  Type = "cpf"
	CNPJ Type = "cnpj"
)

type Document struct {
	ID        uuid.UUID `json:"id"`
	Type      Type      `json:"type"`
	Value     string    `json:"value"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(docTypeStr, value string, userID uuid.UUID) (*Document, error) {
	docType := Type(helpers.CleanString(docTypeStr))
	if docType != CPF && docType != CNPJ {
		return nil, errors.New("invalid document type")
	}

	if userID == uuid.Nil {
		return nil, errors.New("user_id is required")
	}

	cleanedValue := helpers.CleanString(value)

	if docType == CPF {
		if len(cleanedValue) != 11 {
			return nil, errors.New("CPF must have 11 digits")
		}
	} else if docType == CNPJ {
		if len(cleanedValue) != 14 {
			return nil, errors.New("CNPJ must have 14 digits")
		}
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, errors.Wrap(err, "generating document ID")
	}

	now := time.Now()

	return &Document{
		ID:        id,
		Type:      docType,
		Value:     cleanedValue,
		UserID:    userID,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (d *Document) IsValid() bool {
	if d.Type == CPF && len(d.Value) == 11 {
		return true
	}
	if d.Type == CNPJ && len(d.Value) == 14 {
		return true
	}
	return false
}

func (d *Document) Update(value string) error {
	cleanedValue := helpers.CleanString(value)

	if d.Type == CPF {
		if len(cleanedValue) != 11 {
			return errors.New("CPF must have 11 digits")
		}
	} else if d.Type == CNPJ {
		if len(cleanedValue) != 14 {
			return errors.New("CNPJ must have 14 digits")
		}
	}

	d.Value = cleanedValue
	d.UpdatedAt = time.Now()
	return nil
}

func (d *Document) Format() string {
	if d.Type == CPF {
		if len(d.Value) == 11 {
			return d.Value[0:3] + "." + d.Value[3:6] + "." + d.Value[6:9] + "-" + d.Value[9:11]
		}
	} else if d.Type == CNPJ {
		if len(d.Value) == 14 {
			return d.Value[0:2] + "." + d.Value[2:5] + "." + d.Value[5:8] + "/" + d.Value[8:12] + "-" + d.Value[12:14]
		}
	}
	return d.Value
}
