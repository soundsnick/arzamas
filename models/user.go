package models

import (
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	Model

	Email    string `form:"email" binding:"required"`
	Name     string `form:"name"`
	LastName string `form:"last_name"`
	Password string `form:"password" binding:"required"`
}

var user User

// BeforeSave - hook being executed before each save
func (user *User) BeforeSave() (err error) {
	hashedPassword, err := EncryptPassword(user.Password)
	if err != nil {
		return
	}
	user.Password = hashedPassword
	return
}

// EncryptPassword generates hash from password
func EncryptPassword(password string) (string, error) {
	var hash []byte
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// GetUserByEmail - get user row by email
func GetUserByEmail(email string) User {
	user := User{}
	GetDB().Where("email = ?", email).First(&user)
	return user
}

// GetUsersByName - get user collection by name
func GetUsersByName(name string) interface{} {
	users := GetDB().Where("name LIKE ?", "%"+name+"%").Find(&user)
	if users.Error != nil {
		return nil
	}
	return users.Value
}

// GetUsersByLastName - get user collection by last name
func GetUsersByLastName(lastName string) interface{} {
	users := GetDB().Where("last_name LIKE ?", "%"+lastName+"%").Find(&user)
	if users.Error != nil {
		return nil
	}
	return users.Value
}
