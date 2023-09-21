package service

import (
	"reggie/dao"
)

type userService struct{}

var UserService = new(userService)

func (userService) Login() map[string]any {
	user := dao.UserDao.Login("陈烨")
	return map[string]any{
		"username": user.UserName,
		"id":       user.Id,
	}
}
