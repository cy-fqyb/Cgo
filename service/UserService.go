package service

import (
	"Cgo/dao"
	"Cgo/models"
	"Cgo/utils"
	"errors"
)

type userService struct{}

var UserService = new(userService)

func (userService) Login(users *models.User) (string, error) {
	if err := dao.UserDao.Login(users); err != nil {
		return "", err
	}
	token := utils.CreateToken(map[string]any{
		"username": users.Name,
		"id":       users.Id,
	})
	return token, nil
}

// 用户注册的service函数
func (userService) Register(users models.User) error {
	if dao.UserDao.Register(users) {
		return nil
	} else {
		return errors.New("用户注册失败")
	}
}
