package main

import (
	"github.com/gin-gonic/gin"
	"github.com/soundsnick/arzamas/config"
	"github.com/soundsnick/arzamas/controllers"
	"github.com/soundsnick/arzamas/models"
)

func main() {
	router := gin.Default()

	// Load configurations
	config.LoadConfig()
	models.SetDB(config.GetConnectionString())
	models.AutoMigrate()
	config.SeedTestPosts()

	router.GET("/", controllers.IndexPage)
	router.GET("/posts", controllers.PostIndex)
	router.GET("/users/email/:email", controllers.UserByEmail)
	router.GET("/users/name/:name", controllers.UsersByName)
	router.GET("/users/last_name/:name", controllers.UsersByLastName)

	// Run listener
	router.Run()
}
