package usertype

import (
	"errors"
	"strings"
)

type UserType string

const (
	Individual UserType = "individual"
	Company    UserType = "company"
)

var validTypes = map[UserType]bool{
	Individual: true,
	Company:    true,
}

func New(value string) (UserType, error) {
	userType := UserType(strings.ToLower(strings.TrimSpace(value)))
	if !validTypes[userType] {
		return "", errors.New("invalid user type: must be 'individual' or 'company'")
	}
	return userType, nil
}

func (u UserType) String() string {
	return string(u)
}
