package helpers

import (
	"errors"
	"net"
	"regexp"
	"strings"
)

// UserRegistrationForm is a type for UserRegistration input
type UserRegistrationForm struct {
	Email                string
	Name                 string
	LastName             string
	Password             string
	PasswordConfirmation string
}

// ErrValidation for validation erroring
var ErrValidation = errors.New("validation error")
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// ValidateUserRegistration validates UserRegistration input
func ValidateUserRegistration(form UserRegistrationForm) (string, error) {
	if !IsEmailValid(form.Email) {
		return "email", ErrValidation
	}
	if !IsNameValid(form.Name) {
		return "name", ErrValidation
	}
	if !IsNameValid(form.LastName) {
		return "last_name", ErrValidation
	}
	if !IsPasswordValid(form.Password) {
		return "password", ErrValidation
	}
	if !IsPasswordConfirmationValid(form.Password, form.PasswordConfirmation) {
		return "password_confirmation", ErrValidation
	}
	return "ok", nil
}

// IsEmailValid validates email field
func IsEmailValid(e string) bool {
	if len(e) < 3 || len(e) > 254 {
		return false
	}
	if !emailRegex.MatchString(e) {
		return false
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false
	}
	return true
}

// IsNameValid validates name and last name field
func IsNameValid(n string) bool {
	if len(n) < 2 || len(n) > 30 {
		return false
	}
	if len(strings.Split(n, " ")) != 1 {
		return false
	}
	return true
}

// IsPasswordValid validates password
func IsPasswordValid(n string) bool {
	if len(n) < 6 {
		return false
	}
	return true
}

// IsPasswordConfirmationValid validates password confirmation
func IsPasswordConfirmationValid(password string, passwordConfirmation string) bool {
	if !IsPasswordValid(passwordConfirmation) || password != passwordConfirmation {
		return false
	}
	return true
}
