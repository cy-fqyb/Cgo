package models

import "time"

type RoomApply struct {
	MasterId   string    `json:"master_id" binding:"required"`
	UserId     string    `json:"user_id" binding:"required"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"autoUpdateTime"`
}
