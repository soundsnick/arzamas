package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soundsnick/arzamas/controllers"
	"github.com/soundsnick/arzamas/config"
	"github.com/soundsnick/arzamas/models"
)

func main() {
	router := gin.Default()

	config.LoadConfig()
	models.SetDB(config.GetConnectionString())
	models.AutoMigrate()

	router.GET("/", controllers.IndexPage)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // 0.0.0.0:8080
}