package user

import (
	"time"

	"github.com/felipeversiane/donation-server/pkg/vo/document"
	"github.com/felipeversiane/donation-server/pkg/vo/email"
	"github.com/felipeversiane/donation-server/pkg/vo/password"
	"github.com/felipeversiane/donation-server/pkg/vo/phone"

	"github.com/google/uuid"
)

type user struct {
	id           uuid.UUID
	name         string
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
	email email.Email,
	password password.Password,
	phone phone.Phone,
	document document.Document,
	name, avatar string,
	isEnterprise bool,
) UserInterface {
	return &user{
		id:           uuid.Must(uuid.NewV7()),
		email:        email,
		password:     password,
		phone:        phone,
		name:         name,
		avatar:       avatar,
		document:     document,
		isEnterprise: isEnterprise,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}
}

func (u *user) GetID() uuid.UUID        { return u.id }
func (u *user) GetEmail() string        { return u.email.String() }
func (u *user) GetPassword() string     { return u.password.String() }
func (u *user) GetPhone() string        { return u.phone.String() }
func (u *user) GetName() string         { return u.name }
func (u *user) GetAvatar() string       { return u.avatar }
func (u *user) GetDocument() string     { return u.document.String() }
func (u *user) IsEnterprise() bool      { return u.isEnterprise }
func (u *user) GetCreatedAt() time.Time { return u.createdAt }
func (u *user) GetUpdatedAt() time.Time { return u.updatedAt }

func (u *user) ComparePassword(raw string) bool {
	return u.password.Compare(raw)
}
