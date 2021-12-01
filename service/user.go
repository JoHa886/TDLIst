package service

import (
	"TDList/model"
	"TDList/pkg/utils"
	"TDList/serializer"
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

// 注册
func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int64
	model.GlobalDB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "注册失败，用户名已存在",
		}
	}
	user.UserName = service.UserName
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	// 创建用户
	if err := model.GlobalDB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建成功",
	}
}

// 登录
func (service *UserService) Login() serializer.Response {
	var user model.User
	if err := model.GlobalDB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return serializer.Response{
				Status: 400,
				Msg:    "登录失败，用户不存在",
			}
		}
		return serializer.Response{
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}
	//验证密码
	if !user.CheckPassword(service.Password) {
		return serializer.Response{
			Status: 400,
			Msg:    "登录失败，密码错误",
		}
	}
	// 返回一个token
	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "token生成错误:" + err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登录成功",
	}
}
