package shared

import (
	"github.com/soundsnick/arzamas/core"
	"github.com/soundsnick/arzamas/post"
	"github.com/soundsnick/arzamas/session"
)

// DeleteSessionsByUserID ends all session of the user
func DeleteSessionsByUserID(ID uint64) {
	var sessions []session.Session
	core.GetDB().Where("user_id = ?", ID).Delete(&sessions)
}

// DeletePostsByUserID ends all session of the user
func DeletePostsByUserID(ID uint64) {
	var posts []post.Post
	core.GetDB().Where("user_id = ?", ID).Delete(&posts)
}
