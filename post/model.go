package post

import (
	"github.com/soundsnick/arzamas/core"
	"github.com/soundsnick/arzamas/user"
)

// Post model
type Post struct {
	core.Model

	Title   string    `form:"title" binding:"required" json:"title"`
	Content string    `form:"content" json:"content"`
	Description string `form:"description" json:"description"`
	Cover   string    `form:"cover" json:"cover"`
	UserID  uint64    `form:"user_id" json:"user_id"`
	User    user.User `binding:"-" gorm:"association_autoupdate:false;association_autocreate:false"`
}
