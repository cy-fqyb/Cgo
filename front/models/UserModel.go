package models

import "time"

//用户基础信息
type Users struct {
	Id         string    `json:"id"`
	Name       string    `json:"name" binding:"required"`
	Password   string    `json:"password" binding:"required"`
	Email      string    `json:"email" binding:"required"`
	Sex        string    `json:"sex" binding:"required"`
	Avatar     string    `json:"avatar"`
	Status     int       `json:"status"`
	Synopsis   string    `json:"synopsis"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime" `
}

//用户好友表
type UserFriend struct {
	UserId     string    `json:"user_id"`
	FriendId   string    `json:"friend_id"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime" `
}

//用户群组表
type UserRoom struct {
	UserId     string    `json:"user_id"`
	RoomId     string    `json:"room_id"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime" `
}
