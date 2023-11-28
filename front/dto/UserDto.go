package dto

type UserLoginDto struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
