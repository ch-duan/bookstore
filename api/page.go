package api

import (
	"bookstore/serializer"
	"bookstore/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Home 首页
func Home(c *gin.Context) {
	user := CurrentUser(c)
	log.Println("user:", user)
	var pageNum int64
	if c.Query("pageNum") == "" {
		pageNum = 1
	} else {
		pageNum, _ = strconv.ParseInt(c.Query("pageNum"), 10, 0)
	}
	search := c.PostForm("search")
	page, err := service.QueryBooks(search, int(pageNum), 4)
	if user != nil {
		page.Username = user.Username
		log.Println("home:", err, page)
		c.HTML(http.StatusOK, "index.html", serializer.Response{
			Status: serializer.StatusOK,
			Data:   page,
		})
	} else {
		log.Println("home:", err, page)
		c.HTML(http.StatusOK, "index.html", serializer.Response{
			Status: serializer.StatusNoLogin,
			Data:   page,
		})
	}
}

//Search 搜索图书
func Search(c *gin.Context) {

}
