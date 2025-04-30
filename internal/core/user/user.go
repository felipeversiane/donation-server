package user

import (
	"time"

	"github.com/felipeversiane/donation-server/pkg/helpers"
	"github.com/felipeversiane/donation-server/pkg/vo/email"
	"github.com/felipeversiane/donation-server/pkg/vo/password"
	"github.com/felipeversiane/donation-server/pkg/vo/phone"

	"github.com/felipeversiane/donation-server/pkg/vo/uuid"
	"github.com/pkg/errors"
)

type User struct {
	ID        uuid.UUID         `json:"id"`
	Name      string            `json:"name"`
	Email     email.Email       `json:"email"`
	Password  password.Password `json:"-"`
	Phone     phone.Phone       `json:"phone"`
	Avatar    string            `json:"avatar"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func New(
	name, emailStr, passwordStr, phoneStr, avatar string,
) (*User, error) {
	if err := helpers.ValidateRequired(name, "name"); err != nil {
		return nil, errors.Wrap(err, "validating name")
	}
	if err := helpers.ValidateMinLength(name, 2, "name"); err != nil {
		return nil, errors.Wrap(err, "validating minimum name length")
	}
	if err := helpers.ValidateMaxLength(name, 100, "name"); err != nil {
		return nil, errors.Wrap(err, "validating maximum name length")
	}

	if err := helpers.ValidateRequired(avatar, "avatar"); err != nil {
		return nil, errors.Wrap(err, "validating avatar")
	}

	emailVO, err := email.New(emailStr)
	if err != nil {
		return nil, errors.Wrap(err, "creating email")
	}

	passVO, err := password.New(passwordStr)
	if err != nil {
		return nil, errors.Wrap(err, "creating password")
	}

	phoneVO, err := phone.New(phoneStr)
	if err != nil {
		return nil, errors.Wrap(err, "creating phone")
	}

	id, err := uuid.New()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &User{
		ID:        id,
		Name:      name,
		Email:     emailVO,
		Password:  passVO,
		Phone:     phoneVO,
		Avatar:    avatar,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (u *User) Update(name, avatar string, phone phone.Phone) error {
	if err := helpers.ValidateRequired(name, "name"); err != nil {
		return errors.Wrap(err, "validating name")
	}
	if err := helpers.ValidateMinLength(name, 2, "name"); err != nil {
		return errors.Wrap(err, "validating minimum name length")
	}
	if err := helpers.ValidateMaxLength(name, 100, "name"); err != nil {
		return errors.Wrap(err, "validating maximum name length")
	}
	if err := helpers.ValidateRequired(avatar, "avatar"); err != nil {
		return errors.Wrap(err, "validating avatar")
	}

	u.Name = name
	u.Avatar = avatar
	u.Phone = phone
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) ChangePassword(newPasswordStr string) error {
	newPassword, err := password.New(newPasswordStr)
	if err != nil {
		return errors.Wrap(err, "validating new password")
	}

	u.Password = newPassword
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) ChangeEmail(newEmailStr string) error {
	newEmail, err := email.New(newEmailStr)
	if err != nil {
		return errors.Wrap(err, "validating new email")
	}

	u.Email = newEmail
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) ComparePassword(raw string) bool {
	return u.Password.Compare(raw)
}
