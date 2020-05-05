package middleware

import (
	"bookstore/model"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		ID := session.Get("userID")
		if ID != nil {
			user, err := model.GetUser(ID)
			if err == nil {
				c.Set("user", user)
			}
		}
		rootID := session.Get("rootID")
		log.Println("session:rootid:", rootID)
		if rootID != nil {
			root, err := model.GetRootByID(rootID)
			log.Println("sss", err)
			if err == nil {
				c.Set("root", &root)
			}
		}
		c.Next()
	}
}

//AuthUserRequired 验证用户是否登录
func AuthUserRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		ID := session.Get("userID")
		if ID != nil {
			c.Next()
			return
		}
		c.HTML(http.StatusOK, "login.html", nil)
		c.Abort()
	}
}

//AuthRootRequired 验证管理员是否登录
func AuthRootRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		ID := session.Get("rootID")
		if ID != nil {
			c.Next()
			return
		}
		c.HTML(http.StatusOK, "rootLogin.html", nil)
		c.Abort()
	}
}
