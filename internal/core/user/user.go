package user

import (
	"time"

	"github.com/felipeversiane/donation-server/internal/core/vo/email"
	"github.com/felipeversiane/donation-server/internal/core/vo/password"
	"github.com/felipeversiane/donation-server/internal/core/vo/role"
	"github.com/felipeversiane/donation-server/internal/core/vo/uuid"
)

type User struct {
	ID        uuid.UUID         `json:"id"`
	Email     email.Email       `json:"email"`
	Password  password.Password `json:"-"`
	AvatarID  *uuid.UUID        `json:"avatar_id,omitempty"`
	Role      role.Role         `json:"role"`
	Type      string            `json:"type"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func New(emailStr, passwordStr, roleStr, userType string, avatarID *uuid.UUID) (*User, error) {
	id, err := uuid.New()
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
		Type:      userType,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
