package user

import (
	"net"
	"regexp"
	"strings"

	"github.com/soundsnick/arzamas/core"
	"golang.org/x/crypto/bcrypt"
)

// RegistrationForm type for UserRegistration input
type RegistrationForm struct {
	Email                string
	Name                 string
	LastName             string
	Password             string
	PasswordConfirmation string
}

// UpdateForm type for UserRegistration input
type UpdateForm struct {
	Name     string
	LastName string
	Avatar   string
	Token    string
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// ValidateRegistrationForm validates UserRegistration input
func ValidateRegistrationForm(form RegistrationForm) (string, error) {
	if !IsEmailValid(form.Email) {
		return "email", core.ErrValidation
	}
	if !IsNameValid(form.Name) {
		return "name", core.ErrValidation
	}
	if !IsNameValid(form.LastName) {
		return "last_name", core.ErrValidation
	}
	if !IsPasswordValid(form.Password) {
		return "password", core.ErrValidation
	}
	if !IsPasswordConfirmationValid(form.Password, form.PasswordConfirmation) {
		return "password_confirmation", core.ErrValidation
	}
	return "ok", nil
}

// ValidateUpdateForm validates User input
func ValidateUpdateForm(form *UpdateForm, userCurrent User) (string, error) {
	if form.Name != "" {
		if !IsNameValid(form.Name) {
			return "name", core.ErrValidation
		}
	} else {
		form.Name = userCurrent.Name
	}
	if form.LastName != "" {
		if !IsNameValid(form.LastName) {
			return "last_name", core.ErrValidation
		}
	} else {
		form.LastName = userCurrent.LastName
	}
	if form.Avatar != "" {
		if !core.IsURL(form.Avatar) {
			return "avatar", core.ErrValidation
		}
	} else {
		form.Avatar = userCurrent.Avatar
	}
	if form.Token == "" {
		return "token", core.ErrValidation
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

// EncryptPassword generates hash from password
func EncryptPassword(password string) (string, error) {
	var hash []byte
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
