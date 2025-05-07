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
	return hasMinLength(password) &&
		hasUppercase(password) &&
		hasLowercase(password) &&
		hasDigit(password) &&
		hasSpecialChar(password)
}

func hasMinLength(p string) bool {
	return len(p) >= 8
}

func hasUppercase(p string) bool {
	for _, ch := range p {
		if unicode.IsUpper(ch) {
			return true
		}
	}
	return false
}

func hasLowercase(p string) bool {
	for _, ch := range p {
		if unicode.IsLower(ch) {
			return true
		}
	}
	return false
}

func hasDigit(p string) bool {
	for _, ch := range p {
		if unicode.IsDigit(ch) {
			return true
		}
	}
	return false
}

func hasSpecialChar(p string) bool {
	for _, ch := range p {
		if unicode.IsPunct(ch) || unicode.IsSymbol(ch) {
			return true
		}
	}
	return false
}
