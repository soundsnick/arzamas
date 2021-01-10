package post

import (
	"github.com/soundsnick/arzamas/core"
)

// CreationForm type for Post create input
type CreationForm struct {
	Title   string
	Content string
	Cover   string
	Token   string
}

// ValidateCreationForm validates Post input
func ValidateCreationForm(form CreationForm) (string, error) {
	if form.Title == "" {
		return "title", core.ErrValidation
	}
	if !IsContentValid(form.Content) {
		return "content", core.ErrValidation
	}
	if !core.IsURL(form.Cover) {
		return "cover", core.ErrValidation
	}
	if form.Token == "" {
		return "token", core.ErrValidation
	}
	return "ok", nil
}

// ValidateUpdateForm validates Post input
func ValidateUpdateForm(form *CreationForm, postCurrent Post) (string, error) {
	if form.Title == "" {
		form.Title = postCurrent.Title
	}
	if form.Content != "" {
		if !IsContentValid(form.Content) {
			return "content", core.ErrValidation
		}
	} else {
		form.Content = postCurrent.Content
	}
	if form.Cover != "" {
		if !core.IsURL(form.Cover) {
			return "cover", core.ErrValidation
		}
	} else {
		form.Cover = postCurrent.Cover
	}
	if form.Token == "" {
		return "token", core.ErrValidation
	}
	return "ok", nil
}

// IsContentValid validates content
func IsContentValid(t string) bool {
	if len(t) < 30 {
		return false
	}
	return true
}
