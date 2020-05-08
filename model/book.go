package model

import (
	"bookstore/utils"
)

//Book 图书结构体
type Book struct {
	ID             int     `form:"bookID"`
	Title          string  `form:"title"`  //书名
	Author         string  `form:"author"` //作者
	Price          float64 `form:"price"`  //单价
	Sales          int     `form:"sales"`  //销售数据
	Stock          int     `form:"stock"`  //库存
	Classification string  //分类
	Publisher      string  //出版商
	ImgPath        string  `form:"imgpath" gorm:"column:imgpath"` //图书图片路径
	Ebook          bool    //是否电子书
}

//AddBook 添加图书
func (book *Book) AddBook() error {
	err := utils.Db.Create(&book).Error
	return err
}

//UpdateBook 更新图书
func (book *Book) UpdateBook() error {
	err := utils.Db.Save(&book).Error
	return err
}

//DeleteBookByID 通过ID删除图书
func DeleteBookByID(ID interface{}) error {
	err := utils.Db.Delete(&Book{}, ID).Error
	return err
}

//GetBookByID 通过ID获取图书信息
func GetBookByID(ID interface{}) (*Book, error) {
	book := &Book{}
	err := utils.Db.First(&book, ID).Error
	if err != nil {
		return nil, err
	}
	return book, nil
}

//GetBooks 获取所有图书信息
func GetBooks() ([]*Book, error) {
	var books []*Book
	utils.Db.Find(&books)
	return books, nil
}

//GetPage 获取图书分页信息
func GetPage(title string, pageNum int, pageSize int) (*Page, error) {
	page := &Page{}
	utils.Db.Find(&page.Books).Count(&page.TotalRecurd)
	if page.TotalRecurd%pageSize == 0 {
		page.TotalPageNum = page.TotalRecurd / pageSize
	} else {
		page.TotalPageNum = page.TotalRecurd/pageSize + 1
	}
	utils.Db.Where("title like ?", "%"+title+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&page.Books)
	page.PageNum = pageNum
	page.PageSize = pageSize
	return page, nil
}
