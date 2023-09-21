package dao

import (
	"reggie/global"
	"reggie/models"
)

type userDao struct{}

var UserDao = new(userDao)

// 用户登录
func (userDao) Login(username string) models.User {
	var user models.User
	// 验证用户名和密码
	global.DB.Raw("select * from sysuser where UserName =?", username).Scan(&user)
	return user
}
