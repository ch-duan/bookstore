package model

import (
	"bookstore/utils"
	"log"
)

//OrderItem 结构
type OrderItem struct {
	ID             int     //订单项的id
	Count          int     //图书数量
	Amount         float64 //图书金额小计
	Title          string  //图书名
	Author         string  //图书作者
	Price          float64 //图书价格
	Classification string  //分类
	Publisher      string  //出版商
	ImgPath        string  `gorm:"column:imgpath"` //图书图片路径
	OrderID        string  `gorm:"column:orderid"` //订单项所属的订单
}

//TableName 修改gorm默认表名
func (OrderItem) TableName() string {
	return "orderitems"
}

//AddOrderItem 添加订单项
func (OrderItem *OrderItem) AddOrderItem() error {
	err := utils.Db.Create(&OrderItem).Error
	return err
}

//QueryOrderItemsByOrderID 查询订单中的所有项
func QueryOrderItemsByOrderID(ID interface{}) ([]*OrderItem, error) {
	var orderitems []*OrderItem
	err := utils.Db.Where("orderid=?", ID).Find(&orderitems).Error
	for _, v := range orderitems {
		log.Println("qqq:", v)
	}
	return orderitems, err

}
