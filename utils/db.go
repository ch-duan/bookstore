package utils

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	//mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

//Db is sql
var Db *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:1q2w3e4r@/bookstore?charset=utf8&parseTime=true&loc=Local")
	db.LogMode(true)
	if err != nil {
		log.Println("连接数据库失败", err)
		panic(err)
	}
	if gin.Mode() == "release" {
		Db.LogMode(false)
	}
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)
	Db = db
}
