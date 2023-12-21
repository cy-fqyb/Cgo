package controller

import (
	"Cgo/common"
	"Cgo/front/models"

	"github.com/gin-gonic/gin"
)

func RoomController(r *gin.RouterGroup) {
	r.POST("/createRoom", common.HandlerFunc(createRoom))
	r.GET("/getRoomApply", common.HandlerFunc(getRoomApply))
	r.POST("/applyJoinRoom", common.HandlerFunc(applyJoinRoom))
	r.POST("/handleRoomApplication", common.HandlerFunc(handleRoomApplication))
	r.GET("/getRoomUsers", common.HandlerFunc(getRoomUsers))
}

// @Summary 创建房间
// @Schemes
// @Description 创建房间
// @Tags 前端房间接口
// @Param room body models.Room true "房间信息"
// @Accept json
// @Produce json
// @Success 200 {object} common.RT[string]
// @Router /front/room/createRoom [post]
func createRoom(ctx *gin.Context) common.Result {
	var room models.Room
	if err := ctx.ShouldBindJSON(&room); err != nil {
		return common.R.Fail(err.Error())
	}
	if err := RoomService.CreateRoom(room); err != nil {
		return common.R.Fail(err.Error())
	}
	return common.R.Success("创建房间成功")
}

// @Summary 申请加入房间
// @Schemes
// @Description 申请加入房间
// @Tags 前端房间接口
// @Param room body models.RoomApply true "用户申请信息"
// @Accept json
// @Produce json
// @Success 200 {object} common.RT[string]
// @Router /front/room/applyJoinRoom [post]
func applyJoinRoom(ctx *gin.Context) common.Result {
	var roomApply models.RoomApply
	if err := ctx.ShouldBindJSON(&roomApply); err != nil {
		return common.R.Fail(err.Error())
	}
	if err := RoomService.ApplyJoinRoom(roomApply); err != nil {
		return common.R.Fail(err.Error())
	}
	return common.R.Success("申请成功，等待房主同意")
}

// @Summary 获取房间的申请信息
// @Schemes
// @Description 获取房间的申请信息
// @Tags 前端房间接口
// @Param master_id query string true "房主id"
// @Accept json
// @Produce json
// @Success 200 {object} common.RT[models.RoomApply]
// @Router /front/room/getRoomApply [get]
func getRoomApply(ctx *gin.Context) common.Result {
	master_id := ctx.Query("master_id")
	if master_id == "" {
		return common.R.Fail("房主id不能为空")
	}
	roomApply, err := RoomService.GetRoomApply(master_id)
	if err != nil {
		return common.R.Fail(err.Error())
	}
	return common.R.Success(roomApply)
}

// @Summary 处理用户加群申请
// @Schemes
// @Description 处理用户加群申请
// @Tags 前端房间接口
// @Param room body models.RoomApply true "处理用户申请信息"
// @Accept json
// @Produce json
// @Success 200 {object} common.RT[string]
// @Router /front/room/handleRoomApplication [post]
func handleRoomApplication(ctx *gin.Context) common.Result {
	var roomApply models.RoomApply
	if err := ctx.ShouldBindJSON(&roomApply); err != nil {
		return common.R.Fail(err.Error())
	}
	return common.R.Success("处理成功")
}

// @Summary 获取聊天群中的所有用户
// @Schemes
// @Description 获取所有用户的信息
// @Tags 前端房间接口
// @Param room_id query string true "群组id"
// @Accept json
// @Produce json
// @Success 200 {object} common.RT[models.Users]
// @Router /front/room/getRoomUsers [get]
func getRoomUsers(ctx *gin.Context) common.Result {
	room_id := ctx.Query("room_id")
	if room_id == "" {
		return common.R.Fail("群组id不能为空")
	}
	userArr, err := RoomService.GetRoomUsers(room_id)
	if err != nil {
		return common.R.Fail(err.Error())
	}
	return common.R.Success(userArr)
}
