package models

import (
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type Register struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type User struct {
	Model

	Email    string `form:"email" binding:"required"`
	Name     string `form:"name"`
	LastName string `form:"last_name"`
	Password string `form:"password" binding:"required"`
}
var user User

func (user *User) BeforeSave() (err error) {
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user.Password = string(hash)
	return
}




// User methods
func GetUserByEmail(email string) interface {} {
	user := GetDB().Take(&user, "email = ?", email)
	if user.Error != nil {
		return nil
	}
	return user.Value
}

func GetUsersByName(name string) interface {} {
	users := GetDB().Where("name LIKE ?", "%"+name+"%").Find(&user)
	if users.Error != nil {
		return nil
	}
	return users.Value
}