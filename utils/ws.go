package utils

import (
	"Cgo/global"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type ws struct {
	upgrade websocket.Upgrader
}

var Ws = ws{
	upgrade: websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	},
}

func (s *ws) Socket(ctx *gin.Context) {
	conn, err := s.upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			global.Logger.Error(err.Error())
		}
	}(conn)

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		var messageData map[string]interface{}
		err = json.Unmarshal(p, &messageData)
		if err != nil {
			global.Logger.Error("data jsonUnmarshal failed:", err.Error())
			return
		}
		if err := conn.WriteJSON(gin.H{"code": 1, "data": messageData}); err != nil {
			global.Logger.Error("ws send msg failed:", err.Error())
			return
		}
	}
}
