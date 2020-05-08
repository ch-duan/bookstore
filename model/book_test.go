package model

import (
	"bookstore/utils"
	"fmt"
	"log"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println("开始测试")
	order := &Order{}
	var count int
	log.Println(order)
	log.Println(order.ID == "")
	utils.Db.Where("ID=?", "52fdfc07-2182-454f-563f-5f0f9a621d72").First(&order)
	utils.Db.Model(&order).Where("ID=?", "").Count(&count)
	log.Println(order, count)
	log.Println(order.ID == "")
	fmt.Println("结束测试")
}
