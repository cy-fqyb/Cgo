package controller

import (
	"Cgo/common"

	"github.com/gin-gonic/gin"
)

func MsgController(r *gin.RouterGroup) {
	r.GET("/getMsgNum", common.HandlerFunc(getMsgNum))
	r.GET("/getMsgHistory", common.HandlerFunc(GetMsgHistory))
}

// @Summary 获取用户未读消息数量
// @Schemes
// @Description 获取用户未读消息数量
// @Tags 前端消息接口
// @Param user_id query string true "用户id"
// @Accept json
// @Produce json
// @Success 200 {object} common.RT[models.MsgNum]
// @Router /front/msg/getMsgNum [get]
func getMsgNum(ctx *gin.Context) common.Result {
	// 获取用户未读消息数量
	user_id := ctx.Query("user_id")
	if user_id == "" {
		return common.R.Fail("用户id不能为空")
	}
	results, err := MsgService.GetMsgNum(user_id)
	if err != nil {
		return common.R.Fail(err.Error())
	}
	return common.R.Success(results)
}

func GetMsgHistory(ctx *gin.Context) common.Result {
	// 获取用户未读消息数量
	user_id := ctx.Query("user_id")

	from_id := ctx.Query("from_id")
	if user_id == "" {
		return common.R.Fail("用户id不能为空")
	}
	if from_id == "" {
		return common.R.Fail("好友id不能为空")
	}
	results, err := MsgService.GetMsgHistory(from_id, user_id)
	if err != nil {
		return common.R.Fail(err.Error())
	}
	return common.R.Success(results)
}
