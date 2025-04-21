package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		{"123A", "123A", true},
		{"456", "456", true},
		{"", "", false},
		{"A123", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			n, err := New(tt.input)
			if tt.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, n.String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
