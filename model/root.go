package model

import "bookstore/utils"

//Root 管理员结构体
type Root struct {
	ID       int
	Username string `gorm:"cloumn:username;type:varchar(100)" form:"username"`
	Password string `gorm:"cloumn:password;type:varchar(100)" form:"password"`
}

//GetRootByID 获取管理员
func GetRootByID(ID interface{}) (*Root, error) {
	root := &Root{}
	err := utils.Db.First(&root, ID).Error
	return root, err
}
