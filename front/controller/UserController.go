package controller

import (
	"Cgo/common"
	"github.com/gin-gonic/gin"
)

func UserController(r *gin.RouterGroup) {
	r.GET("/login", common.HandlerFunc(login))
	r.POST("/register", common.HandlerFunc(register))
}

func login(ctx *gin.Context) common.Result {
	return common.R.Success(UserService.Login())
}

func register(ctx *gin.Context) common.Result {
	return common.R.Success(UserService.Register())
}
