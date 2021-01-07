package models

import (
	"fmt"
	"html/template"
	"strings"
)

// Post model
type Post struct {
	Model

	Title   string `form:"title" binding:"required"`
	Content string `form:"content"`
	Cover   string `form:"cover"`
	UserID  uint64
	User    User `binding:"-" gorm:"association_autoupdate:false;association_autocreate:false"`
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
func GetPostsAll() []Post {
	var posts []Post
	GetDB().Preload("User").Find(&posts)
	return posts
}

// GetPostByID return post by id
func GetPostByID(id uint64) Post {
	var post Post
	GetDB().Preload("User").Where("id = ?", id).Find(&post)
	return post
}

// GetPostsByTitle returns posts by title
func GetPostsByTitle(payload string) []Post {
	var posts []Post
	GetDB().Preload("User").Where("LOWER(title) LIKE ?", strings.ToLower(payload)+"%").Find(&posts)
	return posts
}

// GetPostsByUserID return posts of user
func GetPostsByUserID(userID uint64) []Post {
	var posts []Post
	GetDB().Preload("User").Where("user_id = ?", userID).Find(&posts)
	return posts
}
