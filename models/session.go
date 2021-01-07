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
	IP     string
	User   User `binding:"-" gorm:"association_autoupdate:false;association_autocreate:false"`
}

// GenerateToken returns hash token
func GenerateToken() string {
	b := make([]byte, 20)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// DeleteOrphanSession deletes currently active sessions from the same IP
func DeleteOrphanSession(ip string) {
	sessions := []Session{}
	GetDB().Debug().Where("ip = ?", ip).Delete(&sessions)
}
