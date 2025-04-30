package address

import (
	"testing"

	"github.com/felipeversiane/donation-server/pkg/vo/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	validUserID, _ := uuid.New()
	validCountry := "BR"
	validZipCode := "12345-678"
	validState := "São Paulo"
	validCity := "São Paulo"
	validNeighborhood := "Centro"
	validStreet := "Rua A"
	validNumber := "100"
	validComplement := "Apto 1"

	tests := []struct {
		name           string
		zipCodeStr     string
		countryStr     string
		state          string
		city           string
		neighborhood   string
		street         string
		number         string
		complement     string
		userID         uuid.UUID
		expectError    bool
		expectedErrMsg string
	}{
		{
			name:         "valid address",
			zipCodeStr:   validZipCode,
			countryStr:   validCountry,
			state:        validState,
			city:         validCity,
			neighborhood: validNeighborhood,
			street:       validStreet,
			number:       validNumber,
			complement:   validComplement,
			userID:       validUserID,
			expectError:  false,
		},
		{
			name:           "missing state",
			zipCodeStr:     validZipCode,
			countryStr:     validCountry,
			state:          "",
			city:           validCity,
			neighborhood:   validNeighborhood,
			street:         validStreet,
			number:         validNumber,
			complement:     validComplement,
			userID:         validUserID,
			expectError:    true,
			expectedErrMsg: "state is required",
		},
		{
			name:           "invalid user_id",
			zipCodeStr:     validZipCode,
			countryStr:     validCountry,
			state:          validState,
			city:           validCity,
			neighborhood:   validNeighborhood,
			street:         validStreet,
			number:         validNumber,
			complement:     validComplement,
			userID:         uuid.UUID{},  
			expectError:    true,
			expectedErrMsg: "user_id is required",
		},
		{
			name:           "invalid country",
			zipCodeStr:     validZipCode,
			countryStr:     "INVALID",
			state:          validState,
			city:           validCity,
			neighborhood:   validNeighborhood,
			street:         validStreet,
			number:         validNumber,
			complement:     validComplement,
			userID:         validUserID,
			expectError:    true,
			expectedErrMsg: "invalid country code",
		},
		{
			name:           "invalid zipcode",
			zipCodeStr:     "invalid_zip",
			countryStr:     validCountry,
			state:          validState,
			city:           validCity,
			neighborhood:   validNeighborhood,
			street:         validStreet,
			number:         validNumber,
			complement:     validComplement,
			userID:         validUserID,
			expectError:    true,
			expectedErrMsg: "invalid zipcode format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			address, err := New(
				tt.zipCodeStr,
				tt.countryStr,
				tt.state,
				tt.city,
				tt.neighborhood,
				tt.street,
				tt.number,
				tt.complement,
				tt.userID,
			)

			if tt.expectError {
				assert.Error(t, err)
				if tt.expectedErrMsg != "" {
					assert.Contains(t, err.Error(), tt.expectedErrMsg)
				}
				assert.Nil(t, address)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, address)
				assert.Equal(t, tt.state, address.State)
				assert.Equal(t, tt.city, address.City)
				assert.Equal(t, tt.neighborhood, address.Neighborhood)
				assert.Equal(t, tt.street, address.Street)
				assert.Equal(t, tt.number, address.Number)
				assert.Equal(t, tt.complement, address.Complement)
				assert.Equal(t, tt.userID.UUID(), address.UserID.UUID())
			}
		})
	}
}
