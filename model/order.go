package model

import (
	"time"
)

//Order 订单结构体
type Order struct {
	ID          string    //订单号
	CreateTime  time.Time //创建订单的时间
	TotalCount  int       //订单图书总计
	TotalAmount float64   //订单总金额
	State       int       //订单状态0未发货，1已发货，2交易完成
	UserID      int       //订单所属用户
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
