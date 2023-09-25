package controller

import (
	"Cgo/models"
	"Cgo/service"

	"github.com/gin-gonic/gin"
)

func UserController(r *gin.RouterGroup) {
	// 用户注册
	r.POST("/register", HandlerFunc(register))
}

// 用户注册
func register(ctx *gin.Context) Result {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return R.Fail(err.Error())
	}
	err := service.UserService.Register(user)
	if err != nil {
		return R.Fail(err.Error())
	}
	return R.Success("用户注册成功")
}

// 更新用户数据
func updateUser(ctx *gin.Context) Result {
	// var user models.User
	// if err := ctx.ShouldBindJSON(&user); err != nil {
	// 	return R.Fail(err.Error())
	// }
	// err := service.UserService.UpdateUser(user)
	// if err != nil {
	// 	return R.Fail(err.Error())
	// }
	return R.Success("更新用户数据")
}
