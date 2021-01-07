package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/soundsnick/arzamas/models"
	"golang.org/x/crypto/bcrypt"
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

// UserAuthenticate authenticates user by email and password
func UserAuthenticate(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")

	if len(email) > 0 && len(password) > 0 {
		// Get user by email
		user := models.GetUserByEmail(email)

		// If user not found by email OR passwords doesn't match
		if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
			c.JSON(422, gin.H{
				"error": "wrong email or password",
			})
		} else {
			models.DeleteOrphanSession(c.ClientIP())
			session := models.Session{Token: models.GenerateToken(), UserID: user.ID, IP: c.ClientIP()}
			sessionRes := models.GetDB().Create(&session)
			if sessionRes.Error != nil {
				c.JSON(422, gin.H{
					"error": sessionRes.Error,
				})
			} else {
				c.JSON(200, gin.H{
					"user":  user,
					"token": session.Token,
				})
			}
		}
	} else {
		c.JSON(400, gin.H{
			"error": "email and password required",
		})
	}
}
