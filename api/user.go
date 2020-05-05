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

//CheckUsername 检查用户名
func CheckUsername(c *gin.Context) {
	username := c.PostForm("username")
	if err := model.CheckUsername(username); err != nil {
		c.JSON(http.StatusOK, "<font style='color:green'>用户名可用！</font>")
	} else {
		c.JSON(http.StatusOK, "用户名不可用！")
	}
}

//UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var service service.UserRegister
	if err := c.ShouldBind(&service); err == nil {
		if response := service.Register(); response == nil {
			c.HTML(http.StatusOK, "register_success.html", serializer.Response{
				Status: serializer.StatusOK,
				Msg:    "注册成功",
				Data:   service,
				Error:  "",
			})
			return
		}
	}
	c.HTML(http.StatusOK, "register.html", serializer.Response{
		Status: serializer.StatusError,
		Msg:    "注册失败",
	})
}

//UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		if user, response := service.Login(); response != nil {
			c.HTML(http.StatusOK, "login.html", response)
		} else {
			s := sessions.Default(c)
			s.Set("userID", user.ID)
			s.Save()
			c.HTML(http.StatusOK, "login_success.html", serializer.Response{
				Status: serializer.StatusOK,
				Msg:    "用户登录成功",
				Data:   user,
				Error:  "",
			})
		}
	} else {
		c.HTML(http.StatusOK, "login.html", serializer.Response{
			Status: serializer.StatusError,
			Msg:    "用户名或密码错误",
		})
	}
}

//CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	session := sessions.Default(c)
	ID := session.Get("userID")
	log.Println("ID", ID)
	if ID != nil {
		user, err := model.GetUser(ID)
		if err == nil {
			return user
		}
	}
	return nil
}

//Logout 用户登出
func Logout(c *gin.Context, key string) {
	s := sessions.Default(c)
	s.Delete(key)
	s.Save()
	c.Redirect(http.StatusMovedPermanently, "http://localhost:8080")
}
