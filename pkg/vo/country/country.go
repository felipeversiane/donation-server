package country

import (
	"errors"

	"github.com/pariz/gountries"
)

type Country struct {
	value string
}

func New(input string) (Country, error) {
	query := gountries.New()
	_, err := query.FindCountryByName(input)
	if err != nil {
		return Country{}, errors.New("invalid country name")
	}

	return Country{value: input}, nil
}

func (c Country) String() string {
	return c.value
}
