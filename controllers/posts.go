package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/soundsnick/arzamas/models"
)

// PostIndex returns all posts
func PostIndex(c *gin.Context) {
	posts := models.GetPostsAll()
	c.JSON(200, gin.H{
		"data": posts,
	})
}
