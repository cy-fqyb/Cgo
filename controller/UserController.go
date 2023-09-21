package controller

import (
	"reggie/service"

	"github.com/gin-gonic/gin"
)

func UserController(r *gin.RouterGroup) {
	// 用户登录
	r.GET("/login", HandlerFunc(login))
	// 用户注册
	r.POST("/register")
}

// 用户请求函数
func login(c *gin.Context) Result {
	// 获取请求参数
	// 调用服务层方法
	user := service.UserService.Login()

	return R.Success(user)
}
