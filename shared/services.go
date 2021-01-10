package shared

import (
	"github.com/soundsnick/arzamas/comment"
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
	core.GetDB().Where("user_id = ?", ID).Find(&posts)
	// TODO: Refactor
	for _, s := range posts {
		DeleteCommentsByPostID(s.ID)
	}
	core.GetDB().Delete(&posts)
}

// DeleteCommentsByPostID deletes all comments by post
func DeleteCommentsByPostID(ID uint64) {
	var comments []comment.Comment
	core.GetDB().Debug().Where("post_id = ?", ID).Delete(&comments)
}

// DeleteCommentsByUserID deletes all comments by user
func DeleteCommentsByUserID(ID uint64) {
	var comments []comment.Comment
	core.GetDB().Where("user_id = ?", ID).Delete(&comments)
}
