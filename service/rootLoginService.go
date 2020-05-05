package service

import (
	"bookstore/model"
	"bookstore/serializer"
	"bookstore/utils"
)

// RootLoginService 管理员登录管理服务
type RootLoginService struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

//RootLogin 登录
func (service *RootLoginService) RootLogin() (model.Root, *serializer.Response) {
	var root model.Root
	if err := utils.Db.Where("username = ? AND password = ?", service.Username, service.Password).First(&root).Error; err != nil {
		return root, &serializer.Response{
			Status: serializer.StatusError,
			Msg:    "用户名或密码错误",
		}
	}
	return root, nil
}
