package address

import (
	"time"

	"github.com/felipeversiane/donation-server/pkg/helpers"
	"github.com/felipeversiane/donation-server/pkg/vo/country"
	"github.com/felipeversiane/donation-server/pkg/vo/zipcode"
	"github.com/google/uuid"
)

type address struct {
	id           uuid.UUID
	country      country.Country
	zipCode      zipcode.ZipCode
	state        string
	city         string
	neighborhood string
	street       string
	number       string
	complement   string
	userID       uuid.UUID
	createdAt    time.Time
	updatedAt    time.Time
}

type AddressInterface interface {
	GetID() uuid.UUID
	GetCountry() country.Country
	GetZipCode() zipcode.ZipCode
	GetState() string
	GetCity() string
	GetNeighborhood() string
	GetStreet() string
	GetNumber() string
	GetComplement() string
	GetUserID() uuid.UUID
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

func New(
	zipCodeStr, countryStr, state, city, neighborhood, street, number, complement string,
) (AddressInterface, error) {

	if err := helpers.ValidateRequired(state, "state"); err != nil {
		return nil, err
	}
	if err := helpers.ValidateRequired(city, "city"); err != nil {
		return nil, err
	}
	if err := helpers.ValidateRequired(street, "street"); err != nil {
		return nil, err
	}

	if err := helpers.ValidateMaxLength(state, 100, "state"); err != nil {
		return nil, err
	}
	if err := helpers.ValidateMaxLength(city, 100, "city"); err != nil {
		return nil, err
	}
	if err := helpers.ValidateMaxLength(neighborhood, 100, "neighborhood"); err != nil {
		return nil, err
	}
	if err := helpers.ValidateMaxLength(street, 255, "street"); err != nil {
		return nil, err
	}
	if err := helpers.ValidateMaxLength(number, 20, "number"); err != nil {
		return nil, err
	}
	if err := helpers.ValidateMaxLength(complement, 255, "complement"); err != nil {
		return nil, err
	}

	country, err := country.New(countryStr)
	if err != nil {
		return nil, err
	}

	zipCode, err := zipcode.New(zipCodeStr)
	if err != nil {
		return nil, err
	}

	address := &address{
		id:           uuid.Must(uuid.NewV7()),
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

func (a *address) GetID() uuid.UUID            { return a.id }
func (a *address) GetCountry() country.Country { return a.country }
func (a *address) GetZipCode() zipcode.ZipCode { return a.zipCode }
func (a *address) GetState() string            { return a.state }
func (a *address) GetCity() string             { return a.city }
func (a *address) GetNeighborhood() string     { return a.neighborhood }
func (a *address) GetStreet() string           { return a.street }
func (a *address) GetNumber() string           { return a.number }
func (a *address) GetComplement() string       { return a.complement }
func (a *address) GetCreatedAt() time.Time     { return a.createdAt }
func (a *address) GetUpdatedAt() time.Time     { return a.updatedAt }
func (a *address) GetUserID() uuid.UUID        { return a.userID }
