package server

import (
	"bookstore/api"
	"bookstore/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	log.Println(os.Getwd())
	//中间件
	r.Use(middleware.Session("secret"))

	r.Static("static", "./view/static")
	r.LoadHTMLGlob("view/pages/**/*")
	r.GET("/", api.Home)
	r.POST("/", api.Home)
	//用户
	user := r.Group("/user")
	{
		user.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", nil)
		})

		user.POST("/login", api.UserLogin)
		user.GET("/register", func(c *gin.Context) {
			c.HTML(http.StatusOK, "register.html", nil)
		})
		user.POST("/register", api.UserRegister)

		user.POST("/checkUsername", api.CheckUsername)

		user.GET("/logout", middleware.AuthUserRequired(), func(c *gin.Context) {
			api.Logout(c, "userID")
		})
	}

	//后台管理
	manager := r.Group("/manager")
	{
		manager.POST("/login", api.RootLogin)
		auth := manager.Group("")
		auth.Use(middleware.AuthRootRequired())
		{
			auth.GET("", api.Manager)
			auth.GET("/logout", func(c *gin.Context) {
				api.Logout(c, "rootID")
			})
			auth.GET("/bookManager", api.BookManager)
			auth.POST("/updateOrAddBook", api.UpdateOrAddBook)
			auth.GET("/updateBook", api.UpdateBook)
			auth.GET("/deleteBook", api.DeleteBook)
		}
	}
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
	return r
}
