package model

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
	ImgPath        string  //图书图片路径
	OrderID        string  //订单项所属的订单
}
