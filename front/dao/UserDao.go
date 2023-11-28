package dao

import (
	"Cgo/front/models"
	"Cgo/global"
)

type userDao struct{}

var UserDao = new(userDao)

// 用户登录
func (userDao) Login(user *models.Users) error {
	if r := global.DB.Where(user).First(user); r.RowsAffected > 0 {
		return nil
	} else {
		return r.Error
	}
}

// 用户注册
func (userDao) Register(user models.Users) error {
	if r := global.DB.Create(&user); r.RowsAffected > 0 {
		return nil
	} else {
		return r.Error
	}
}

// 获取用户好友
func (userDao) GetFriends(user *models.Users) ([]models.Users, error) {
	var friends []models.Users
	//select * from users where id in (select friend_id from user_friends where user_id = 1)
	if r := global.DB.Table("users").Where("id in (select friend_id from user_friend where user_id = ?)", user.Id).Find(&friends); r.RowsAffected > 0 {
		return friends, nil
	} else {
		return nil, r.Error
	}
}

// 获取用户信息
func (userDao) GetUserMsg(user *models.Users) (models.Users, error) {
	if r := global.DB.Where(user).First(user); r.RowsAffected > 0 {
		return *user, nil
	} else {
		return *user, r.Error
	}
}
