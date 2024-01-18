package service

import (
	"Cgo/front/models"
	"errors"
)

type roomService struct{}

var RoomService = roomService{}

func (roomService) CreateRoom(room models.Room) error {
	if err := RoomDao.CreateRoom(room); err != nil {
		return err
	}
	return nil
}

// 用户申请加入房间
func (roomService) ApplyJoinRoom(roomApply models.RoomApply) error {
	if err := RoomDao.ApplyJoinRoom(roomApply); err != nil {
		return err
	}
	return nil
}

// 获取用户申请信息
func (roomService) GetRoomApply(master_id string) ([]models.RoomApply, error) {
	var roomApplyArr []models.RoomApply
	err := RoomDao.GetRoomApply(&roomApplyArr, master_id)
	if err != nil || len(roomApplyArr) <= 0 {
		return nil, errors.New("查询不到申请信息")
	}
	return roomApplyArr, nil
}

// 处理用户申请
func (roomService) HandleRoomApply(roomApply models.RoomApply) error {
	if err := RoomDao.HandleRoomApply(roomApply); err != nil {
		return err
	}
	return nil
}

// 获取群组内所有用户的信息
func (roomService) GetRoomUsers(room_id string) ([]models.Users, error) {
	//FIXME: 查询错了表，应该查询用户群组表，查出改群组中所有的用户，然后根据查出来的用户id去用户表中查出用户信息
	var userArr []models.Users
	err := RoomDao.GetRoomUsers(&userArr, room_id)
	if err != nil || len(userArr) <= 0 {
		return nil, errors.New("查询不到用户信息")
	}
	return userArr, nil
}
