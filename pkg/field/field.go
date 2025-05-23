package field

import (
	"fmt"
	"strings"
)

func ValidateRequired(value, fieldName string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

func ValidateMaxLength(value string, max int, fieldName string) error {
	if len(value) > max {
		return fmt.Errorf("%s must be at most %d characters", fieldName, max)
	}
	return nil
}

func ValidateMinLength(value string, min int, fieldName string) error {
	if len(value) < min {
		return fmt.Errorf("%s must be at least %d characters", fieldName, min)
	}
	return nil
}
