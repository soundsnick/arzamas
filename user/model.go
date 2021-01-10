package user

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/soundsnick/arzamas/core"
	"gorm.io/gorm"
)

// User model for user
type User struct {
	core.Model

	Email    string `form:"email" binding:"required" json:"-"`
	Name     string `form:"name"`
	LastName string `form:"last_name"`
	Password string `form:"password" binding:"required" json:"-"`
}

var user User

// BeforeSave - hook being executed before each save
func (user *User) BeforeSave(*gorm.DB) (err error) {
	hashedPassword, err := EncryptPassword(user.Password)
	if err != nil {
		return
	}
	user.Password = hashedPassword
	user.Name = strings.Title(user.Name)
	user.LastName = strings.Title(user.LastName)
	logrus.WithFields(logrus.Fields{
		"event": "NEWUSER",
	}).Infof("%s is registered", user.Email)
	return
}
