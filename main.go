package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/soundsnick/arzamas/config"
	"github.com/soundsnick/arzamas/controllers"
	"github.com/soundsnick/arzamas/models"
)

func main() {
	router := gin.Default()

	// Load configurations
	initLogger()
	config.LoadConfig()
	models.SetDB(config.GetConnectionString())
	models.AutoMigrate()
	config.SeedTestPosts()

	// Application routes
	router.GET("/", controllers.IndexPage)

	// Post routes
	router.GET("/posts", controllers.PostIndex)

	// User routes
	router.GET("/users/email/:email", controllers.UserByEmail)
	router.GET("/users/name/:name", controllers.UsersByName)
	router.GET("/users/last_name/:name", controllers.UsersByLastName)
	router.GET("/user/auth", controllers.UserAuthenticate)

	// Run listener
	router.Run()
}

//initLogger initializes logrus logger with some defaults
func initLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	//logrus.SetOutput(os.Stderr)
	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
