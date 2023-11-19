package dao

import (
	"Cgo/backend/models"
	"Cgo/global"
)

type userDao struct{}

var UserDao = new(userDao)

// 用户登录
func (userDao) Login(user *models.User) error {
	if r := global.DB.Raw("").Scan(user); r.RowsAffected > 0 {
		return nil
	} else {
		return r.Error
	}
}
