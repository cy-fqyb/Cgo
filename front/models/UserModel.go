package models

import "time"

type User struct {
	Id         uint64    `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Role       string    `json:"role"`
	Status     string    `json:"status"`
	CreateTime time.Time `json:"create_time,omitempty" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time,omitempty" gorm:"autoUpdateTime" `
}
