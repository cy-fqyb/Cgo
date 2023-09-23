package dao

import (
	"Cgo/global"
	"Cgo/models"
)

type userDao struct{}

var UserDao = new(userDao)

// 用户登录
func (userDao) Login(user models.User) models.User {
	// 验证用户名和密码
	global.DB.Raw("select * from sysuser where UserName =?", user.UserName).Scan(&user)
	return user
}
