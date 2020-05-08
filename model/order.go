package model

import (
	"bookstore/utils"
	"log"
	"time"
)

//Order 订单结构体
type Order struct {
	ID          string    //订单号
	CreateTime  time.Time `gorm:"column:createtime"`  //创建订单的时间
	TotalCount  int       `gorm:"column:totalcount"`  //订单图书总计
	TotalAmount float64   `gorm:"column:totalamount"` //订单总金额
	State       int       //订单状态0未发货，1已发货，2交易完成
	UserID      int       `gorm:"column:userid"` //订单所属用户
}

//GetTime 返回格式化后的时间
func (order *Order) GetTime() string {
	return order.CreateTime.Format("2006-01-02 15:04:05")
}

//NoSend 未发货
func (order *Order) NoSend() bool {
	return order.State == 0
}

//SendComplate 已发货
func (order *Order) SendComplate() bool {
	return order.State == 1
}

//Complate 交易完成
func (order *Order) Complate() bool {
	return order.State == 2
}

//AddOrder 添加订单
func (order *Order) AddOrder() error {
	err := utils.Db.Create(&order).Error
	return err
}

//QueryAllOrder 查询所有订单
func QueryAllOrder() ([]*Order, error) {
	var order []*Order
	err := utils.Db.Find(&order).Error
	return order, err
}

//QueryOrderByUserID 查询用户所有订单
func QueryOrderByUserID(ID interface{}) ([]*Order, error) {
	var orders []*Order
	err := utils.Db.Where("userid=?", ID).Find(&orders).Error
	for _, v := range orders {
		log.Println("QueryOrderByUserID: ", v, v.CreateTime)
	}
	return orders, err
}

//CheckOrderID 检查购物车ID是否可用
func CheckOrderID(orderID string) bool {
	var count int
	utils.Db.Model(&Order{}).Where("ID=?", orderID).Count(&count)
	if count == 0 {
		return true
	}
	return false
}

//UpdateOrderstatus 更新订单
func UpdateOrderstatus(orderID string, state int) error {
	err := utils.Db.Model(&Order{}).Where("ID=?", orderID).Update("state", state).Error
	return err
}
