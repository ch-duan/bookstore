package service

import (
	"bookstore/model"
)

//QueryBooks 查询图书
func QueryBooks(title string, pageNum int, pageSize int) (*model.Page, error) {
	page, err := model.GetPage(title, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return page, nil
}
