package post

import (
	"strings"

	"github.com/soundsnick/arzamas/core"
)

// GetAll returns a collection of all posts
func GetAll() []Post {
	var posts []Post
	core.GetDB().Preload("User").Order("created_at desc").Find(&posts)
	return posts
}

// GetLast returns a collection of all posts
func GetLast() []Post {
	var posts []Post
	core.GetDB().Preload("User").Order("created_at desc").Limit(9).Find(&posts)
	return posts
}

// GetByID return post by id
func GetByID(id uint64) Post {
	var post Post
	core.GetDB().Preload("User").Where("id = ?", id).Find(&post)
	return post
}

// GetByTitle returns posts by title
func GetByTitle(payload string) []Post {
	var posts []Post
	core.GetDB().Preload("User").Where("LOWER(title) LIKE ?", strings.ToLower(payload)+"%").Order("created_at desc").Find(&posts)
	return posts
}

// GetByUserID return posts of user
func GetByUserID(userID uint64) []Post {
	var posts []Post
	core.GetDB().Preload("User").Where("user_id = ?", userID).Order("created_at desc").Find(&posts)
	return posts
}

// DeleteByID deletes post
func DeleteByID(id uint64) Post {
	var post Post
	core.GetDB().Where("id = ?", id).Delete(&post)
	return post
}
