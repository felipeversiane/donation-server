package user

import (
	"strings"
	"testing"

	"github.com/felipeversiane/donation-server/pkg/vo/document"
	"github.com/felipeversiane/donation-server/pkg/vo/email"
	"github.com/felipeversiane/donation-server/pkg/vo/password"
	"github.com/felipeversiane/donation-server/pkg/vo/phone"
	"github.com/stretchr/testify/assert"
)

func makeValidDependencies() (email.Email, password.Password, phone.Phone, document.Document) {
	email, _ := email.New("user@example.com")
	password, _ := password.New("StrongPass123!")
	phone, _ := phone.New("+5511999999999")
	document, _ := document.New("12345678901")
	return email, password, phone, document
}

func TestNewUser_Success(t *testing.T) {
	email, password, phone, document := makeValidDependencies()

	userObj, err := New(email, password, phone, document, "Valid Name", "avatar.png", false)

	assert.NoError(t, err)
	assert.NotNil(t, userObj)
	assert.Equal(t, "Valid Name", userObj.GetName())
}

func TestNewUser_EmptyName(t *testing.T) {
	email, password, phone, document := makeValidDependencies()

	userObj, err := New(email, password, phone, document, "", "", false)

	assert.Error(t, err)
	assert.Nil(t, userObj)
	assert.Equal(t, "name is required", err.Error())
}

func TestNewUser_NameTooLong(t *testing.T) {
	email, password, phone, document := makeValidDependencies()
	name := strings.Repeat("a", 101)

	userObj, err := New(email, password, phone, document, name, "", false)

	assert.Error(t, err)
	assert.Nil(t, userObj)
	assert.Equal(t, "name must be at most 100 characters", err.Error())
}

func TestNewUser_EmptyEmail(t *testing.T) {
	_, password, phone, document := makeValidDependencies()
	email := email.Email{}

	userObj, err := New(email, password, phone, document, "John Doe", "", false)

	assert.Error(t, err)
	assert.Nil(t, userObj)
	assert.Equal(t, "email is required", err.Error())
}

func TestNewUser_EmptyPassword(t *testing.T) {
	email, _, phone, document := makeValidDependencies()
	password := password.Password{}

	userObj, err := New(email, password, phone, document, "John Doe", "", false)

	assert.Error(t, err)
	assert.Nil(t, userObj)
	assert.Equal(t, "password is required", err.Error())
}

func TestNewUser_EmptyPhone(t *testing.T) {
	email, password, _, document := makeValidDependencies()
	phone := phone.Phone{}

	userObj, err := New(email, password, phone, document, "John Doe", "", false)

	assert.Error(t, err)
	assert.Nil(t, userObj)
	assert.Equal(t, "phone is required", err.Error())
}

func TestNewUser_EmptyDocument(t *testing.T) {
	email, password, phone, _ := makeValidDependencies()
	document := document.Document{}

	userObj, err := New(email, password, phone, document, "John Doe", "", false)

	assert.Error(t, err)
	assert.Nil(t, userObj)
	assert.Equal(t, "document is required", err.Error())
}
