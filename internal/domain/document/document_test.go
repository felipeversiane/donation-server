package document

import (
	"testing"
	"time"

	"github.com/felipeversiane/donation-server/pkg/helpers"
	"github.com/felipeversiane/donation-server/pkg/vo/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	validUserID, _ := uuid.New()

	tests := []struct {
		name         string
		docType      string
		value        string
		userID       uuid.UUID
		expectedErr  bool
		errorMessage string
	}{
		{
			name:        "valid CPF",
			docType:     "cpf",
			value:       "12345678901",
			userID:      validUserID,
			expectedErr: false,
		},
		{
			name:        "valid CPF with formatting",
			docType:     "cpf",
			value:       "123.456.789-01",
			userID:      validUserID,
			expectedErr: false,
		},
		{
			name:        "valid CNPJ",
			docType:     "cnpj",
			value:       "12345678000190",
			userID:      validUserID,
			expectedErr: false,
		},
		{
			name:        "valid CNPJ with formatting",
			docType:     "cnpj",
			value:       "12.345.678/0001-90",
			userID:      validUserID,
			expectedErr: false,
		},
		{
			name:         "invalid document type",
			docType:      "rg",
			value:        "123456789",
			userID:       validUserID,
			expectedErr:  true,
			errorMessage: "invalid document type",
		},
		{
			name:         "short CPF",
			docType:      "cpf",
			value:        "1234567890",
			userID:       validUserID,
			expectedErr:  true,
			errorMessage: "CPF must have 11 digits",
		},
		{
			name:         "short CNPJ",
			docType:      "cnpj",
			value:        "1234567890123",
			userID:       validUserID,
			expectedErr:  true,
			errorMessage: "CNPJ must have 14 digits",
		},
		{
			name:         "nil user ID",
			docType:      "cpf",
			value:        "12345678901",
			userID:       uuid.UUID{},
			expectedErr:  true,
			errorMessage: "user ID is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc, err := New(tt.docType, tt.value, tt.userID)

			if tt.expectedErr {
				assert.Error(t, err)
				if tt.errorMessage != "" {
					assert.Contains(t, err.Error(), tt.errorMessage)
				}
				assert.Nil(t, doc)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, doc)

				if tt.docType == "cpf" {
					assert.Equal(t, CPF, doc.Type)
				} else {
					assert.Equal(t, CNPJ, doc.Type)
				}

				if tt.docType == "cpf" {
					assert.Len(t, doc.Value, 11)
				} else {
					assert.Len(t, doc.Value, 14)
				}

				assert.Equal(t, tt.userID, doc.UserID)
				assert.NotEqual(t, uuid.UUID{}, doc.ID)
				assert.WithinDuration(t, time.Now(), doc.CreatedAt, 2*time.Second)
				assert.WithinDuration(t, time.Now(), doc.UpdatedAt, 2*time.Second)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		docType  Type
		value    string
		expected bool
	}{
		{
			name:     "valid CPF length",
			docType:  CPF,
			value:    "12345678901",
			expected: true,
		},
		{
			name:     "invalid CPF length",
			docType:  CPF,
			value:    "123456",
			expected: false,
		},
		{
			name:     "valid CNPJ length",
			docType:  CNPJ,
			value:    "12345678000190",
			expected: true,
		},
		{
			name:     "invalid CNPJ length",
			docType:  CNPJ,
			value:    "123456",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := &Document{
				Type:  tt.docType,
				Value: tt.value,
			}

			result := doc.IsValid()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUpdate(t *testing.T) {
	validUserID, _ := uuid.New()

	cpfDoc, _ := New("cpf", "12345678901", validUserID)
	cnpjDoc, _ := New("cnpj", "12345678000190", validUserID)

	originalCPFTime := cpfDoc.UpdatedAt
	originalCNPJTime := cnpjDoc.UpdatedAt

	time.Sleep(10 * time.Millisecond)

	tests := []struct {
		name         string
		doc          *Document
		newValue     string
		expectedErr  bool
		errorMessage string
	}{
		{
			name:        "update CPF to valid value",
			doc:         cpfDoc,
			newValue:    "10987654321",
			expectedErr: false,
		},
		{
			name:        "update CPF with formatting",
			doc:         cpfDoc,
			newValue:    "109.876.543-21",
			expectedErr: false,
		},
		{
			name:         "update CPF to short value",
			doc:          cpfDoc,
			newValue:     "1234567890",
			expectedErr:  true,
			errorMessage: "CPF must have 11 digits",
		},
		{
			name:        "update CNPJ to valid value",
			doc:         cnpjDoc,
			newValue:    "09876543000121",
			expectedErr: false,
		},
		{
			name:         "update CNPJ to short value",
			doc:          cnpjDoc,
			newValue:     "1234567890123",
			expectedErr:  true,
			errorMessage: "CNPJ must have 14 digits",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDoc := &Document{
				ID:        tt.doc.ID,
				Type:      tt.doc.Type,
				Value:     tt.doc.Value,
				UserID:    tt.doc.UserID,
				CreatedAt: tt.doc.CreatedAt,
				UpdatedAt: tt.doc.UpdatedAt,
			}

			err := testDoc.Update(tt.newValue)

			if tt.expectedErr {
				assert.Error(t, err)
				if tt.errorMessage != "" {
					assert.Contains(t, err.Error(), tt.errorMessage)
				}
			} else {
				assert.NoError(t, err)

				cleanedValue := helpers.CleanString(tt.newValue)
				assert.Equal(t, cleanedValue, testDoc.Value)

				if testDoc.Type == CPF {
					assert.True(t, testDoc.UpdatedAt.After(originalCPFTime))
				} else {
					assert.True(t, testDoc.UpdatedAt.After(originalCNPJTime))
				}
			}
		})
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		name     string
		docType  Type
		value    string
		expected string
	}{
		{
			name:     "format CPF",
			docType:  CPF,
			value:    "12345678901",
			expected: "123.456.789-01",
		},
		{
			name:     "format incomplete CPF",
			docType:  CPF,
			value:    "123456",
			expected: "123456",
		},
		{
			name:     "format CNPJ",
			docType:  CNPJ,
			value:    "12345678000190",
			expected: "12.345.678/0001-90",
		},
		{
			name:     "format incomplete CNPJ",
			docType:  CNPJ,
			value:    "123456",
			expected: "123456",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := &Document{
				Type:  tt.docType,
				Value: tt.value,
			}

			result := doc.Format()
			assert.Equal(t, tt.expected, result)
		})
	}
}
