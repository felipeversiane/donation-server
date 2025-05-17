package role

import (
	"errors"
	"strings"
)

type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
)

var validRoles = map[Role]bool{
	Admin: true,
	User:  true,
}

func New(value string) (Role, error) {
	role := Role(strings.ToLower(value))
	if !validRoles[role] {
		return "", errors.New("invalid role: must be 'admin' or 'user'")
	}
	return role, nil
}

func (r Role) String() string {
	return string(r)
}
