package document

import (
	"errors"
	"strings"
	"time"

	"github.com/felipeversiane/donation-server/pkg/helpers"
	"github.com/google/uuid"
)

type Type string

const (
	CPF  Type = "cpf"
	CNPJ Type = "cnpj"
)

type document struct {
	id        uuid.UUID
	docType   Type
	value     string
	userID    uuid.UUID
	createdAt time.Time
	updatedAt time.Time
}

type DocumentInterface interface {
	GetID() uuid.UUID
	GetType() Type
	GetValue() string
	GetUserID() uuid.UUID
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

func New(docTypeStr, value string) (DocumentInterface, error) {
	docType := Type(strings.ToLower(docTypeStr))

	if docType != CPF && docType != CNPJ {
		return nil, errors.New("invalid document type")
	}

	cleanedValue := helpers.CleanString(value)

	if len(cleanedValue) < 11 {
		return nil, errors.New("invalid document value")
	}

	doc := &document{
		id:        uuid.Must(uuid.NewV7()),
		docType:   docType,
		value:     cleanedValue,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	return doc, nil
}

func (d *document) GetID() uuid.UUID        { return d.id }
func (d *document) GetType() Type           { return d.docType }
func (d *document) GetValue() string        { return d.value }
func (d *document) GetCreatedAt() time.Time { return d.createdAt }
func (d *document) GetUpdatedAt() time.Time { return d.updatedAt }
func (d *document) GetUserID() uuid.UUID    { return d.userID }
