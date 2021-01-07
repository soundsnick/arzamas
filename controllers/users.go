package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/soundsnick/arzamas/helpers"
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
			// Authorise user and return token
			session := models.Session{Token: models.GenerateToken(), UserID: user.ID, IP: c.ClientIP()}
			sessionRes := models.GetDB().Create(&session)
			if sessionRes.Error != nil {
				c.JSON(422, gin.H{
					"error": sessionRes.Error,
				})
			} else {
				c.JSON(200, gin.H{
					"user":  user.FilterPassword(),
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

// UserRegister registers user
func UserRegister(c *gin.Context) {
	form := helpers.UserRegistrationForm{
		Email:                c.Query("email"),
		Name:                 c.Query("name"),
		LastName:             c.Query("last_name"),
		Password:             c.Query("password"),
		PasswordConfirmation: c.Query("password_confirmation"),
	}

	// Check validation
	validatedField, validateErr := helpers.ValidateUserRegistration(form)

	if validateErr == nil {

		// Check if user exists
		user := models.GetUserByEmail(form.Email)
		if user.ID == 0 {
			userNew := models.User{Name: form.Name, LastName: form.LastName, Email: form.Email, Password: form.Password}
			models.GetDB().Create(&userNew)

			// Authorise user and return token
			session := models.Session{Token: models.GenerateToken(), UserID: userNew.ID, IP: c.ClientIP()}
			sessionRes := models.GetDB().Create(&session)
			if sessionRes.Error != nil {
				c.JSON(422, gin.H{
					"error": sessionRes.Error,
				})
			} else {
				c.JSON(200, gin.H{
					"user":  userNew,
					"token": session.Token,
				})
			}
		} else {
			c.JSON(422, gin.H{
				"error": "already exists",
			})
		}
	} else {
		c.JSON(400, gin.H{
			"error": validateErr,
			"field": validatedField,
		})
	}
}
