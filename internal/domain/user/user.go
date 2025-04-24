package user

import (
	"errors"
	"time"

	"github.com/felipeversiane/donation-server/pkg/helpers"
	"github.com/felipeversiane/donation-server/pkg/vo/email"
	"github.com/felipeversiane/donation-server/pkg/vo/password"
	"github.com/felipeversiane/donation-server/pkg/vo/phone"

	"github.com/google/uuid"
)

type user struct {
	id         uuid.UUID
	name       string
	email      email.Email
	password   password.Password
	phone      phone.Phone
	avatar     string
	documentID uuid.UUID
	addressID  uuid.UUID
	createdAt  time.Time
	updatedAt  time.Time
}

type UserInterface interface {
	GetID() uuid.UUID
	GetEmail() string
	GetPassword() string
	GetPhone() string
	GetName() string
	GetAvatar() string
	GetDocumentID() uuid.UUID
	GetAddressID() uuid.UUID
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	ComparePassword(raw string) bool
}

func New(
	name, emailStr, passwordStr, phoneStr, documentStr, avatar string,
	documentID, addressID uuid.UUID,
) (UserInterface, error) {

	if err := helpers.ValidateRequired(name, "name"); err != nil {
		return nil, err
	}

	if err := helpers.ValidateMaxLength(name, 100, "name"); err != nil {
		return nil, err
	}

	if avatar == "" {
		return nil, errors.New("avatar is required")
	}

	email, err := email.New(emailStr)
	if err != nil {
		return nil, err
	}

	password, err := password.New(passwordStr)
	if err != nil {
		return nil, err
	}

	phone, err := phone.New(phoneStr)
	if err != nil {
		return nil, err
	}

	user := &user{
		id:         uuid.Must(uuid.NewV7()),
		name:       name,
		email:      email,
		password:   password,
		phone:      phone,
		avatar:     avatar,
		documentID: documentID,
		addressID:  addressID,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}

	return user, nil
}

func (u *user) GetID() uuid.UUID         { return u.id }
func (u *user) GetEmail() string         { return u.email.String() }
func (u *user) GetPassword() string      { return u.password.String() }
func (u *user) GetPhone() string         { return u.phone.String() }
func (u *user) GetName() string          { return u.name }
func (u *user) GetAvatar() string        { return u.avatar }
func (u *user) GetDocumentID() uuid.UUID { return u.documentID }
func (u *user) GetAddressID() uuid.UUID  { return u.addressID }
func (u *user) GetCreatedAt() time.Time  { return u.createdAt }
func (u *user) GetUpdatedAt() time.Time  { return u.updatedAt }

func (u *user) ComparePassword(raw string) bool {
	return u.password.Compare(raw)
}
