package complement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewComplement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		{"Apto 32", "Apto 32", true},
		{"", "", true},
		{"Bloco A", "Bloco A", true},
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
