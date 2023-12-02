package dao

import (
	"Cgo/front/models"
	"Cgo/global"

	"gorm.io/gorm"
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

// 获取用户好友申请
func (userDao) GetFriendApply(user *models.Users) ([]models.Users, error) {

	return DaoErrorHandle[[]models.Users](func() (*gorm.DB, []models.Users) {
		var apply []models.Users
		return global.DB.Model(&models.Apply{}).
			Select("users.*").
			Joins("JOIN users ON apply.apply_id = users.id").
			Where("apply.user_id = ?", user.Id).
			Find(&apply), apply
	}), nil
}

// 处理好友的申请
func (userDao) HandleApplication(userId string, friendId string, isAccept bool) bool {
	var deleteApply = func() {
		global.DB.Where("user_id = ? and apply_id = ?", userId, friendId).Delete(&models.Apply{})
	}
	global.Logger.Info("userId:", userId, "friendId:", friendId, "isAccept:", isAccept)
	//根据用户传来的信息判断是否接收好友
	if isAccept && global.DB.Where(" user_id = ? and apply_id = ?", userId, friendId).Find(&models.Apply{}).RowsAffected > 0 {
		if r := global.DB.Create(&models.UserFriend{UserId: userId, FriendId: friendId}); r.RowsAffected > 0 {
			deleteApply()
			return true
		} else {
			return false
		}
	}
	global.Logger.Info("userId:", userId, "friendId:", friendId, "isAccept:", isAccept)
	//不管好友申请有没有通过都要把好友申请给干掉
	return false
}
