package documenttype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		input   string
		isValid bool
	}{
		{"cpf", true},
		{"CNPJ", true},
		{"rg", false},
		{"cnh", false},
		{"", false},
		{"passaporte", false},
		{" cpf ", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			dt, err := New(tt.input)
			if tt.isValid {
				assert.NoError(t, err)
				assert.NotEmpty(t, dt.String())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
