package session

import (
	"github.com/soundsnick/arzamas/core"
	"github.com/soundsnick/arzamas/user"
)

// DeleteOrphanSessions deletes currently active sessions from the same IP
func DeleteOrphanSessions(ip string) {
	var sessions []Session
	core.GetDB().Where("ip = ?", ip).Delete(&sessions)
}

// DeleteSessions deletes currently active sessions from the same IP
func DeleteSessions(token string) {
	var sessions []Session
	core.GetDB().Where("token = ?", token).Delete(&sessions)
}

// GetUserByToken returns user by token
func GetUserByToken(token string) user.User {
	var session Session
	core.GetDB().Preload("User").Where("token = ?", token).Find(&session)
	return session.User
}
