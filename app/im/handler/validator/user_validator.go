package validator

import (
//"liaotian/middlewares/validate/translate"
)

// 登录验证器
type LoginValidator struct {
	Account  string `validate:"required,min=1,max=20"`
	Password string `validate:"required,min=1,max=20"`
}

//注册验证器
type RegisterValidator struct {
	Account  string `validate:"required,min=1,max=20"`
	Name     string `validate:"required,min=1,max=20"`
	Password string `validate:"required,min=1,max=20"`
}

//获取用户信息验证器
type GetUserInfoValidator struct {
	Id int64 `validate:"required,min=1"`
}

//更新用户信息验证器
type UpdateUserInfoValidator struct {
	Id       int64  `validate:"required,min=1"`
	Name     string `validate:"required,min=1,max=20"`
	Password string `validate:"required,min=1,max=20"`
	Avatar   string
}
