package address

import (
	"testing"

	"github.com/felipeversiane/donation-server/pkg/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	userID, _ := uuid.New()
	cityID, _ := uuid.New()

	tests := []struct {
		name         string
		zipCode      string
		expectedZip  string
		neighborhood string
		street       string
		number       string
		complement   string
		expectErr    bool
	}{
		{
			name:         "valid address",
			zipCode:      "74000-000",
			expectedZip:  "74000000",
			neighborhood: "Centro",
			street:       "Rua 1",
			number:       "123",
			complement:   "Ap 101",
			expectErr:    false,
		},
		{
			name:         "empty zip code",
			zipCode:      "",
			expectedZip:  "",
			neighborhood: "Centro",
			street:       "Rua 1",
			number:       "123",
			expectErr:    true,
		},
		{
			name:         "invalid zip code format",
			zipCode:      "ABC123",
			expectedZip:  "",
			neighborhood: "Centro",
			street:       "Rua 1",
			number:       "123",
			expectErr:    true,
		},
		{
			name:         "missing neighborhood",
			zipCode:      "74000-000",
			expectedZip:  "",
			neighborhood: "",
			street:       "Rua 1",
			number:       "123",
			expectErr:    true,
		},
		{
			name:         "missing street",
			zipCode:      "74000-000",
			expectedZip:  "",
			neighborhood: "Centro",
			street:       "",
			number:       "123",
			expectErr:    true,
		},
		{
			name:         "missing number",
			zipCode:      "74000-000",
			expectedZip:  "",
			neighborhood: "Centro",
			street:       "Rua 1",
			number:       "",
			expectErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addr, err := New(tt.zipCode, tt.neighborhood, tt.street, tt.number, tt.complement, userID, cityID)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, addr)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, addr)
				assert.Equal(t, tt.expectedZip, addr.ZipCode.String())
			}
		})
	}
}
