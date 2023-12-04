package controller

import (
	"Cgo/front/models"
	"Cgo/global"
	"Cgo/utils"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 以下为websocket的业务逻辑
type ws struct {
	upgrade     websocket.Upgrader
	connections map[string]*utils.Connection
}

var Ws = ws{
	upgrade: websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	},
	connections: map[string]*utils.Connection{},
}

func (s *ws) Socket(ctx *gin.Context) {
	wsConn, err := s.upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		global.Logger.Error("websocket:连接失败:", err.Error())
	}
	conn, err := utils.InitConnection(wsConn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": 0, "msg": err.Error()})
		return
	}
	//获取url中的路径参数
	userId := ctx.Param("id")
	//TODO:用户连接数可以用redis来存储，这样可以实现分布式
	//储存连接
	s.connections[userId] = conn
	global.Rdb.SAdd(context.Background(), "online", userId)
	global.Logger.Info("websocket:有新的连接,现有连接数:", s.GetOnlineUserNum())
	defer func(conn *utils.Connection) {
		conn.Close()
		//删除储存的连接
		delete(s.connections, userId)
		global.Rdb.SRem(context.Background(), "online", userId)
		global.Logger.Info("websocket:有连接断开,现有连接数:", s.GetOnlineUserNum())
	}(conn)

	for {
		p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		var messageData map[string]interface{}
		err = json.Unmarshal(p, &messageData)
		if err != nil {
			global.Logger.Error("data jsonUnmarshal failed:", err.Error())
			return
		}
		if messageData["type"] == "1" {
			//单发消息
			s.sendOne(messageData, userId)
		}
		if messageData["type"] == "2" {
			//群发消息
			s.sendArr(messageData["to_id"].([]string), messageData)
		}
	}
}

// 单发消息
func (s *ws) sendOne(dataMap map[string]any, fromId string) {

	//判断发送的用户是否在线
	if !s.userIsOnline(dataMap["to_id"].(string)) {
		if err := s.connections[fromId].WriteMessage(convertMapToBytes(gin.H{"data": "对方不在线"})); err != nil {
			global.Logger.Error("ws send msg failed:", err.Error())
			return
		}
		return
	}
	//TODO:判断是否是好友,先放着吧
	// 存储消息
	global.DB.Table("msg").Create(&models.Message{FromId: fromId, ToId: dataMap["to_id"].(string), Msg: dataMap["msg"].(string), Type: "1"})
	if err := s.connections[dataMap["to_id"].(string)].WriteMessage(convertMapToBytes(gin.H{"data": dataMap})); err != nil {
		global.Logger.Error("ws send msg failed:", err.Error())
		return
	}
}

// 群发消息
func (s *ws) sendArr(userId []string, data map[string]any) {
	for _, id := range userId {
		if _, ok := s.connections[id]; ok {
			//FIXME:如何取出用户所在的群组的所有用户的id，因为调用sendOne如果一个用户不在线就会返回一个状态提示信息，
			//导致如果一个用户在群发的时候如果很多用户不在线会多次收到状态提示信息
			s.sendOne(data, id)
		}
	}
}

// 获取在线用户数
func (s *ws) GetOnlineUserNum() int64 {
	return global.Rdb.SCard(context.Background(), "online").Val()
}

// 判断用户是否在线
func (s *ws) userIsOnline(userId string) bool {
	return global.Rdb.SIsMember(context.Background(), "online", userId).Val()
}

// map转换为[]byte
func convertMapToBytes(data gin.H) []byte {
	jsonPayload, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return jsonPayload
}
