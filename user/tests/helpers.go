package tests

import (
	"testing"

	helpers "github.com/soundsnick/arzamas/user"
)

// TestIsEmailValid is a test case that checks email
func TestIsEmailValid(t *testing.T) {
	validMx := "example@gmail.com"
	invalidMx := "example@golangcode-example.com"
	invalidLength := "ex"

	testCaseMxMustValid := helpers.IsEmailValid(validMx)
	testCaseMxMustInvalid := helpers.IsEmailValid(invalidMx)
	testCaseLengthMustInvalid := helpers.IsEmailValid(invalidLength)

	if !testCaseMxMustValid {
		t.Errorf("%s must be valid", validMx)
	}
	if testCaseMxMustInvalid {
		t.Errorf("%s must be invalid (MX)", invalidMx)
	}
	if testCaseLengthMustInvalid {
		t.Errorf("%s must be invalid (Length)", invalidLength)
	}
}

// TestIsNameValid is a test case that checks name
func TestIsNameValid(t *testing.T) {
	valid := "yernazar"
	invalidLength := "y"
	invalidWords := "Qasym Zhomart"

	if !helpers.IsNameValid(valid) {
		t.Errorf("%s must be valid", valid)
	}
	if helpers.IsNameValid(invalidLength) {
		t.Errorf("%s must be invalid (Length)", invalidLength)
	}
	if helpers.IsNameValid(invalidWords) {
		t.Errorf("%s must be invalid (Length of words)", invalidWords)
	}
}

// TestIsPasswordConfirmationValid is a test cast that checks password
func TestIsPasswordConfirmationValid(t *testing.T) {
	password := "qwerty"
	passwordInvalid := "qwe"
	passwordConfirmation := "qwerty"
	passwordConfirmationInvalid := "ytrewq"

	if !helpers.IsPasswordValid(password) {
		t.Errorf("%s must be valid", password)
	}
	if helpers.IsPasswordValid(passwordInvalid) {
		t.Errorf("%s must be invalid", passwordInvalid)
	}
	if !helpers.IsPasswordConfirmationValid(password, passwordConfirmation) {
		t.Errorf("%s, %s must be valid", password, passwordConfirmation)
	}
	if helpers.IsPasswordConfirmationValid(password, passwordConfirmationInvalid) {
		t.Errorf("%s, %s must be invalid", password, passwordConfirmationInvalid)
	}
}
