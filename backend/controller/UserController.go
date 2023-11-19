package controller

import (
	"Cgo/backend/models"
	"Cgo/backend/service"
	"Cgo/common"
	"github.com/gin-gonic/gin"
)

// 用户控制器
func UserController(r *gin.RouterGroup) {
	// 用户更新
	r.POST("/update", common.HandlerFunc(updateUser))
	// 删除用户
	r.POST("/delete", common.HandlerFunc(deleteUser))
}

// updateUser 更新用户数据
func updateUser(ctx *gin.Context) common.Result {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return common.R.Fail("更新失败")
	}
	//调用service层方法进行后续操作
	err = service.Update(user)
	if err != nil {
		return common.R.Fail(err.Error())
	}
	return common.R.Success("更新用户数据")
}

// 删除用户
func deleteUser(ctx *gin.Context) common.Result {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return common.R.Fail("删除失败")
	}
	err = service.Delete(user.Id)
	if err != nil {
		return common.R.Fail(err.Error())
	}
	return common.R.Success("删除用户成功")
}
