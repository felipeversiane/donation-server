package email

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
		{"test@example.com", "test@example.com", true},
		{"test.email+1234@example.com", "test.email+1234@example.com", true},
		{"TEST@EXAMPLE.COM", "test@example.com", true},
		{"test@com", "", false},
		{"@example.com", "", false},
		{"test@.com", "", false},
		{"test@example..com", "", false},
		{" test@example.com ", "test@example.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			email, err := New(tt.input)
			if tt.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, email.String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
