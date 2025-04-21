package zipcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewZipCode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		{"12345678", "12345678", true},
		{"12345", "", false},
		{"ABCDE123", "", false},
		{"12345-678", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			zipCode, err := New(tt.input)
			if tt.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, zipCode.String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
