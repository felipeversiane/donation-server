package document

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDocument(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		{"12345678909", "12345678909", true},
		{"00000000000", "", false},
		{"12345678901234", "12345678901234", true},
		{"00000000000000", "", false},
		{"12.345.678/0001-95", "12345678000195", true},
		{"1234", "", false},
		{"1234567890", "", false},
		{"1234567890a", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			doc, err := New(tt.input)
			if tt.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, doc.String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
