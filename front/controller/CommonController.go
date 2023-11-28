package controller

import (
	_ "Cgo/front/dto"
	"Cgo/utils"

	"github.com/gin-gonic/gin"
)

type WsDto struct {
	ToId uint64 `json:"to_id"`
	Msg  string `json:"msg"`
	Type string `json:"type"`
}

func FrontConmmonController(r *gin.RouterGroup) {
	r.GET("/ws/:id", WebSocket)
}

// @Summary websocket
// @Schemes
// @Description websocket对话通信接口
// @Tags 前端用户接口
// @Param id path string true "用户id"
// @Param msg body dto.WsDto true "消息体" Example:={"msg": "Hello, World!"}
// @Accept json
// @Produce json
// @Success 200 {object} string 返回结果 200 类型（object就是结构体） 类型 注释
// @Router /front/common/ws/{id} [get]
func WebSocket(ctx *gin.Context) {
	utils.Ws.Socket(ctx)
}
