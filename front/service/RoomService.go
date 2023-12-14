package service

import (
	"Cgo/front/models"
	"errors"
)

type roomService struct {
}

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
