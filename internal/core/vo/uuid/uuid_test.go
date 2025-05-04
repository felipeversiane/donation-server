package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	u, err := New()
	assert.NoError(t, err)
	assert.NotEmpty(t, u.String())
}

func TestFromString(t *testing.T) {
	validUUID := "550e8400-e29b-41d4-a716-446655440000"
	invalidUUID := "invalid-uuid"

	t.Run("valid UUID", func(t *testing.T) {
		u, err := FromString(validUUID)
		assert.NoError(t, err)
		assert.Equal(t, validUUID, u.String())
	})

	t.Run("invalid UUID", func(t *testing.T) {
		_, err := FromString(invalidUUID)
		assert.Error(t, err)
	})
}
