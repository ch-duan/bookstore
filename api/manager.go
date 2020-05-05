package api

import (
	"bookstore/model"
	"bookstore/serializer"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Manager 后台管理
func Manager(c *gin.Context) {
	root := CurrentRoot(c)
	c.HTML(http.StatusOK, "manager.html", serializer.Response{
		Status: serializer.StatusOK,
		Data:   root,
	})
}

//BookManager 后台图书管理
func BookManager(c *gin.Context) {
	books, err := model.GetBooks()
	if err == nil {
		c.HTML(http.StatusOK, "book_manager.html", serializer.Response{
			Data: books,
		})
	}
}

//UpdateOrAddBook 更新或修改图书
func UpdateOrAddBook(c *gin.Context) {
	book := &model.Book{}
	c.ShouldBind(book)
	log.Println("test:", book)
	file, err := c.FormFile("newImgPath")
	if err != nil {
		log.Println("UpdateOrAddBook:图片上传失败", err)
	} else {
		dir, _ := os.Getwd()
		dst := fmt.Sprintf(dir+"/view/static/img/%s", file.Filename)
		c.SaveUploadedFile(file, dst)
		book.ImgPath = "/static/img/" + file.Filename
	}
	if c.PostForm("bookID") == "" {
		book.AddBook()
	} else {
		book.UpdateBook()
	}
	c.Redirect(http.StatusMovedPermanently, "/manager/bookManager")
}

//UpdateBook 要更新图书的信息
func UpdateBook(c *gin.Context) {
	ID := c.Query("bookID")
	log.Println("id", ID)
	if ID == "" {
		c.HTML(http.StatusOK, "book_edit.html", serializer.Response{
			Data: &model.Book{},
		})
	} else {
		book, _ := model.GetBookByID(ID)
		log.Println("test", book, book.ID)
		c.HTML(http.StatusOK, "book_edit.html", serializer.Response{
			Data: book,
		})
	}

}

//DeleteBook 删除图书
func DeleteBook(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Query("bookID"), 10, 0)
	err := model.DeleteBookByID(ID)
	log.Println("DeleteBook", err)
	c.Redirect(http.StatusMovedPermanently, "/manager/bookManager")
}
