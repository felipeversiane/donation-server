package field

import (
	"testing"
)

func TestValidateRequired(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		fieldName string
		wantErr   bool
	}{
		{"Empty string", "", "Name", true},
		{"Whitespace only", "   ", "Email", true},
		{"Valid input", "John", "Name", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateRequired(tt.value, tt.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateRequired() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateMaxLength(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		max       int
		fieldName string
		wantErr   bool
	}{
		{"Too long", "abcdef", 5, "Username", true},
		{"Exact length", "abcde", 5, "Username", false},
		{"Short", "abc", 5, "Username", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMaxLength(tt.value, tt.max, tt.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateMaxLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateMinLength(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		min       int
		fieldName string
		wantErr   bool
	}{
		{"Too short", "abc", 5, "Password", true},
		{"Exact length", "abcde", 5, "Password", false},
		{"Longer", "abcdef", 5, "Password", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMinLength(tt.value, tt.min, tt.fieldName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateMinLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
