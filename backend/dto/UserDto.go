package dto

type UserDto struct {
	Name     string `json:"name" binding:"required"`
	Id       int    `json:"id"`
	Password string `json:"password,omitempty" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Sex      string `json:"sex"`
	Avatar   string `json:"avatar"`
	Status   int    `json:"status"`
}
