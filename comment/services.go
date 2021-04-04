package comment

import (
	"github.com/soundsnick/arzamas/core"
)

// GetByID reads comment by id
func GetByID(ID uint64) Comment {
	var comment Comment
	core.GetDB().Preload("User").Preload("Post").Where("id = ?", ID).Find(&comment)
	return comment
}

// GetByPostID return comments by post
func GetByPostID(ID uint64) []Comment {
	var comments []Comment
	core.GetDB().Preload("User").Preload("Post").Where("post_id = ?", ID).Order("created_at desc").Find(&comments)
	return comments
}

// DeleteByID deletes comment by id
func DeleteByID(ID uint64) {
	var comment Comment
	core.GetDB().Where("id = ?", ID).Delete(&comment)
}
