package zipcode

import (
	"errors"
	"regexp"

	"github.com/felipeversiane/donation-server/pkg/str"
)

type ZipCode struct {
	value string
}

func New(value string) (ZipCode, error) {
	cleaned := str.CleanString(value)

	if cleaned == "" {
		return ZipCode{}, errors.New("zipcode is required")
	}
	if match, _ := regexp.MatchString(`^\d{8}$`, cleaned); !match {
		return ZipCode{}, errors.New("invalid zip code format")
	}
	return ZipCode{value: cleaned}, nil
}

func (z ZipCode) String() string {
	return z.value
}
