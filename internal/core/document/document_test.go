package document

import (
	"testing"

	"github.com/felipeversiane/donation-server/internal/core/vo/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	validUUID, _ := uuid.New()

	tests := []struct {
		name      string
		value     string
		docType   string
		userID    uuid.UUID
		expectErr bool
	}{
		{
			name:      "valid CPF",
			value:     "123.456.789-01",
			docType:   "cpf",
			userID:    validUUID,
			expectErr: false,
		},
		{
			name:      "valid CNPJ",
			value:     "12.345.678/0001-99",
			docType:   "cnpj",
			userID:    validUUID,
			expectErr: false,
		},
		{
			name:      "invalid CPF format",
			value:     "123456789",
			docType:   "cpf",
			userID:    validUUID,
			expectErr: true,
		},
		{
			name:      "invalid CNPJ format",
			value:     "123456789012",
			docType:   "cnpj",
			userID:    validUUID,
			expectErr: true,
		},
		{
			name:      "invalid document type",
			value:     "12345678901",
			docType:   "rg",
			userID:    validUUID,
			expectErr: true,
		},
		{
			name:      "empty value",
			value:     "",
			docType:   "cpf",
			userID:    validUUID,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := New(tt.value, tt.docType, tt.userID)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, doc)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, doc)
				assert.Equal(t, tt.docType, doc.Type.String())
			}
		})
	}
}
