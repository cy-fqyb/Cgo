package service

import (
	"Cgo/front/dto"
	"Cgo/front/models"
	"Cgo/utils"
)

type userService struct{}

var UserService = userService{}

func (userService) Login(userDto dto.UserLoginDto) (*models.Users, error) {
	//TODO:调用dao中的方法进行登录验证
	user := models.Users{
		Email:    userDto.Email,
		Password: userDto.Password,
	}
	if err := UserDao.Login(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (userService) Register(user models.Users) error {
	user.Id = utils.Md5(user.Email)
	if err := UserDao.Register(user); err != nil {
		return err
	}
	return nil
}

// 获取用户好友
func (userService) GetFriends(user *models.Users) ([]models.Users, error) {
	if friends, err := UserDao.GetFriends(user); err != nil {
		return nil, err
	} else {
		return friends, nil
	}
}

// 获取用户好友信息
func (userService) GetFriendApply(user *models.Users) ([]models.Apply, error) {
	if apply, err := UserDao.GetFriendApply(user); err != nil {
		return nil, err
	} else {
		return apply, nil
	}
}
