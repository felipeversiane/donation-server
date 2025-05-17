package country

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{
			name:      "valid country name",
			input:     "Brasil",
			expectErr: false,
		},
		{
			name:      "empty name",
			input:     "",
			expectErr: true,
		},
		{
			name:      "name too short",
			input:     "A",
			expectErr: true,
		},
		{
			name:      "name too long",
			input:     string(make([]byte, 101)),
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := New(tt.input)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, c)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, c)
				assert.Equal(t, tt.input, c.Name)
			}
		})
	}
}
