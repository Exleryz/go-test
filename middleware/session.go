package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
)

// Session 初始化session
func Session(secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	sessionNames := []string{"session", os.Getenv("SESSION_NAME")}
	//Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})

	return sessions.SessionsMany(sessionNames, store)
}
