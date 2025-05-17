package usertype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		input   string
		isValid bool
	}{
		{"individual", true},
		{"company", true},
		{"Individual", true},
		{"Company", true},
		{"pessoa física", false},
		{"", false},
		{" admin", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			typ, err := New(tt.input)
			if tt.isValid {
				assert.NoError(t, err)
				assert.NotEmpty(t, typ.String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
