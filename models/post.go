package models

import (
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

// GetPostsAll returns a collection of all posts
func GetPostsAll() []Post {
	var posts []Post
	GetDB().Preload("User").Order("created_at desc").Find(&posts)
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
	GetDB().Preload("User").Where("LOWER(title) LIKE ?", strings.ToLower(payload)+"%").Order("created_at desc").Find(&posts)
	return posts
}

// GetPostsByUserID return posts of user
func GetPostsByUserID(userID uint64) []Post {
	var posts []Post
	GetDB().Preload("User").Where("user_id = ?", userID).Order("created_at desc").Find(&posts)
	return posts
}
