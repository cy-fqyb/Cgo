package controller

import (
	"Cgo/backend/dto"
	"Cgo/backend/models"
	"Cgo/backend/service"
	"Cgo/utils"

	"github.com/gin-gonic/gin"
)

func CommonController(r *gin.RouterGroup) {
	// 文件上传接口
	r.POST("/upload", HandlerFunc(uploadFile))
	r.POST("/login", HandlerFunc(login))
	r.POST("/register", HandlerFunc(register))
	r.GET("/ws/:id", WebSocket)
}

// 文件上传
func uploadFile(c *gin.Context) Result {
	file, err := c.FormFile("file")
	if err != nil {
		return R.Fail("上传失败:" + err.Error())
	}
	return R.Success(utils.UploadAliyunOss(file))
}

func WebSocket(ctx *gin.Context) {
	utils.Ws.Socket(ctx)
}

// 用户登录接口
func login(c *gin.Context) Result {
	// 获取请求参数
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return R.Fail(err.Error())
	}
	// 调用服务层方法
	token, err := service.UserService.Login(&user)
	if err != nil {
		return R.Fail(err.Error())
	}
	return R.Success(map[string]any{
		"user":  user,
		"token": token,
	})
}

// register 用户注册
func register(ctx *gin.Context) Result {
	var userDto dto.UserDto
	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		return R.Fail(err.Error())
	}
	err := service.UserService.Register(userDto)
	if err != nil {
		return R.Fail(err.Error())
	}
	return R.Success("用户注册成功")
}
