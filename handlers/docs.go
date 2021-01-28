package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/soundsnick/arzamas/docs"
)

func DocsIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"posts": "/docs/posts",
		"users": "/docs/users",
		"comment": "/docs/comments",
	})
}

func DocsPosts(c *gin.Context) {
	c.JSON(200, gin.H{
		"index":    docs.PostIndex,
		"search":   docs.PostSearch,
		"user":     docs.PostUser,
		"comments": docs.PostComments,
		"create":   docs.PostCreate,
		"get":      docs.PostGet,
		"update":   docs.PostUpdate,
		"delete":   docs.PostDelete,
	})
}

func DocsUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"authenticate":     docs.UserAuthenticate,
		"register":         docs.UserRegister,
		"read":             docs.UserRead,
		"update":           docs.UserUpdate,
		"delete":           docs.UserDelete,
		"get by email":     docs.UserByEmail,
		"get by name":      docs.UsersByName,
		"get by last name": docs.UsersByLastName,
	})
}

func DocsComments(c *gin.Context) {
	c.JSON(200, gin.H{
		"create": docs.CommentCreate,
		"read": docs.CommentRead,
		"update": docs.CommentUpdate,
		"delete": docs.CommentDelete,
	})
}