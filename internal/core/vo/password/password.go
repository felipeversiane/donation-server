package password

import (
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrWeakPassword      = errors.New("password must have at least 8 characters, including uppercase, lowercase, number, and symbol")
	ErrHashingFailed     = errors.New("failed to hash password")
	ErrInvalidComparison = errors.New("invalid password")
)

type Password struct {
	hashed string
}

func New(value string) (Password, error) {
	if value == "" {
		return Password{}, errors.New("password is required")
	}
	if !isStrong(value) {
		return Password{}, ErrWeakPassword
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, ErrHashingFailed
	}

	return Password{hashed: string(hash)}, nil
}

func FromHashed(hashed string) Password {
	return Password{hashed: hashed}
}

func (p Password) Compare(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.hashed), []byte(plain))
	return err == nil
}

func (p Password) String() string {
	return p.hashed
}

func (p Password) Equals(other Password) bool {
	return p.hashed == other.hashed
}

func isStrong(password string) bool {
	var hasMinLen, hasUpper, hasLower, hasDigit, hasSpecial bool

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	return hasMinLen && hasUpper && hasLower && hasDigit && hasSpecial
}
