package user

import (
	"time"

	"github.com/felipeversiane/donation-server/pkg/helpers"
	"github.com/felipeversiane/donation-server/pkg/vo/email"
	"github.com/felipeversiane/donation-server/pkg/vo/password"
	"github.com/felipeversiane/donation-server/pkg/vo/phone"

	"github.com/google/uuid"
)

type user struct {
	id        uuid.UUID
	name      string
	email     email.Email
	password  password.Password
	phone     phone.Phone
	avatar    string
	createdAt time.Time
	updatedAt time.Time
}

type UserInterface interface {
	GetID() uuid.UUID
	GetEmail() email.Email
	GetPassword() password.Password
	GetPhone() phone.Phone
	GetName() string
	GetAvatar() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	ComparePassword(raw string) bool
}

func New(
	name, emailStr, passwordStr, phoneStr, documentStr, avatar string,
) (UserInterface, error) {

	if err := helpers.ValidateRequired(name, "name"); err != nil {
		return nil, err
	}

	if err := helpers.ValidateMaxLength(name, 100, "name"); err != nil {
		return nil, err
	}

	if err := helpers.ValidateRequired(avatar, "avatar"); err != nil {
		return nil, err
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
		id:        uuid.Must(uuid.NewV7()),
		name:      name,
		email:     email,
		password:  password,
		phone:     phone,
		avatar:    avatar,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	return user, nil
}

func (u *user) GetID() uuid.UUID               { return u.id }
func (u *user) GetEmail() email.Email          { return u.email }
func (u *user) GetPassword() password.Password { return u.password }
func (u *user) GetPhone() phone.Phone          { return u.phone }
func (u *user) GetName() string                { return u.name }
func (u *user) GetAvatar() string              { return u.avatar }
func (u *user) GetCreatedAt() time.Time        { return u.createdAt }
func (u *user) GetUpdatedAt() time.Time        { return u.updatedAt }

func (u *user) ComparePassword(raw string) bool {
	return u.password.Compare(raw)
}
