package api

import (
	"bookstore/model"
	"bookstore/serializer"
	"bookstore/service"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//RootLogin 管理员登录
func RootLogin(c *gin.Context) {
	var service service.RootLoginService
	if err := c.ShouldBind(&service); err == nil {
		if root, response := service.RootLogin(); response != nil {
			c.HTML(http.StatusOK, "rootLogin.html", response)
		} else {
			s := sessions.Default(c)
			s.Set("rootID", root.ID)
			s.Save()
			log.Println(s.Get("rootID"), root)
			c.HTML(http.StatusOK, "manager.html", serializer.Response{
				Status: serializer.StatusOK,
				Msg:    "管理员登录成功",
				Data:   root,
				Error:  "",
			})
		}
	}
}

//CurrentRoot 获取当前用户
func CurrentRoot(c *gin.Context) *model.Root {
	session := sessions.Default(c)
	ID := session.Get("rootID")
	if ID != nil {
		root, err := model.GetRootByID(ID)
		if err == nil {
			return root
		}
	}
	return nil
}
