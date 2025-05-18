package user

import (
	"time"

	"github.com/felipeversiane/donation-server/internal/core/vo/email"
	"github.com/felipeversiane/donation-server/internal/core/vo/password"
	"github.com/felipeversiane/donation-server/internal/core/vo/role"
	"github.com/felipeversiane/donation-server/internal/core/vo/usertype"
	"github.com/felipeversiane/donation-server/internal/core/vo/uuid"
)

type User struct {
	ID        uuid.UUID
	Email     email.Email
	Password  password.Password
	AvatarID  *uuid.UUID
	Role      role.Role
	Type      usertype.UserType
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(emailStr, passwordStr, roleStr, userTypeStr string, avatarID *uuid.UUID) (*User, error) {
	id, err := uuid.New()
	if err != nil {
		return nil, err
	}

	userTypeVO, err := usertype.New(userTypeStr)
	if err != nil {
		return nil, err
	}

	emailVO, err := email.New(emailStr)
	if err != nil {
		return nil, err
	}

	passwordVO, err := password.New(passwordStr)
	if err != nil {
		return nil, err
	}

	roleVO, err := role.New(roleStr)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &User{
		ID:        id,
		Email:     emailVO,
		Password:  passwordVO,
		AvatarID:  avatarID,
		Role:      roleVO,
		Type:      userTypeVO,
		IsActive:  false,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
