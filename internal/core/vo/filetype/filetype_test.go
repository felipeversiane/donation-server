package filetype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		input   string
		isValid bool
	}{
		{"application/pdf", false},
		{"image/png", true},
		{"image/jpeg", true},
		{"image/jpg", true},
		{"application/zip", false},
		{"text/plain", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			ft, err := New(tt.input)
			if tt.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tt.input, ft.String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
