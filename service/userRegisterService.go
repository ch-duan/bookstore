package service

import (
	"bookstore/model"
	"bookstore/serializer"
	"bookstore/utils"
)

//UserRegister 管理用户注册服务
type UserRegister struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
	PhoneNum string `form:"phoneNum" json:"phoneNum" gorm:"cloumn:phonenum"`
}

//Register 用户注册
func (userRegister *UserRegister) Register() *serializer.Response {
	user := model.User{
		Username: userRegister.Username,
		Password: userRegister.Password,
		Email:    userRegister.Email,
		PhoneNum: userRegister.PhoneNum,
	}
	err := utils.Db.Create(&user).Error
	if err != nil {
		return &serializer.Response{
			Status: serializer.StatusError,
			Msg:    "注册失败",
		}
	}
	return nil
}
