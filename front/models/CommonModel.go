package models

import "time"

// Message 消息记录表
type Message struct {
	FromId     string    `json:"from_id"`
	ToId       string    `json:"to_id"`
	Msg        string    `json:"msg"`
	Type       string    `json:"type"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime" `
}

// Room 消息记录表
type Room struct {
	Name       string    `json:"name"`
	Avatar     string    `json:"avatar"`
	MasterId   string    `json:"master_id"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime" `
}

// Apply 好友申请表
type Apply struct {
	UserId     string    `json:"user_id"`
	ApplyId    string    `json:"apply_id"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime" `
}
