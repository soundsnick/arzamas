package models

import (
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	Model

	Email    string `form:"email" binding:"required" json:"-"`
	Name     string `form:"name"`
	LastName string `form:"last_name"`
	Password string `form:"password" binding:"required" json:"-"`
}

var user User

// BeforeSave - hook being executed before each save
func (user *User) BeforeSave() (err error) {
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

// EncryptPassword generates hash from password
func EncryptPassword(password string) (string, error) {
	var hash []byte
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// GetUserByEmail returns user by email
func GetUserByEmail(email string) User {
	var user User
	GetDB().Where("email = ?", email).First(&user)
	return user
}

// GetUsersByName returns user collection by name
func GetUsersByName(name string) []User {
	var users []User
	GetDB().Where("name LIKE ?", "%"+name+"%").Find(&users)
	return users
}

// GetUsersByLastName returns user collection by last name
func GetUsersByLastName(lastName string) []User {
	var users []User
	GetDB().Where("last_name LIKE ?", "%"+lastName+"%").Find(&users)
	return users
}

// GetUserByToken returns user by token
func GetUserByToken(token string) User {
	var session Session
	GetDB().Preload("User").Where("token = ?", token).Find(&session)
	return session.User
}
