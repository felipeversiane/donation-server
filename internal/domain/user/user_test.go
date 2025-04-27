package user

import (
	"testing"
	"time"

	"github.com/felipeversiane/donation-server/pkg/vo/phone"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name          string
		inputName     string
		inputEmail    string
		inputPassword string
		inputPhone    string
		inputAvatar   string
		expectedErr   bool
		errorMessage  string
	}{
		{
			name:          "valid user",
			inputName:     "John Doe",
			inputEmail:    "john@example.com",
			inputPassword: "StrongPass123!",
			inputPhone:    "+5511999999999",
			inputAvatar:   "avatar.jpg",
			expectedErr:   false,
		},
		{
			name:          "empty name",
			inputName:     "",
			inputEmail:    "john@example.com",
			inputPassword: "StrongPass123!",
			inputPhone:    "+5511999999999",
			inputAvatar:   "avatar.jpg",
			expectedErr:   true,
			errorMessage:  "name is required",
		},
		{
			name:          "name too short",
			inputName:     "A",
			inputEmail:    "john@example.com",
			inputPassword: "StrongPass123!",
			inputPhone:    "+5511999999999",
			inputAvatar:   "avatar.jpg",
			expectedErr:   true,
			errorMessage:  "minimum name length",
		},
		{
			name:          "name too long",
			inputName:     string(make([]byte, 101)),  
			inputEmail:    "john@example.com",
			inputPassword: "StrongPass123!",
			inputPhone:    "+5511999999999",
			inputAvatar:   "avatar.jpg",
			expectedErr:   true,
			errorMessage:  "maximum name length",
		},
		{
			name:          "empty avatar",
			inputName:     "John Doe",
			inputEmail:    "john@example.com",
			inputPassword: "StrongPass123!",
			inputPhone:    "+5511999999999",
			inputAvatar:   "",
			expectedErr:   true,
			errorMessage:  "avatar is required",
		},
		{
			name:          "invalid email",
			inputName:     "John Doe",
			inputEmail:    "invalid-email",
			inputPassword: "StrongPass123!",
			inputPhone:    "+5511999999999",
			inputAvatar:   "avatar.jpg",
			expectedErr:   true,
			errorMessage:  "creating email",
		},
		{
			name:          "weak password",
			inputName:     "John Doe",
			inputEmail:    "john@example.com",
			inputPassword: "123",
			inputPhone:    "+5511999999999",
			inputAvatar:   "avatar.jpg",
			expectedErr:   true,
			errorMessage:  "creating password",
		},
		{
			name:          "invalid phone",
			inputName:     "John Doe",
			inputEmail:    "john@example.com",
			inputPassword: "StrongPass123!",
			inputPhone:    "invalid",
			inputAvatar:   "avatar.jpg",
			expectedErr:   true,
			errorMessage:  "creating phone",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := New(
				tt.inputName,
				tt.inputEmail,
				tt.inputPassword,
				tt.inputPhone,
				tt.inputAvatar,
			)

			if tt.expectedErr {
				assert.Error(t, err)
				if tt.errorMessage != "" {
					assert.Contains(t, err.Error(), tt.errorMessage)
				}
				assert.Nil(t, u)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, u)
				assert.Equal(t, tt.inputName, u.Name)
				assert.Equal(t, tt.inputEmail, u.Email.String())
				assert.Equal(t, tt.inputPhone, u.Phone.String())
				assert.Equal(t, tt.inputAvatar, u.Avatar)
				assert.True(t, u.ComparePassword(tt.inputPassword))
				assert.NotEqual(t, uuid.Nil, u.ID)
				assert.WithinDuration(t, time.Now(), u.CreatedAt, 2*time.Second)
				assert.WithinDuration(t, time.Now(), u.UpdatedAt, 2*time.Second)
			}
		})
	}
}

