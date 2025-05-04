package address

import (
	"time"

	"github.com/felipeversiane/donation-server/pkg/vo/uuid"
	"github.com/felipeversiane/donation-server/pkg/vo/zipcode"
)

type Address struct {
	ID           uuid.UUID       `json:"id"`
	ZipCode      zipcode.ZipCode `json:"zip_code"`
	Neighborhood string          `json:"neighborhood"`
	Street       string          `json:"street"`
	Number       string          `json:"number,omitempty"`
	Complement   string          `json:"complement,omitempty"`
	UserID       uuid.UUID       `json:"user_id,omitempty"`
	CityID       uuid.UUID       `json:"city_id"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

func New(zipCode, neighborhood, street, number, complement string, userID, cityID uuid.UUID) (*Address, error) {
	id, err := uuid.New()
	if err != nil {
		return nil, err
	}
	zip, err := zipcode.New(zipCode)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Address{
		ID:           id,
		ZipCode:      zip,
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
