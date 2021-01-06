package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/soundsnick/arzamas/models"
)

// UserByEmail get user by email
func UserByEmail(c *gin.Context) {
	email := c.Param("email")
	user := models.GetUserByEmail(email)
	c.JSON(200, gin.H{
		"data": user,
	})
}

// UsersByName get users by name
func UsersByName(c *gin.Context) {
	name := c.Param("name")
	users := models.GetUsersByName(name)
	c.JSON(200, gin.H{
		"data": users,
	})
}

// UsersByLastName get users by name
func UsersByLastName(c *gin.Context) {
	lastName := c.Param("name")
	users := models.GetUsersByLastName(lastName)
	c.JSON(200, gin.H{
		"data": users,
	})
}
