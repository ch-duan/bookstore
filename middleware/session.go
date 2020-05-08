package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//Session Session
func Session(key string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(key))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 3600, Path: "/"})
	return sessions.Sessions("my-session", store)
}
