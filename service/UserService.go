package service

import (
	"Cgo/dao"
	"Cgo/models"
)

type userService struct{}

var UserService = new(userService)

func (userService) Login(users models.User) map[string]any {
	user := dao.UserDao.Login(users)
	return map[string]any{
		"username": user.UserName,
		"id":       user.Id,
	}
}
