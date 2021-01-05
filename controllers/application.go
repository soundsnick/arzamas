package controllers

import (
	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
    c.JSON(200, gin.H{
		"title": "Arzamas index page",
	})
}