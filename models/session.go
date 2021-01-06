package models

import (
	"crypto/rand"
	"fmt"
)

// Session model
type Session struct {
	Model

	Token  string `form:"token" binding:"required"`
	UserID uint64
	User   User `binding:"-" gorm:"association_autoupdate:false;association_autocreate:false"`
}

// GenerateToken returns hash token
func GenerateToken() string {
	b := make([]byte, 20)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
