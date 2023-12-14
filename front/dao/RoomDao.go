package dao

import (
	"Cgo/front/models"
	"Cgo/global"
)

type roomDao struct {
}

var RoomDao = roomDao{}

func (roomDao) CreateRoom(room models.Room) error {
	if err := global.DB.Create(&room).Error; err != nil {
		return err
	}
	return nil
}
func (roomDao) ApplyJoinRoom(roomApply models.RoomApply) error {
	if err := global.DB.Create(&roomApply).Error; err != nil {
		return err
	}
	return nil
}
func (roomDao) GetRoomApply(roomApplyArr *[]models.RoomApply, master_id string) error {
	if err := global.DB.Where("master_id = ?", master_id).Find(roomApplyArr).Error; err != nil {
		return err
	}
	return nil
}

func (roomDao) HandleRoomApply(apply models.RoomApply) error {
	if err := global.DB.Model(&apply).Updates(apply).Error; err != nil {
		return err
	}
	//先把申请信息删除
	if r := global.DB.Delete(&apply); r.RowsAffected == 0 {
		return r.Error
	}
	//如果是同意申请，则把用户加入房间

	return nil
}
