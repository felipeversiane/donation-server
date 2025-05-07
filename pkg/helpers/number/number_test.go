package number

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeIntToInt32(t *testing.T) {
	tests := []struct {
		name      string
		input     int
		expected  int32
		expectErr bool
	}{
		{
			name:      "within range",
			input:     123,
			expected:  123,
			expectErr: false,
		},
		{
			name:      "max int32",
			input:     math.MaxInt32,
			expected:  math.MaxInt32,
			expectErr: false,
		},
		{
			name:      "min int32",
			input:     math.MinInt32,
			expected:  math.MinInt32,
			expectErr: false,
		},
		{
			name:      "above int32 range",
			input:     math.MaxInt32 + 1,
			expectErr: true,
		},
		{
			name:      "below int32 range",
			input:     math.MinInt32 - 1,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SafeIntToInt32(tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
