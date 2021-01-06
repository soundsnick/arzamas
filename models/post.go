package models

import (
	"fmt"
	"html/template"
)

// Post model
type Post struct {
	Model

	Title     string `form:"title" binding:"required"`
	Content   string `form:"content"`
	Published bool   `form:"published"`
	Cover     string `form:"cover"`
	UserID    uint64
	User      User `binding:"-" gorm:"association_autoupdate:false;association_autocreate:false"`
}

// HTMLContent returns html content that won't be escaped
func (post *Post) HTMLContent() template.HTML {
	return template.HTML(post.Content)
}

// URL returns the post's canonical url
func (post *Post) URL() string {
	return fmt.Sprintf("/posts/%d", post.ID)
}

// GetPostsAll returns a collection of all posts
func GetPostsAll() interface{} {
	var posts []Post
	return GetDB().Preload("User").Find(&posts).Value
}