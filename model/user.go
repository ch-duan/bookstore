package model

import (
	"bookstore/utils"
	"log"
)

//User 用户结构体
type User struct {
	ID       int
	Username string `gorm:"column:username;type:varchar(100)"` //用户名
	Password string `gorm:"column:password;type:varchar(100)"` //密码
	Email    string `gorm:"column:email;type:varchar(100)"`    //邮箱
	PhoneNum string `gorm:"column:phonenum;type:varchar(100)"` //手机号
}

//GetUser 通过ID获取用户
func GetUser(ID interface{}) (*User, error) {
	user := &User{}
	err := utils.Db.First(&user, ID).Error
	log.Println("测试", user, err)
	return user, err
}

//CheckUsername 查询用户名是否可用
func CheckUsername(username string) error {
	var user User
	err := utils.Db.Where("username=?", username).First(&user).Error
	return err
}

//GetPassword 获取密码
func GetPassword(username string) (User, error) {
	var user User
	result := utils.Db.Where("username=?", username).First(&user)
	return user, result.Error
}
