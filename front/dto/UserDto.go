package dto

type UserLoginDto struct {
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
type UserApplyDto struct {
	UserId   string `json:"user_id" binding:"required"`
	IsAccept bool   `json:"is_accept" binding:"required"`
	FriendId string `json:"friend_id"`
}
