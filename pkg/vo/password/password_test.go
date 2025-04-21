package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		{"Password123!", "Password123!", true},
		{"password", "", false},
		{"Password", "", false},
		{"1234567890", "", false},
		{"!@#password123", "!@#password123", true},
		{"1234", "", false},
		{"Password1234", "Password1234", true},
		{" pass word ", "password", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			password, err := New(tt.input)
			if tt.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, password.String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
