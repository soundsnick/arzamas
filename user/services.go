package user

import (
	"github.com/soundsnick/arzamas/core"
)

// GetByEmail returns user by email
func GetByEmail(email string) User {
	var user User
	core.GetDB().Where("email = ?", email).First(&user)
	return user
}

// GetByName returns user collection by name
func GetByName(name string) []User {
	var users []User
	core.GetDB().Where("name LIKE ?", "%"+name+"%").Find(&users)
	return users
}

// GetByLastName returns user collection by last name
func GetByLastName(lastName string) []User {
	var users []User
	core.GetDB().Where("last_name LIKE ?", "%"+lastName+"%").Find(&users)
	return users
}
