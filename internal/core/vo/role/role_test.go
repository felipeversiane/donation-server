package role

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		input   string
		isValid bool
	}{
		{"admin", true},
		{"user", true},
		{"ADMIN", true},
		{"User", true},
		{"manager", false},
		{"", false},
		{"superuser", false},
		{"admin ", false}, 
		{" user", false},  
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r, err := New(tt.input)
			if tt.isValid {
				assert.NoError(t, err)
				assert.NotEmpty(t, r.String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
