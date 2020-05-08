package api

import (
	"bookstore/model"
	"bookstore/serializer"
	"bookstore/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AddCart 添加图书到购物车
func AddCart(c *gin.Context) {
	user := CurrentUser(c)
	log.Println("AddCart: user:", user)
	bookID := c.Query("bookID")
	book, _ := model.GetBookByID(bookID)
	cart, _ := model.QueryCartByUserID(user.ID)
	log.Println("cart:", cart)
	//购物车存在
	if cart.UserID > 0 {
		cartItem, _ := model.QueryCartItem(bookID)
		if cartItem.ID > 0 {
			cts := cart.CartItems
			for _, v := range cts {
				if v.Book.ID == cartItem.Book.ID {
					v.Count = v.Count + 1
					v.Amount = float64(v.Count) * v.Book.Price
					v.UpdateCartItem()
				}
			}
		} else {
			cartItem := *&model.CartItem{
				Book:   book,
				BookID: book.ID,
				Count:  1,
				Amount: book.Price * 1,
				CartID: cart.ID,
			}
			cartItem.AddCartItem()
			cart.CartItems = append(cart.CartItems, &cartItem)
		}
		cart.TotalCount = cart.GetTotalCount()
		cart.TotalAmount = cart.GetTotalAmount()
		cart.UpdateCart()
	} else {
		cartID := utils.CreateUUID()
		for model.CheckCartID(cartID) == false {
			cartID = utils.CreateUUID()
		}
		var cartItems []*model.CartItem
		cartItem := &model.CartItem{
			Book:   book,
			BookID: book.ID,
			Count:  1,
			Amount: book.Price * 1,
			CartID: cartID,
		}
		cartItems = append(cartItems, cartItem)
		cart1 := &model.Cart{
			ID:        cartID,
			CartItems: cartItems,
			UserID:    user.ID,
		}
		cart1.TotalAmount = cart1.GetTotalAmount()
		cart1.TotalCount = cart1.GetTotalCount()
		cart1.AddCart()
		cartItem.AddCartItem()
	}
	c.JSON(http.StatusOK, serializer.Response{
		Msg: "您刚刚将" + book.Title + "添加到了购物车！",
	})
}

//Cart 查看购物车内容
func Cart(c *gin.Context) {
	user := CurrentUser(c)
	cart, _ := model.QueryCartByUserID(user.ID)
	c.HTML(http.StatusOK, "cart.html", serializer.Response{
		Status: serializer.StatusOK,
		Data: struct {
			*model.Cart
			*model.User
		}{
			Cart: cart,
			User: user,
		},
	})
}

//UpdateCartItem 更新一些购物项
func UpdateCartItem(c *gin.Context) {
	user := CurrentUser(c)
	cartItemID, _ := strconv.Atoi(c.PostForm("cartItemID"))
	bookCount, _ := strconv.Atoi(c.PostForm("bookCount"))
	cart, _ := model.QueryCartByUserID(user.ID)
	cartItems := cart.CartItems
	for _, v := range cartItems {
		if v.ID == cartItemID {
			v.Count = bookCount
			v.Amount = float64(bookCount) * v.Book.Price
			err := v.UpdateCartItem()
			if err != nil {
				log.Println("UpdateCartItem: UpdateCartItem error!", err)
			}
		}
	}
	cart.TotalCount = cart.GetTotalCount()
	cart.TotalAmount = cart.GetTotalAmount()
	err := cart.UpdateCart()
	if err != nil {
		log.Println("UpdateCartItem：UpdateCart error!", err)
	}
	cart, _ = model.QueryCartByUserID(user.ID)
	totalCount := cart.TotalCount
	totalAmount := cart.TotalAmount
	var amount float64
	cIs := cart.CartItems
	for _, v := range cIs {
		if v.ID == cartItemID {
			amount = v.Amount
		}
	}
	data := struct {
		Amount      float64
		TotalAmount float64
		TotalCount  int
	}{
		Amount:      amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
	}
	// response, _ := json.Marshal(data)
	c.JSON(http.StatusOK, data)
}

//DeleteCartItem 删除购物项
func DeleteCartItem(c *gin.Context) {
	user := CurrentUser(c)
	cartItemID, _ := strconv.Atoi(c.Query("cartItemID"))
	cart, _ := model.QueryCartByUserID(user.ID)
	for k, v := range cart.CartItems {
		if v.ID == cartItemID {
			cart.CartItems = append(cart.CartItems[:k], cart.CartItems[k+1:]...)
			err := model.DeleteCartItemByID(cartItemID)
			if err != nil {
				log.Println("DeleteCartItem: DeleteCartItemByID error!", err)
			}
		}
	}
	cart.TotalCount = cart.GetTotalCount()
	cart.TotalAmount = cart.GetTotalAmount()
	cart.UpdateCart()
	Cart(c)
}

//DeleteCart 清空购物车
func DeleteCart(c *gin.Context) {
	user := CurrentUser(c)
	cartID := c.Query("cartID")
	cart, _ := model.QueryCartByUserID(user.ID)
	model.DeleteCartItemByCartID(cartID)
	cart.CartItems = nil
	cart.TotalAmount = 0
	cart.TotalCount = 0
	cart.UpdateCart()
	Cart(c)
}
