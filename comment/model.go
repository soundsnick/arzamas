package comment

import (
	"github.com/soundsnick/arzamas/core"
	"github.com/soundsnick/arzamas/post"
	"github.com/soundsnick/arzamas/user"
)

// Comment model for comment
type Comment struct {
	core.Model

	Content string `form:"content" binding:"required"`
	PostID  uint64
	Post    post.Post `binding:"-" gorm:"association_autoupdate:false;association_autocreate:false"`
	UserID  uint64
	User    user.User `binding:"-" gorm:"association_autoupdate:false;association_autocreate:false"`
}
