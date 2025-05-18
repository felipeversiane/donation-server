package address

import (
	"time"

	"github.com/felipeversiane/donation-server/internal/core/vo/zipcode"
	"github.com/felipeversiane/donation-server/pkg/field"
	"github.com/felipeversiane/donation-server/pkg/uuid"
)

type Address struct {
	ID           uuid.UUID
	ZipCode      zipcode.ZipCode
	Neighborhood string
	Street       string
	Number       string
	Complement   string
	UserID       uuid.UUID
	CityID       uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func New(zipCodeStr, neighborhood, street, number, complement string, userID, cityID uuid.UUID) (*Address, error) {
	if err := field.ValidateRequired(zipCodeStr, "zip code"); err != nil {
		return nil, err
	}
	if err := field.ValidateRequired(neighborhood, "neighborhood"); err != nil {
		return nil, err
	}
	if err := field.ValidateRequired(street, "street"); err != nil {
		return nil, err
	}
	if err := field.ValidateRequired(number, "number"); err != nil {
		return nil, err
	}

	zipCodeVO, err := zipcode.New(zipCodeStr)
	if err != nil {
		return nil, err
	}

	id, err := uuid.New()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Address{
		ID:           id,
		ZipCode:      zipCodeVO,
		Neighborhood: neighborhood,
		Street:       street,
		Number:       number,
		Complement:   complement,
		UserID:       userID,
		CityID:       cityID,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}
