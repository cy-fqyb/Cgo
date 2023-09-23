package models

type User struct {
	UserName string `json:"userName" binding:"required"`
	Id       int    `json:"id"`
	Password string `json:"password,omitempty"`
	Mobile   string `json:"mobile" binding:"mobile"`
}
