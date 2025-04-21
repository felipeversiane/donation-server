package name

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		isValid  bool
	}{
		{"Valid", "John Doe", "John Doe", true},
		{"Empty", "", "", false},
		{"TooLong", strings.Repeat("a", 101), "", false},
		{"Trimmed", "   Maria  ", "Maria", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := New(tt.input)

			if tt.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result.String())
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "name")
			}
		})
	}
}
