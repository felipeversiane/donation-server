package user

import (
	"time"

	"github.com/felipeversiane/donation-server/pkg/vo/document"
	"github.com/felipeversiane/donation-server/pkg/vo/email"
	"github.com/felipeversiane/donation-server/pkg/vo/name"
	"github.com/felipeversiane/donation-server/pkg/vo/password"
	"github.com/felipeversiane/donation-server/pkg/vo/phone"

	"github.com/google/uuid"
)

type user struct {
	id           uuid.UUID
	name         name.Name
	email        email.Email
	password     password.Password
	phone        phone.Phone
	avatar       string
	document     document.Document
	isEnterprise bool
	createdAt    time.Time
	updatedAt    time.Time
}

type UserInterface interface {
	GetID() uuid.UUID
	GetEmail() string
	GetPassword() string
	GetPhone() string
	GetName() string
	GetAvatar() string
	GetDocument() string
	IsEnterprise() bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	ComparePassword(raw string) bool
}

func New(
	nameStr, emailStr, passwordStr, phoneStr, documentStr, avatar string,
	isEnterprise bool,
) (UserInterface, error) {

	name, err := name.New(nameStr)
	if err != nil {
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

	document, err := document.New(documentStr)
	if err != nil {
		return nil, err
	}

	user := &user{
		id:           uuid.Must(uuid.NewV7()),
		name:         name,
		email:        email,
		password:     password,
		phone:        phone,
		avatar:       avatar,
		document:     document,
		isEnterprise: isEnterprise,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}

	return user, nil
}

func (u *user) GetID() uuid.UUID        { return u.id }
func (u *user) GetEmail() string        { return u.email.String() }
func (u *user) GetPassword() string     { return u.password.String() }
func (u *user) GetPhone() string        { return u.phone.String() }
func (u *user) GetName() string         { return u.name.String() }
func (u *user) GetAvatar() string       { return u.avatar }
func (u *user) GetDocument() string     { return u.document.String() }
func (u *user) IsEnterprise() bool      { return u.isEnterprise }
func (u *user) GetCreatedAt() time.Time { return u.createdAt }
func (u *user) GetUpdatedAt() time.Time { return u.updatedAt }

func (u *user) ComparePassword(raw string) bool {
	return u.password.Compare(raw)
}
