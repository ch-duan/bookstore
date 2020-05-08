package model

import (
	"bookstore/utils"
	"log"
)

//CartItem 购物项结构体
type CartItem struct {
	ID     int     //购物项的id
	BookID int     `gorm:"column:bookid"` //图书ID，因为使用gorm暂时未找到如何将结构体中结构体字段使用来更新sql
	Book   *Book   `gorm:"-"`             //购物项中的图书信息
	Count  int     //购物项中图书的数量
	Amount float64 //购物项中图书的金额小计，通过计算得到
	CartID string  `gorm:"column:cartid"` //当前购物项属于哪一个购物车
}

//TableName 修改gorm默认表名
func (CartItem) TableName() string {
	return "cartitems"
}

//GetAmount 获取购物项中图书的金额小计，用图书的价格和图书的数量计算得到
func (cartItem *CartItem) GetAmount() float64 {
	//获取当前购物项中图书的价格
	return float64(cartItem.Count) * cartItem.Book.Price
}

//QueryCartItems 查询购物车中所有东西
func QueryCartItems(cartID interface{}) ([]*CartItem, error) {
	var cartitems []*CartItem
	err := utils.Db.Where("cartid=?", cartID).Find(&cartitems).Error
	if err != nil {
		return cartitems, err
	}
	for _, v := range cartitems {
		book := &Book{}
		utils.Db.First(&book, v.BookID)
		v.Book = book
	}
	log.Println("cartItems:", cartitems)
	for _, v := range cartitems {
		log.Println(v)
	}
	return cartitems, nil
}

//QueryCartItem 查询购物项
func QueryCartItem(bookID interface{}) (*CartItem, error) {
	cartItem := &CartItem{}
	err := utils.Db.Where("bookid=?", bookID).First(&cartItem).Error
	log.Println("cartitem:query", err, cartItem)
	book, _ := GetBookByID(bookID)
	cartItem.Book = book
	if err != nil {
		return cartItem, err
	}
	return cartItem, nil
}

//AddCartItem 添加购物项
func (cartItem *CartItem) AddCartItem() error {
	err := utils.Db.Create(&cartItem).Error
	log.Println("cartitem:add", err)
	return err
}

//UpdateCartItem 更新购物项
func (cartItem *CartItem) UpdateCartItem() error {
	err := utils.Db.Save(&cartItem).Error
	log.Println("cartitem:save", err)
	return err
}

//DeleteCartItemByID 删除购物项
func DeleteCartItemByID(ID interface{}) error {
	err := utils.Db.Delete(&CartItem{}, ID).Error
	return err
}

//DeleteCartItemByCartID 删除购物车里所有购物项
func DeleteCartItemByCartID(cartID interface{}) error {
	err := utils.Db.Where("cartid=?", cartID).Delete(&CartItem{}).Error
	return err
}
