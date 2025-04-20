package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

type user struct {
	id           uuid.UUID
	firstName    string
	lastName     string
	email        string
	password     string
	phone        string
	avatar       string
	document     string
	isEnterprise bool
	createdAt    time.Time
	updatedAt    time.Time
}

type UserInterface interface {
	GetID() uuid.UUID
	GetEmail() string
	GetPassword() string
	GetPhone() string
	GetFirstName() string
	GetLastName() string
	GetAvatar() string
	GetDocument() string
	IsEnterprise() bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	ComparePassword(password string) bool
}

func NewUser(email, password, phone, firstName, lastName, avatar, document string, isEnterprise bool) UserInterface {
	u := &user{
		id:           uuid.Must(uuid.NewRandom()),
		email:        email,
		password:     hashPassword(password),
		phone:        phone,
		firstName:    firstName,
		lastName:     lastName,
		avatar:       avatar,
		document:     document,
		isEnterprise: isEnterprise,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}
	return u
}

func (u *user) GetID() uuid.UUID {
	return u.id
}

func (u *user) GetEmail() string {
	return u.email
}

func (u *user) GetPassword() string {
	return u.password
}

func (u *user) GetPhone() string {
	return u.phone
}

func (u *user) GetFirstName() string {
	return u.firstName
}

func (u *user) GetLastName() string {
	return u.lastName
}

func (u *user) GetAvatar() string {
	return u.avatar
}

func (u *user) GetDocument() string {
	return u.document
}

func (u *user) IsEnterprise() bool {
	return u.isEnterprise
}

func (u *user) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *user) GetUpdatedAt() time.Time {
	return u.updatedAt
}

func (u *user) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err == nil
}

func hashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}
