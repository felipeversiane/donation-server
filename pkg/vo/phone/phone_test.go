package phone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPhone(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		{"+5511998765432", "+5511998765432", true},
		{"1198765432", "1198765432", true},
		{"+1-800-555-5555", "+18005555555", true},
		{"(11) 98765-4321", "11987654321", true},
		{"+55 (11) 98765-4321", "+5511987654321", true},
		{"119876543", "", false},
		{"12345abcd", "", false},
		{"1234", "", false},
		{"(11) 98765-4321 ", "11987654321", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			phone, err := New(tt.input)
			if tt.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, phone.String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
