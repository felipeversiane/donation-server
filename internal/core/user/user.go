package user

import (
	"time"

	"github.com/felipeversiane/donation-server/pkg/vo/email"
	"github.com/felipeversiane/donation-server/pkg/vo/password"
	"github.com/felipeversiane/donation-server/pkg/vo/uuid"
)

type User struct {
	ID        uuid.UUID         `json:"id"`
	Email     email.Email       `json:"email"`
	Password  password.Password `json:"-"`
	AvatarID  *uuid.UUID        `json:"avatar_id,omitempty"`
	Role      string            `json:"role"`
	Type      string            `json:"type"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func New(emailStr, passwordStr, role, userType string, avatarID *uuid.UUID) (*User, error) {
	id, err := uuid.New()
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

	now := time.Now()

	return &User{
		ID:        id,
		Email:     email,
		Password:  password,
		AvatarID:  avatarID,
		Role:      role,
		Type:      userType,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
