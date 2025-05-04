package zipcode

import (
	"errors"
	"regexp"
)

type ZipCode struct {
	value string
}

func New(value string) (ZipCode, error) {
	if value == "" {
		return ZipCode{}, errors.New("zipcode is required")
	}
	if match, _ := regexp.MatchString(`^\d{8}$`, value); !match {
		return ZipCode{}, errors.New("invalid zip code format")
	}
	return ZipCode{value: value}, nil
}

func (z ZipCode) String() string {
	return z.value
}
