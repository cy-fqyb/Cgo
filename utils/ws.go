package utils

import (
	"Cgo/global"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ws struct {
	upgrade         websocket.Upgrader
	connectionCount int
	connections     map[string]*websocket.Conn
}

var Ws = ws{
	upgrade: websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	},
	connections: map[string]*websocket.Conn{},
}

func (s *ws) Socket(ctx *gin.Context) {
	conn, err := s.upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	//获取url中的路径参数
	userId := ctx.Param("id")
	//连接数改变
	s.connectionCount++
	//储存连接
	s.connections[userId] = conn

	global.Logger.Info("websocket:有新的连接,现有连接数:", s.connectionCount)
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			global.Logger.Error(err.Error())
		}
		//删除储存的连接
		delete(s.connections, userId)
		s.connectionCount--
		global.Logger.Info("websocket:有连接断开,现有连接数:", s.connectionCount)
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
		var users = []string{"1", "2"}
		s.sendArr(users, messageData)
	}
}

// 单发消息
func (s *ws) sendOne(conn *websocket.Conn, data any) {
	if err := conn.WriteJSON(gin.H{"data": data}); err != nil {
		global.Logger.Error("ws send msg failed:", err.Error())
		return
	}
}

// 群发消息
func (s *ws) sendArr(userId []string, data any) {
	for _, id := range userId {
		if conn, ok := s.connections[id]; ok {
			s.sendOne(conn, data)
		}
	}
}
