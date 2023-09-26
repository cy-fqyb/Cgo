package models

import "time"

type User struct {
	Name       string    `json:"name" binding:"required"`
	Id         uint64    `json:"id"`
	Password   string    `json:"password,omitempty" binding:"required"`
	Email      string    `json:"email" binding:"required"`
	Sex        string    `json:"sex"`
	Avatar     string    `json:"avatar"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime" `
}
