package post

import (
	"github.com/soundsnick/arzamas/core"
	"github.com/soundsnick/arzamas/user"
)

// Post model
type Post struct {
	core.Model

	Title   string `form:"title" binding:"required"`
	Content string `form:"content"`
	Cover   string `form:"cover"`
	UserID  uint64
	User    user.User `binding:"-" gorm:"association_autoupdate:false;association_autocreate:false"`
}
