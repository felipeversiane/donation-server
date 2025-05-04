package str

import "testing"

func TestCleanString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Only digits", "123456", "123456"},
		{"With letters", "abc123xyz", "123"},
		{"With symbols", "a1!b2@c3#", "123"},
		{"Empty string", "", ""},
		{"Only non-digits", "abc!@#", ""},
		{"Mixed unicode", "1٢٣四5", "15"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CleanString(tt.input)
			if result != tt.expected {
				t.Errorf("CleanString(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
