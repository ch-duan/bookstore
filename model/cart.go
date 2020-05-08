package model

import (
	"bookstore/utils"
)

//Cart 购物车结构体
type Cart struct {
	ID          string      //购物车的id
	CartItems   []*CartItem `gorm:"-"`                  //购物车中所有的购物项
	TotalCount  int         `gorm:"column:totalcount"`  //购物车中图书的总数量，通过计算得到
	TotalAmount float64     `gorm:"column:totalamount"` //购物车中图书的总金额，通过计算得到
	UserID      int         `gorm:"column:userid"`      //当前购物车所属的用户
}

//GetTotalCount 获取购物车中图书的总数量
func (cart *Cart) GetTotalCount() int {
	var totalCount int
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

//GetTotalAmount 获取购物车中图书的总金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount

}

//QueryCartByUserID 查询用户购物车信息
func QueryCartByUserID(userID interface{}) (*Cart, error) {
	cart := &Cart{}
	err := utils.Db.Where("userid = ?", userID).First(&cart).Error
	if err != nil {
		return cart, err
	}
	cartItems, err := QueryCartItems(cart.ID)
	if err != nil {
		return cart, err
	}
	cart.CartItems = cartItems
	return cart, nil
}

//AddCart 给用户一个购物车
func (cart *Cart) AddCart() error {
	err := utils.Db.Create(&cart).Error
	return err
}

//UpdateCart 更新用户购物车
func (cart *Cart) UpdateCart() error {
	err := utils.Db.Save(&cart).Error
	return err
}

//CheckCartID 检查购物车ID是否可用
func CheckCartID(cartID string) bool {
	var count int
	utils.Db.Model(&Cart{}).Where("ID=?", cartID).Count(&count)
	if count == 0 {
		return true
	}
	return false
}
