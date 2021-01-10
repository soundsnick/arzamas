package handlers

import (
	"strconv"

	"github.com/soundsnick/arzamas/comment"
	"github.com/soundsnick/arzamas/core"

	"github.com/gin-gonic/gin"
	"github.com/soundsnick/arzamas/post"
	"github.com/soundsnick/arzamas/session"
)

// CommentCreate create operation
func CommentCreate(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	token := c.Query("token")
	content := c.Query("content")
	userFound := session.GetUserByToken(token)
	postFound := post.GetByID(ID)
	if postFound.ID == 0 || err != nil || userFound.ID == 0 {
		c.JSON(422, gin.H{
			"error": "not found",
		})
	} else {
		if content != "" && len(content) > 2 {
			commentNew := comment.Comment{Content: content, UserID: userFound.ID, PostID: postFound.ID}
			core.GetDB().Create(&commentNew)
			c.JSON(200, gin.H{
				"data": commentNew,
			})
		} else {
			c.JSON(422, gin.H{
				"error": "wrong content",
			})
		}
	}
}

// CommentRead read operation
func CommentRead(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	commentFound := comment.GetByID(ID)
	if commentFound.ID == 0 || err != nil {
		c.JSON(422, gin.H{
			"error": "not found",
		})
	} else {
		c.JSON(200, gin.H{
			"data": commentFound,
		})
	}
}

// CommentUpdate update operation
func CommentUpdate(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	token := c.Query("token")
	content := c.Query("content")
	userFound := session.GetUserByToken(token)
	commentFound := comment.GetByID(ID)
	if commentFound.ID == 0 || err != nil || userFound.ID == 0 {
		c.JSON(422, gin.H{
			"error": "not found",
		})
	} else {
		if commentFound.UserID != userFound.ID {
			c.JSON(422, gin.H{
				"error": "wrong token",
			})
		} else {
			if content != "" && len(content) > 2 {
				commentFound.Content = content
				core.GetDB().Save(&commentFound)
				c.JSON(200, gin.H{
					"data": commentFound,
				})
			} else {
				c.JSON(422, gin.H{
					"error": "wrong content",
				})
			}
		}
	}
}

// CommentDelete delete operation
func CommentDelete(c *gin.Context) {
	ID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	token := c.Query("token")
	userFound := session.GetUserByToken(token)
	commentFound := comment.GetByID(ID)
	if commentFound.ID == 0 || err != nil || userFound.ID == 0 {
		c.JSON(422, gin.H{
			"error": "not found",
		})
	} else {
		if commentFound.UserID != userFound.ID {
			c.JSON(422, gin.H{
				"error": "wrong token",
			})
		} else {
			core.GetDB().Delete(&commentFound)
			c.JSON(200, gin.H{
				"message": "deleted",
			})
		}
	}
}
