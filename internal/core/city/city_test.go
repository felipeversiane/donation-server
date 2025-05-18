package city

import (
	"testing"

	"github.com/felipeversiane/donation-server/pkg/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	stateID, _ := uuid.New()

	tests := []struct {
		name      string
		input     string
		stateID   *uuid.UUID
		expectErr bool
	}{
		{
			name:      "valid city name",
			input:     "Anápolis",
			stateID:   &stateID,
			expectErr: false,
		},
		{
			name:      "empty name",
			input:     "",
			stateID:   &stateID,
			expectErr: true,
		},
		{
			name:      "name too short",
			input:     "A",
			stateID:   &stateID,
			expectErr: true,
		},
		{
			name:      "name too long",
			input:     string(make([]byte, 101)),
			stateID:   &stateID,
			expectErr: true,
		},
		{
			name:      "valid city with nil stateID",
			input:     "Goiânia",
			stateID:   nil,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			city, err := New(tt.input, tt.stateID)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, city)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, city)
				assert.Equal(t, tt.input, city.Name)
			}
		})
	}
}
