package controllers

import (
	"github.com/gin-gonic/gin"
)

// IndexPage root page
func IndexPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"title": "Arzamas index page",
	})
}
