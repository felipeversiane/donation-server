package neighborhood

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNeighborhood(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		{"Centro", "Centro", true},
		{"Jardim Paulista", "Jardim Paulista", true},
		{"", "", false},
		{"12345", "", false},
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
