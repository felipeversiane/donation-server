package address

import (
	"errors"
	"time"

	"github.com/felipeversiane/donation-server/pkg/helpers"
	"github.com/felipeversiane/donation-server/pkg/vo/country"
	"github.com/felipeversiane/donation-server/pkg/vo/zipcode"
	"github.com/google/uuid"
)

type Address struct {
	ID           uuid.UUID       `json:"id"`
	Country      country.Country `json:"country"`
	ZipCode      zipcode.ZipCode `json:"zip_code"`
	State        string          `json:"state"`
	City         string          `json:"city"`
	Neighborhood string          `json:"neighborhood"`
	Street       string          `json:"street"`
	Number       string          `json:"number"`
	Complement   string          `json:"complement"`
	UserID       uuid.UUID       `json:"user_id"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

func New(
	zipCodeStr, countryStr, state, city, neighborhood, street, number, complement string,
	userID uuid.UUID,
) (*Address, error) {

	requiredFields := []struct {
		value string
		field string
	}{
		{state, "state"},
		{city, "city"},
		{street, "street"},
	}

	for _, f := range requiredFields {
		if err := helpers.ValidateRequired(f.value, f.field); err != nil {
			return nil, err
		}
	}

	maxLengthChecks := []struct {
		value string
		limit int
		field string
	}{
		{state, 100, "state"},
		{city, 100, "city"},
		{neighborhood, 100, "neighborhood"},
		{street, 255, "street"},
		{number, 20, "number"},
		{complement, 255, "complement"},
	}

	for _, f := range maxLengthChecks {
		if err := helpers.ValidateMaxLength(f.value, f.limit, f.field); err != nil {
			return nil, err
		}
	}

	if userID == uuid.Nil {
		return nil, errors.New("user_id is required")
	}

	c, err := country.New(countryStr)
	if err != nil {
		return nil, err
	}

	z, err := zipcode.New(zipCodeStr)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Address{
		ID:           uuid.Must(uuid.NewV7()),
		Country:      c,
		ZipCode:      z,
		State:        state,
		City:         city,
		Neighborhood: neighborhood,
		Street:       street,
		Number:       number,
		Complement:   complement,
		UserID:       userID,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}
