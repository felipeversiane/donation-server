package address

import (
	"time"

	"github.com/felipeversiane/donation-server/pkg/vo/zipcode"
	"github.com/google/uuid"
)

type address struct {
	id           uuid.UUID
	userID       uuid.UUID
	country      string
	zipCode      zipcode.ZipCode
	state        string
	city         string
	neighborhood string
	street       string
	number       string
	complement   string
	createdAt    time.Time
	updatedAt    time.Time
}

type AddressInterface interface {
	GetID() uuid.UUID
	GetUserID() uuid.UUID
	GetCountry() string
	GetZipCode() string
	GetState() string
	GetCity() string
	GetNeighborhood() string
	GetStreet() string
	GetNumber() string
	GetComplement() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

func New(
	userID uuid.UUID,
	zipCodeStr, country, state, city, neighborhood, street, number, complement string,
) (AddressInterface, error) {

	zipCode, err := zipcode.New(zipCodeStr)
	if err != nil {
		return nil, err
	}

	address := &address{
		id:           uuid.Must(uuid.NewV7()),
		userID:       userID,
		country:      country,
		zipCode:      zipCode,
		state:        state,
		city:         city,
		neighborhood: neighborhood,
		street:       street,
		number:       number,
		complement:   complement,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}

	return address, nil
}

func (a *address) GetID() uuid.UUID        { return a.id }
func (a *address) GetUserID() uuid.UUID    { return a.userID }
func (a *address) GetCountry() string      { return a.country }
func (a *address) GetZipCode() string      { return a.zipCode.String() }
func (a *address) GetState() string        { return a.state }
func (a *address) GetCity() string         { return a.city }
func (a *address) GetNeighborhood() string { return a.neighborhood }
func (a *address) GetStreet() string       { return a.street }
func (a *address) GetNumber() string       { return a.number }
func (a *address) GetComplement() string   { return a.complement }
func (a *address) GetCreatedAt() time.Time { return a.createdAt }
func (a *address) GetUpdatedAt() time.Time { return a.updatedAt }
