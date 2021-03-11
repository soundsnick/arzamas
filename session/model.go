package session

import (
	"github.com/soundsnick/arzamas/core"
	"github.com/soundsnick/arzamas/user"
)

// Session model for session
type Session struct {
	core.Model

	Token  string    `form:"token" binding:"required" json:"token"`
	UserID uint64    `json:"user_id"`
	IP     string    `json:"-"`
	User   user.User `binding:"-" gorm:"association_autoupdate:false;association_autocreate:false"`
}
