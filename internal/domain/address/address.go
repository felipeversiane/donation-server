package address

import (
	"time"

	"github.com/felipeversiane/donation-server/pkg/vo/city"
	"github.com/felipeversiane/donation-server/pkg/vo/complement"
	"github.com/felipeversiane/donation-server/pkg/vo/country"
	"github.com/felipeversiane/donation-server/pkg/vo/neighborhood"
	"github.com/felipeversiane/donation-server/pkg/vo/number"
	"github.com/felipeversiane/donation-server/pkg/vo/state"
	"github.com/felipeversiane/donation-server/pkg/vo/street"
	"github.com/felipeversiane/donation-server/pkg/vo/zipcode"
	"github.com/google/uuid"
)

type address struct {
	id           uuid.UUID
	userID       uuid.UUID
	country      country.Country
	zipCode      zipcode.ZipCode
	state        state.State
	city         city.City
	neighborhood neighborhood.Neighborhood
	street       street.Street
	number       number.Number
	complement   complement.Complement
	createdAt    time.Time
	updatedAt    time.Time
}

type AddressInterface interface {
	GetID() uuid.UUID
	GetUserID() uuid.UUID
	GetCountry() country.Country
	GetZipCode() zipcode.ZipCode
	GetState() state.State
	GetCity() city.City
	GetNeighborhood() neighborhood.Neighborhood
	GetStreet() street.Street
	GetNumber() number.Number
	GetComplement() complement.Complement
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

func New(
	userID uuid.UUID,
	zipCodeStr, countryStr, stateStr, cityStr, neighborhoodStr, streetStr, numberStr, complementStr string,
) (AddressInterface, error) {

	zipCode, err := zipcode.New(zipCodeStr)
	if err != nil {
		return nil, err
	}

	country, err := country.New(countryStr)
	if err != nil {
		return nil, err
	}

	state, err := state.New(stateStr)
	if err != nil {
		return nil, err
	}

	city, err := city.New(cityStr)
	if err != nil {
		return nil, err
	}

	neighborhood, err := neighborhood.New(neighborhoodStr)
	if err != nil {
		return nil, err
	}

	street, err := street.New(streetStr)
	if err != nil {
		return nil, err
	}

	number, err := number.New(numberStr)
	if err != nil {
		return nil, err
	}

	complement, err := complement.New(complementStr)
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

func (a *address) GetID() uuid.UUID                           { return a.id }
func (a *address) GetUserID() uuid.UUID                       { return a.userID }
func (a *address) GetCountry() country.Country                { return a.country }
func (a *address) GetZipCode() zipcode.ZipCode                { return a.zipCode }
func (a *address) GetState() state.State                      { return a.state }
func (a *address) GetCity() city.City                         { return a.city }
func (a *address) GetNeighborhood() neighborhood.Neighborhood { return a.neighborhood }
func (a *address) GetStreet() street.Street                   { return a.street }
func (a *address) GetNumber() number.Number                   { return a.number }
func (a *address) GetComplement() complement.Complement       { return a.complement }
func (a *address) GetCreatedAt() time.Time                    { return a.createdAt }
func (a *address) GetUpdatedAt() time.Time                    { return a.updatedAt }
