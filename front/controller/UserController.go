package controller

import (
	"Cgo/common"

	"github.com/gin-gonic/gin"
)

func UserController(r *gin.RouterGroup) {
	r.GET("/login", common.HandlerFunc(login))
	r.POST("/register", common.HandlerFunc(register))
}

// GetList
// @Summary 用户登录接口
// @Schemes
// @Description user login
// @Tags 前端用户接口
// @Param Id query int true "Id"     //参数 ：@Param 参数名 位置（query 或者 path或者 body） 类型 是否必需 注释
// @Accept json
// @Produce json
// @Success 200 {object} common.Result 返回结果 200 类型（object就是结构体） 类型 注释
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Router /front/user/login [get]
func login(ctx *gin.Context) common.Result {
	return common.R.Success(UserService.Login())
}

// @Summary 用户注册接口
// @Schemes
// @Description user register
// @Tags 前端用户接口
// @Param Id query int true "Id"     //参数 ：@Param 参数名 位置（query 或者 path或者 body） 类型 是否必需 注释
// @Accept json
// @Produce json
// @Success 200 {object} common.Result 返回结果 200 类型（object就是
// @Router /front/user/register [post]
func register(ctx *gin.Context) common.Result {
	return common.R.Success(UserService.Register())
}
