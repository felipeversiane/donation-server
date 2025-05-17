package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		password  string
		role      string
		userType  string
		expectErr bool
	}{
		{
			name:      "valid user",
			email:     "test@example.com",
			password:  "Password123!",
			role:      "admin",
			userType:  "individual",
			expectErr: false,
		},
		{
			name:      "invalid email",
			email:     "invalid-email",
			password:  "Password123!",
			role:      "admin",
			userType:  "individual",
			expectErr: true,
		},
		{
			name:      "weak password",
			email:     "test@example.com",
			password:  "invalid-password",
			role:      "admin",
			userType:  "individual",
			expectErr: true,
		},
		{
			name:      "invalid role",
			email:     "test@example.com",
			password:  "Password123!",
			role:      "manager",
			userType:  "individual",
			expectErr: true,
		},
		{
			name:      "invalid user type",
			email:     "test@example.com",
			password:  "Password123!",
			role:      "user",
			userType:  "pessoa física",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := New(tt.email, tt.password, tt.role, tt.userType, nil)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, u)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, u)
				assert.Equal(t, tt.email, u.Email.String())
				assert.Equal(t, tt.role, u.Role.String())
				assert.Equal(t, tt.userType, u.Type.String())
				assert.NotEmpty(t, u.ID.String())
			}
		})
	}
}
