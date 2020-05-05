package service

import (
	"bookstore/model"
	"bookstore/serializer"
	"bookstore/utils"
)

//UserLoginService 管理用户登录的服务
type UserLoginService struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

//Login 登录
func (service *UserLoginService) Login() (model.User, *serializer.Response) {
	var user model.User
	if err := utils.Db.Where("username = ? AND password = ?", service.Username, service.Password).First(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: serializer.StatusError,
			Msg:    "用户名或密码错误",
		}
	}
	return user, nil
}
