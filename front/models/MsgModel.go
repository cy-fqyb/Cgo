package models

type MsgNum struct {
	ID       string `gorm:"column:id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	MsgCount int    `gorm:"column:msg_count" json:"msgCount"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
}
type MsgHistory struct {
	ID         string `gorm:"column:id" json:"id"`
	FromID     string `gorm:"column:from_id" json:"fromId"`
	ToID       string `gorm:"column:to_id" json:"toId"`
	Msg        string `gorm:"column:content" json:"content"`
	CreateTime string `gorm:"column:create_time" json:"createTime"`
	IsRead     int    `gorm:"column:is_read" json:"isRead"`
	Name       string `gorm:"column:name" json:"name"`
	Avatar     string `gorm:"column:avatar" json:"avatar"`
}
