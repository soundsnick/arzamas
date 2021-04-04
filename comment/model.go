package comment

import (
	"github.com/soundsnick/arzamas/core"
	"github.com/soundsnick/arzamas/post"
	"github.com/soundsnick/arzamas/user"
)

// Comment model for comment
type Comment struct {
	core.Model

	Content string `form:"content" binding:"required" json:"content"`
	PostID  uint64 `json:"post_id"`
	Post    post.Post `binding:"-" gorm:"association_autoupdate:false;association_autocreate:false" json:"post"`
	UserID  uint64 `json:"user_id"`
	User    user.User `binding:"-" gorm:"association_autoupdate:false;association_autocreate:false" json:"user"`
}
