package dao

import (
	"Cgo/global"
	"Cgo/models"
	"errors"
)

type userDao struct{}

var UserDao = new(userDao)

// 用户登录
func (userDao) Login(user *models.User) error {
	// 验证用户名和密码
	//global.DB.Raw("select * from sysuser where UserName =?", user.Name).Scan(&user)

	if r := global.DB.Where(user).Find(user); r.RowsAffected > 0 {
		return nil
	} else {
		return errors.New("登录失败")
	}
}

// 用户注册的函数，通过使用sql影响行数来判断是否插入成功
func (userDao) Register(user models.User) bool {
	if r := global.DB.Create(&user); r.RowsAffected > 0 {
		return true
	} else {
		global.Logger.Info(r.Error.Error())
		return false
	}
}
