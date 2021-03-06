package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/soundsnick/arzamas/comment"
	"github.com/soundsnick/arzamas/core"
	"github.com/soundsnick/arzamas/handlers"
	"github.com/soundsnick/arzamas/post"
	"github.com/soundsnick/arzamas/session"
	"github.com/soundsnick/arzamas/user"
)

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	// Load configurations
	initLogger()
	core.LoadConfig()
	core.SetDB(core.GetConnectionString())
	AutoMigrate()

	// Application routes
	router.GET("/", handlers.IndexPage)

	// Post routes
	router.GET("/posts", handlers.PostIndex)
	router.GET("/posts/last", handlers.PostLast)
	router.GET("/posts/search", handlers.PostSearch)
	router.GET("/posts/user/:user_id", handlers.PostUser)
	router.GET("/post/comments/:id", handlers.PostComments)

	// Post CRUD
	router.POST("/post/create", handlers.PostCreate)
	router.GET("/post/read/:id", handlers.PostRead)
	router.PUT("/post/update/:id", handlers.PostUpdate)
	router.DELETE("/post/delete/:id", handlers.PostDelete)

	// Comment CRUD
	router.POST("/comment/create/:id", handlers.CommentCreate)
	router.GET("/comment/read/:id", handlers.CommentRead)
	router.PUT("/comment/update/:id", handlers.CommentUpdate)
	router.DELETE("/comment/delete/:id", handlers.CommentDelete)

	// User routes
	router.GET("/users/email/:email", handlers.UserByEmail)
	router.GET("/users/name/:name", handlers.UsersByName)
	router.GET("/users/last_name/:name", handlers.UsersByLastName)

	// User CRUD(RUD)
	router.GET("/user/read/:id", handlers.UserRead)
	router.GET("/user/readByToken", handlers.UserReadByToken)
	router.PUT("/user/update", handlers.UserUpdate)
	router.DELETE("/user/delete", handlers.UserDelete)

	// User authentication
	router.POST("/user/auth", handlers.UserAuthenticate)
	router.POST("/user/register", handlers.UserRegister)
	router.DELETE("/user/logout", handlers.UserLogout)

	// Docs
	router.GET("/docs", handlers.DocsIndex)
	router.GET("/docs/posts", handlers.DocsPosts)
	router.GET("/docs/users", handlers.DocsUsers)
	router.GET("/docs/comments", handlers.DocsComments)

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

// AutoMigrate migrates models
func AutoMigrate() {
	core.GetDB().AutoMigrate(&user.User{}, &post.Post{}, &session.Session{}, &comment.Comment{})
}

// CORSMiddleware makes cors real
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
