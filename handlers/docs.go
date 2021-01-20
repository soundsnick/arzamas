package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/soundsnick/arzamas/docs"
)

func DocsIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"posts": docs.PostIndex,
	})
}