func TestComparePassword(t *testing.T) {
	tests := []struct {
		name            string
		initialPassword string
		comparePassword string
		expected        bool
	}{
		{
			name:            "matching password",
			initialPassword: "StrongPass123!",
			comparePassword: "StrongPass123!",
			expected:        true,
		},
		{
			name:            "non-matching password",
			initialPassword: "StrongPass123!",
			comparePassword: "DifferentPass456!",
			expected:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := New(
				"Test User",
				"test@example.com",
				tt.initialPassword,
				"+5511999999999",
				"avatar.jpg",
			)
			assert.NoError(t, err)

			result := u.ComparePassword(tt.comparePassword)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestUpdate(t *testing.T) {
	initialUser, err := New(
		"Initial Name",
		"test@example.com",
		"StrongPass123!",
		"+5511999999999",
		"initial-avatar.jpg",
	)
	assert.NoError(t, err)

	originalTime := initialUser.UpdatedAt

	time.Sleep(10 * time.Millisecond)

	tests := []struct {
		name        string
		newName     string
		newAvatar   string
		newPhone    string
		expectedErr bool
		errMessage  string
	}{
		{
			name:        "valid update",
			newName:     "New Name",
			newAvatar:   "new-avatar.jpg",
			newPhone:    "+5511988888888",
			expectedErr: false,
		},
		{
			name:        "empty name",
			newName:     "",
			newAvatar:   "new-avatar.jpg",
			newPhone:    "+5511988888888",
			expectedErr: true,
			errMessage:  "name is required",
		},
		{
			name:        "name too short",
			newName:     "A",
			newAvatar:   "new-avatar.jpg",
			newPhone:    "+5511988888888",
			expectedErr: true,
			errMessage:  "minimum name length",
		},
		{
			name:        "empty avatar",
			newName:     "New Name",
			newAvatar:   "",
			newPhone:    "+5511988888888",
			expectedErr: true,
			errMessage:  "avatar is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := New(
				initialUser.Name,
				initialUser.Email.String(),
				"StrongPass123!",
				initialUser.Phone.String(),
				initialUser.Avatar,
			)
			assert.NoError(t, err)

			phoneVO, _ := phone.New(tt.newPhone)

			err = u.Update(tt.newName, tt.newAvatar, phoneVO)

			if tt.expectedErr {
				assert.Error(t, err)
				if tt.errMessage != "" {
					assert.Contains(t, err.Error(), tt.errMessage)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.newName, u.Name)
				assert.Equal(t, tt.newAvatar, u.Avatar)
				assert.Equal(t, tt.newPhone, u.Phone.String())
				assert.True(t, u.UpdatedAt.After(originalTime))
			}
		})
	}
}

func TestChangePassword(t *testing.T) {
	u, err := New(
		"Test User",
		"test@example.com",
		"OldPassword123!",
		"+5511999999999",
		"avatar.jpg",
	)
	assert.NoError(t, err)

	originalTime := u.UpdatedAt

	time.Sleep(10 * time.Millisecond)

	tests := []struct {
		name        string
		newPassword string
		expectedErr bool
		shouldMatch bool
	}{
		{
			name:        "valid password change",
			newPassword: "NewStrongPass456!",
			expectedErr: false,
			shouldMatch: true,
		},
		{
			name:        "weak password",
			newPassword: "123",
			expectedErr: true,
			shouldMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testUser, err := New(
				"Test User",
				"test@example.com",
				"OldPassword123!",
				"+5511999999999",
				"avatar.jpg",
			)
			assert.NoError(t, err)

			err = testUser.ChangePassword(tt.newPassword)

			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.True(t, testUser.ComparePassword(tt.newPassword))
				assert.False(t, testUser.ComparePassword("OldPassword123!"))
				assert.True(t, testUser.UpdatedAt.After(originalTime))
			}
		})
	}
}

func TestChangeEmail(t *testing.T) {
	u, err := New(
		"Test User",
		"old@example.com",
		"StrongPass123!",
		"+5511999999999",
		"avatar.jpg",
	)
	assert.NoError(t, err)

	originalTime := u.UpdatedAt

	time.Sleep(10 * time.Millisecond)

	tests := []struct {
		name        string
		newEmail    string
		expectedErr bool
	}{
		{
			name:        "valid email change",
			newEmail:    "new@example.com",
			expectedErr: false,
		},
		{
			name:        "invalid email",
			newEmail:    "invalid-email",
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testUser, err := New(
				"Test User",
				"old@example.com",
				"StrongPass123!",
				"+5511999999999",
				"avatar.jpg",
			)
			assert.NoError(t, err)

			err = testUser.ChangeEmail(tt.newEmail)

			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.newEmail, testUser.Email.String())
				assert.True(t, testUser.UpdatedAt.After(originalTime))
			}
		})
	}
}
