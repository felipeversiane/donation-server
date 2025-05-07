package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		input   string
		isValid bool
	}{
		{"Password123!", true},
		{"password", false},
		{"Password", false},
		{"1234567890", false},
		{"!@#password123", true},
		{"1234", false},
		{"Password1234", true},
		{" pass word ", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			password, err := New(tt.input)
			if tt.isValid {
				assert.NoError(t, err)
				assert.NotEmpty(t, password.String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
