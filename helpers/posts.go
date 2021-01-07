package helpers

import "net/url"

// PostForm type for Post input
type PostForm struct {
	Title   string
	Content string
	Cover   string
	Token   string
}

// ValidatePostForm validates Post input
func ValidatePostForm(form PostForm) (string, error) {
	if form.Title == "" {
		return "title", ErrValidation
	}
	if !IsContentValid(form.Content) {
		return "content", ErrValidation
	}
	if !IsCoverValid(form.Cover) {
		return "cover", ErrValidation
	}
	if form.Token == "" {
		return "token", ErrValidation
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

// IsCoverValid validates cover
func IsCoverValid(c string) bool {
	_, err := url.Parse(c)

	if err != nil {
		return false
	}
	if c == "" {
		return false
	}
	return true
}
