package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name         string
		inputName    string
		inputEmail   string
		inputPass    string
		inputPhone   string
		inputDoc     string
		inputAvatar  string
		isEnterprise bool
		expectError  bool
	}{
		{
			name:         "Valid user",
			inputName:    "Alice Smith",
			inputEmail:   "alice@example.com",
			inputPass:    "StrongPass123!",
			inputPhone:   "+5511988887777",
			inputDoc:     "12345678900",
			inputAvatar:  "https://example.com/avatar.png",
			isEnterprise: false,
			expectError:  false,
		},
		{
			name:         "Invalid email",
			inputName:    "Alice",
			inputEmail:   "alice@@com",
			inputPass:    "StrongPass123!",
			inputPhone:   "+5511988887777",
			inputDoc:     "12345678900",
			inputAvatar:  "",
			isEnterprise: false,
			expectError:  true,
		},
		{
			name:         "Invalid password",
			inputName:    "Bob",
			inputEmail:   "bob@example.com",
			inputPass:    "123",
			inputPhone:   "+5511988887777",
			inputDoc:     "12345678900",
			inputAvatar:  "",
			isEnterprise: false,
			expectError:  true,
		},
		{
			name:         "Empty name",
			inputName:    "",
			inputEmail:   "bob@example.com",
			inputPass:    "StrongPass123!",
			inputPhone:   "+5511988887777",
			inputDoc:     "12345678900",
			inputAvatar:  "",
			isEnterprise: false,
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := New(tt.inputName, tt.inputEmail, tt.inputPass, tt.inputPhone, tt.inputDoc, tt.inputAvatar, tt.isEnterprise)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, u)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, u)
				assert.Equal(t, tt.inputName, u.GetName())
				assert.Equal(t, tt.inputEmail, u.GetEmail())
				assert.Equal(t, tt.inputPhone, u.GetPhone())
				assert.Equal(t, tt.inputDoc, u.GetDocument())
				assert.Equal(t, tt.inputAvatar, u.GetAvatar())
				assert.Equal(t, tt.isEnterprise, u.IsEnterprise())
				assert.True(t, u.ComparePassword(tt.inputPass))
			}
		})
	}
}
