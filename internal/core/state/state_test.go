package state

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		inputName string
		inputUF   string
		expectErr bool
	}{
		{
			name:      "valid input",
			inputName: "Goiás",
			inputUF:   "GO",
			expectErr: false,
		},
		{
			name:      "empty name",
			inputName: "",
			inputUF:   "GO",
			expectErr: true,
		},
		{
			name:      "empty uf",
			inputName: "Goiás",
			inputUF:   "",
			expectErr: true,
		},
		{
			name:      "uf too short",
			inputName: "Goiás",
			inputUF:   "G",
			expectErr: true,
		},
		{
			name:      "uf too long",
			inputName: "Goiás",
			inputUF:   "GOA",
			expectErr: true,
		},
		{
			name:      "name too long",
			inputName: string(make([]byte, 101)),
			inputUF:   "GO",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state, err := New(tt.inputName, tt.inputUF, nil)
			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, state)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, state)
				assert.Equal(t, tt.inputName, state.Name)
				assert.Equal(t, tt.inputUF, state.UF)
			}
		})
	}
}
