package models

import "time"

//消息记录表
type Message struct {
	FromId     string    `json:"from_id"`
	ToId       string    `json:"to_id"`
	Msg        string    `json:"msg"`
	Type       string    `json:"type"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime" `
}

//消息记录表
type Room struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Avatar     string    `json:"avatar"`
	MasterId   string    `json:"master_id"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime" `
}

//好友申请表
type Apply struct {
	Id         string    `json:"id"`
	UserId     string    `json:"user_id"`
	ApplyId    string    `json:"apply_id"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime" `
}
