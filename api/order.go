package api

import (
	"bookstore/model"
	"bookstore/serializer"
	"bookstore/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//Order 查看订单
func Order(c *gin.Context) {
	user := CurrentUser(c)
	orders, _ := model.QueryOrderByUserID(user.ID)
	c.HTML(http.StatusOK, "order.html", serializer.Response{
		Data: struct {
			*model.User
			Orders []*model.Order
		}{
			User:   user,
			Orders: orders,
		},
	})
}

//GetOrderInfo 查看订单信息
func GetOrderInfo(c *gin.Context) {
	orderID := c.Query("orderID")
	orderItems, _ := model.QueryOrderItemsByOrderID(orderID)
	c.HTML(http.StatusOK, "order_info.html", serializer.Response{
		Data: struct {
			OrderItems []*model.OrderItem
		}{
			OrderItems: orderItems,
		},
	})

}

//TakeOrder 确认收货
func TakeOrder(c *gin.Context) {
	orderID := c.Query("orderID")
	model.UpdateOrderstatus(orderID, 2)
	Order(c)
}

//Checkout 结账
func Checkout(c *gin.Context) {
	user := CurrentUser(c)
	cart, _ := model.QueryCartByUserID(user.ID)
	orderID := utils.CreateUUID()
	for model.CheckOrderID(orderID) == false {
		orderID = utils.CreateUUID()
	}
	time := time.Now()
	order := &model.Order{
		ID:          orderID,
		CreateTime:  time,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      user.ID,
	}
	log.Println("order:", order)
	order.AddOrder()
	cartItems := cart.CartItems
	for _, v := range cartItems {
		if v.ID > 0 {
			orderItem := &model.OrderItem{
				Count:          v.Count,
				Amount:         v.Amount,
				Title:          v.Book.Title,
				Author:         v.Book.Author,
				Price:          v.Book.Price,
				Classification: v.Book.Classification,
				Publisher:      v.Book.Publisher,
				ImgPath:        v.Book.ImgPath,
				OrderID:        orderID,
			}
			orderItem.AddOrderItem()
			book := v.Book
			book.Sales = book.Sales + v.Count
			book.Stock = book.Stock - v.Count
			book.UpdateBook()
		}
	}
	model.DeleteCartItemByCartID(cart.ID)
	cart.CartItems = nil
	cart.TotalAmount = 0
	cart.TotalCount = 0
	cart.UpdateCart()
	c.HTML(http.StatusOK, "checkout.html", serializer.Response{
		Data: struct {
			*model.User
			*model.Order
		}{
			User:  user,
			Order: order,
		},
	})
}
