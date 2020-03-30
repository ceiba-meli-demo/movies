package validators

import (
	"errors"
	"strings"
)

func ValidateRequired(field string, message string) error {
	if strings.TrimSpace(field) == "" {
		return errors.New(message)
	}
	return nil
}
func ValidateRequiredDuration(field int64, message string) error {
	if field == 0 || field <0 {
		return errors.New(message)
	}
	return nil
}
