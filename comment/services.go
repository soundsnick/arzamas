package comment

import (
	"github.com/soundsnick/arzamas/core"
)

// DeleteByID deletes comment by id
func DeleteByID(ID uint64) {
	var comment Comment
	core.GetDB().Where("id = ?", ID).Delete(&comment)
}
