package controller

import (
	"Cgo/common"
	"Cgo/front/dto"
	"Cgo/front/models"
	"Cgo/global"
	"Cgo/utils"
	"context"

	"github.com/gin-gonic/gin"
)

func UserController(r *gin.RouterGroup) {
	r.POST("/login", common.HandlerFunc(login))
	r.POST("/register", common.HandlerFunc(register))
	r.GET("/getUserFriends", common.HandlerFunc(getFriends))
	r.GET("/getUserFriendApply", common.HandlerFunc(getUserFriendApply))
	r.POST("/handleApplication", common.HandlerFunc(handleApplication))
	r.POST("/deleteFriend", common.HandlerFunc(deleteFriend))
	r.POST("/addFriend", common.HandlerFunc(addFriend))
	r.POST("/updateUserInfo", common.HandlerFunc(updateUserInfo))
}

// @Summary 用户登录接口
// @Schemes
// @Description user login
// @Tags 前端用户接口
// @Param user body dto.UserLoginDto true "Id"     //参数 ：@Param 参数名 位置（query 或者 path或者 body） 类型 是否必需 注释
// @Accept json
// @Produce json
// @Success 200 {object} common.RT[models.Users] 返回结果 200 类型（object就是结构体） 类型 注释
// @Failure 400 {string} string
// @Router /front/user/login [post]
func login(ctx *gin.Context) common.Result {
	var UserDto dto.UserLoginDto
	if err := ctx.ShouldBindJSON(&UserDto); err != nil {
		return common.R.Fail(err.Error())
	}
	user, err := UserService.Login(UserDto)
	if err != nil {
		return common.R.Fail("账号或密码错误")
	}
	token := utils.CreateToken(user.Email)
	return common.R.Success(gin.H{
		"token": token,
		"user":  user,
	})
}

// @Summary 用户注册接口
// @Schemes
// @Description 用户注册接口
// @Tags 前端用户接口
// @Param user body models.UserRegister true "注册用户的信息"     //参数 ：@Param 参数名 位置（query 或者 path或者 body） 类型 是否必需 注释
// @Accept json
// @Produce json
// @Success 200 {object} common.RT[string] 返回结果 200 类型（object就是
// @Router /front/user/register [post]
func register(ctx *gin.Context) common.Result {
	//TODO: 实现邮箱验证码注册用户
	var user models.UserRegister
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return common.R.Fail(err.Error())
	}
	emailCode, err := global.Rdb.Get(context.Background(), user.Email).Result()
	if err != nil {
		return common.R.Fail("验证码已过期")
	}
	if emailCode != user.EmailCode {
		return common.R.Fail("验证码错误")
	}
	if err := UserService.Register(user.Users); err != nil {
		return common.R.Fail("注册失败")
	}
	return common.R.Success("注册成功")
}

// @Summary 获取用户好友
// @Schemes
// @Description 获取用户好友
// @Tags 前端用户接口
// @Param user_id query string true "用户id"     //参数 ：@Param 参数名 位置（query 或者 path或者 body） 类型 是否必需 注释
// @Accept json
// @Produce json
// @Success 200 {object} common.RT[[]models.Users]
// @Router /front/user/getUserFriends [get]
func getFriends(ctx *gin.Context) common.Result {
	userId := ctx.Query("user_id")
	if userId == "" {
		return common.R.Fail("用户id不能为空")
	}
	user := models.Users{
		Id: userId,
	}
	if friends, err := UserService.GetFriends(&user); err != nil {
		return common.R.Fail("获取好友失败: " + err.Error())
	} else {
		return common.R.Success(friends)
	}
}

// @Summary 获取用户好友申请
// @Description 获取用户好友申请
// @Tags 前端用户接口
// @Param userId query string true "用户id"     //参数 ：@Param 参数名 位置（query 或者 path或者 body） 类型 是否必需 注释
// @Accept json
// @Produce json
// @Success 200 {object} common.Result
// @Router /front/user/getUserFriendApply [get]
func getUserFriendApply(ctx *gin.Context) common.Result {
	userId := ctx.Query("userId")
	if userId == "" {
		return common.R.Fail("参数不能为空")
	}
	user := models.Users{
		Id: userId,
	}
	apply, err := UserService.GetFriendApply(&user)
	if err != nil {
		return common.R.Fail(err.Error())
	}
	return common.R.Success(apply)
}

// @Summary 处理用户好友申请
// @Description 处理用户好友申请
// @Tags 前端用户接口
// @Param user body dto.UserApplyDto true "处理申请的参数"     //参数 ：@Param 参数名 位置（query 或者 path或者 body） 类型 是否必需 注释
// @Accept json
// @Produce json
// @Success 200 {object} common.RT[string]
// @Router /front/user/handleApplication [post]
func handleApplication(ctx *gin.Context) common.Result {
	var applyDto dto.UserApplyDto
	if err := ctx.ShouldBindJSON(&applyDto); err != nil {
		return common.R.Fail("参数错误")
	}
	if err := UserService.HandleApplication(applyDto); err != nil {
		return common.R.Fail(err.Error())
	}

	return common.R.Success("添加成功")
}

// @Summary 删除好友
// @Description 删除好友
// @Tags 前端用户接口
// @Param user body models.UserFriend true "处理申请的参数"
// @Accept json
// @Produce json
// @Success 200 {object} common.RT[string]
// @Router /front/user/deleteFriend [post]
func deleteFriend(ctx *gin.Context) common.Result {
	var friend models.UserFriend
	if err := ctx.ShouldBindJSON(&friend); err != nil {
		return common.R.Fail("参数错误: " + err.Error())
	}
	if err := UserService.DeleteFriend(friend.UserId, friend.FriendId); err != nil {
		return common.R.Fail("删除失败: " + err.Error())
	}
	return common.R.Success("删除成功")
}

// 请求添加好友
// @Summary 请求添加好友
// @Description 请求添加好友
// @Tags 前端用户接口
// @Param user body models.UserFriend true "处理申请的参数"
// @Accept json
// @Produce json
// @Success 200 {object} common.RT[string]
// @Router /front/user/addFriend [post]
func addFriend(ctx *gin.Context) common.Result {
	var apply models.Apply
	if err := ctx.ShouldBindJSON(&apply); err != nil {
		return common.R.Fail("参数错误: " + err.Error())
	}
	if err := UserService.RequestAddFriend(apply.UserId, apply.ApplyId); err != nil {
		return common.R.Fail("添加失败: " + err.Error())
	}
	return common.R.Success("添加成功")
}

// @Summary 修改用户信息
// @Description 修改用户信息
// @Tags 前端用户接口
// @Param user body models.Users true "用户信息"
// @Accept json
// @Produce json
// @Success 200 {object} common.RT[string]
// @Router /front/user/updateUserInfo [post]
func updateUserInfo(ctx *gin.Context) common.Result {
	var user models.Users
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return common.R.Fail("参数错误: " + err.Error())
	}
	if err := UserService.UpdateUserInfo(user); err != nil {
		return common.R.Fail("修改失败: " + err.Error())
	}
	return common.R.Success("修改成功")
}
