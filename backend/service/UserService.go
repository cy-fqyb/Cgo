package service

import (
	"Cgo/backend/dao"
	"Cgo/backend/dto"
	"Cgo/backend/models"
	"Cgo/utils"
	"errors"
)

type userService struct{}

var UserService = new(userService)

func (userService) Login(users *models.User) (string, error) {
	users.Password = utils.Md5(users.Password)
	if err := dao.UserDao.Login(users); err != nil {
		return "", err
	}
	token := utils.CreateToken(map[string]any{
		"username": users.Name,
		"id":       users.Id,
	})
	return token, nil
}

// Register 用户注册的service函数
func (userService) Register(userDto dto.UserDto) error {
	if userDto.Password == "" {
		userDto.Password = "123456u"
	}
	user := models.User{
		Name:     userDto.Name,
		Id:       utils.CreateUserId(),
		Password: utils.Md5(userDto.Password),
		Email:    userDto.Email,
		Sex:      userDto.Sex,
		Avatar:   userDto.Avatar,
		Status:   userDto.Status,
	}
	if dao.UserDao.Register(user) {
		return nil
	} else {
		return errors.New("用户注册失败")
	}
}

func Update(user models.User) error {
	err := dao.UserDao.Update(user)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// 删除用户
func Delete(id uint64) error {
	err := dao.UserDao.Delete(id)
	if err != nil {
		return err
	} else {
		return nil
	}
}
