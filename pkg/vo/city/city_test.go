package city

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCity(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		{"São Paulo", "São Paulo", true},
		{"Rio de Janeiro", "Rio de Janeiro", true},
		{"", "", false},
		{"12345", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			c, err := New(tt.input)
			if tt.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, c.String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
