package models

type MsgNum struct {
	ID       string `gorm:"column:id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	MsgCount int    `gorm:"column:msg_count" json:"msgCount"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
}
