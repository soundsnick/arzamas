package core

import (
	"errors"
	"net/url"
)

// ErrValidation for validation erroring
var ErrValidation = errors.New("validation error")

// IsURL validates cover
func IsURL(c string) bool {
	_, err := url.Parse(c)

	if err != nil {
		return false
	}
	if c == "" {
		return false
	}
	return true
}
