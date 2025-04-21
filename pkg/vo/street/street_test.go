package street

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStreet(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		{"Av. Paulista", "Av. Paulista", true},
		{"Rua dos Três Irmãos", "Rua dos Três Irmãos", true},
		{"", "", false},
		{"12345", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			s, err := New(tt.input)
			if tt.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, s.String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
