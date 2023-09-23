package controller

import (
	"Cgo/models"
	"Cgo/service"

	"github.com/gin-gonic/gin"
)

func UserController(r *gin.RouterGroup) {
	// 用户登录
	r.POST("/login", HandlerFunc(login))
	// 用户注册
	r.POST("/register")
}

// 用户请求函数
func login(c *gin.Context) Result {
	// 获取请求参数
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return R.Fail(err.Error())
	}
	// 调用服务层方法
	u := service.UserService.Login(user)

	return R.Success(u)
}
