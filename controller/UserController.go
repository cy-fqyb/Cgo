package controller

import (
	"Cgo/models"
	"Cgo/service"
	"github.com/gin-gonic/gin"
)

func UserController(r *gin.RouterGroup) {
	// 用户更新
	r.POST("/update", HandlerFunc(updateUser))
}

// updateUser 更新用户数据
func updateUser(ctx *gin.Context) Result {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return R.Fail("更新失败")
	}
	//调用service层方法进行后续操作
	err = service.Update(user)
	if err != nil {
		return R.Fail(err.Error())
	}
	return R.Success("更新用户数据")
}

//
