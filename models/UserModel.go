package models

import "time"

type User struct {
	Name       string    `json:"name" binding:"required"`
	Id         int       `json:"id"`
	Password   string    `json:"password,omitempty"`
	Email      string    `json:"email" binding:"required"`
	Sex        string    `json:"sex" binding:"required"`
	Avatar     string    `json:"avatar"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime" `
}
