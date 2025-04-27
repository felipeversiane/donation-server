package country

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		{"Brazil", "Brazil", true},
		{"United States", "United States", true},
		{"Canada", "Canada", true},
		{"Germany", "Germany", true},
		{"Br", "", false},
		{"Brasília", "", false},
		{"", "", false},
		{"123", "", false},
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
