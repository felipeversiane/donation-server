package address

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewAddress(t *testing.T) {
	tests := []struct {
		userID             uuid.UUID
		zipCodeStr         string
		countryStr         string
		stateStr           string
		cityStr            string
		neighborhoodStr    string
		streetStr          string
		numberStr          string
		complementStr      string
		expectedCountry    string
		expectedState      string
		expectedCity       string
		expectedStreet     string
		expectedNumber     string
		expectedComplement string
		isValid            bool
	}{
		{
			userID:             uuid.New(),
			zipCodeStr:         "12345-678",
			countryStr:         "Brazil",
			stateStr:           "SP",
			cityStr:            "São Paulo",
			neighborhoodStr:    "Centro",
			streetStr:          "Av. Paulista",
			numberStr:          "1000",
			complementStr:      "Apto 32",
			expectedCountry:    "Brazil",
			expectedState:      "SP",
			expectedCity:       "São Paulo",
			expectedStreet:     "Av. Paulista",
			expectedNumber:     "1000",
			expectedComplement: "Apto 32",
			isValid:            true,
		},
		{
			userID:             uuid.New(),
			zipCodeStr:         "98765-432",
			countryStr:         "USA",
			stateStr:           "NY",
			cityStr:            "New York",
			neighborhoodStr:    "Brooklyn",
			streetStr:          "5th Avenue",
			numberStr:          "200",
			complementStr:      "Suite 10B",
			expectedCountry:    "USA",
			expectedState:      "NY",
			expectedCity:       "New York",
			expectedStreet:     "5th Avenue",
			expectedNumber:     "200",
			expectedComplement: "Suite 10B",
			isValid:            true,
		},
		{
			userID:             uuid.New(),
			zipCodeStr:         "54321-987",
			countryStr:         "",
			stateStr:           "",
			cityStr:            "",
			neighborhoodStr:    "",
			streetStr:          "",
			numberStr:          "",
			complementStr:      "",
			expectedCountry:    "",
			expectedState:      "",
			expectedCity:       "",
			expectedStreet:     "",
			expectedNumber:     "",
			expectedComplement: "",
			isValid:            false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.zipCodeStr, func(t *testing.T) {
			address, err := New(tt.userID, tt.zipCodeStr, tt.countryStr, tt.stateStr, tt.cityStr, tt.neighborhoodStr, tt.streetStr, tt.numberStr, tt.complementStr)

			if tt.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCountry, address.GetCountry().String())
				assert.Equal(t, tt.expectedState, address.GetState().String())
				assert.Equal(t, tt.expectedCity, address.GetCity().String())
				assert.Equal(t, tt.expectedStreet, address.GetStreet().String())
				assert.Equal(t, tt.expectedNumber, address.GetNumber().String())
				assert.Equal(t, tt.expectedComplement, address.GetComplement().String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
