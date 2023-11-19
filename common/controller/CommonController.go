package controller

import (
	"Cgo/backend/dto"
	"Cgo/backend/models"
	"Cgo/backend/service"
	"Cgo/common"
	"Cgo/utils"

	"github.com/gin-gonic/gin"
)

func CommonController(r *gin.RouterGroup) {
	// 文件上传接口
	r.POST("/upload", common.HandlerFunc(uploadFile))
	r.POST("/login", common.HandlerFunc(login))
	r.POST("/register", common.HandlerFunc(register))
	r.GET("/ws/:id", WebSocket)
}

// 文件上传
func uploadFile(c *gin.Context) common.Result {
	file, err := c.FormFile("file")
	if err != nil {
		return common.R.Fail("上传失败:" + err.Error())
	}
	return common.R.Success(utils.UploadAliyunOss(file))
}

func WebSocket(ctx *gin.Context) {
	utils.Ws.Socket(ctx)
}

// 用户登录接口
func login(c *gin.Context) common.Result {
	// 获取请求参数
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return common.R.Fail(err.Error())
	}
	// 调用服务层方法
	token, err := service.UserService.Login(&user)
	if err != nil {
		return common.R.Fail(err.Error())
	}
	return common.R.Success(map[string]any{
		"user":  user,
		"token": token,
	})
}

// register 用户注册
func register(ctx *gin.Context) common.Result {
	var userDto dto.UserDto
	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		return common.R.Fail(err.Error())
	}
	err := service.UserService.Register(userDto)
	if err != nil {
		return common.R.Fail(err.Error())
	}
	return common.R.Success("用户注册成功")
}
