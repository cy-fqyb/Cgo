package controller

import (
	"Cgo/common"
	"Cgo/global"
	"context"

	"github.com/gin-gonic/gin"
)

func CommonController(r *gin.RouterGroup) {
	// 获取当前在线用户人数
	r.GET("/onlineUserCount", common.HandlerFunc(getOnlineUserCount))
}

// @Summary 获取当前在线用户人数
// @Schemes
// @Description 获取当前在线用户人数
// @Tags 后台公共接口
// @Accept json
// @Produce json
// @Success 200 {object} common.RT[int]
// @Router /backend/common/onlineUserCount [get]
func getOnlineUserCount(_ *gin.Context) common.Result {
	//从redis set 中获取当前在线用户人数
	count := global.Rdb.SCard(context.Background(), "online")
	//返回当前在线用户人数
	return common.R.Success(count.Val())
}
