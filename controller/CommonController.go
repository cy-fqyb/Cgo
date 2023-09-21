package controller

import (
	"Cgo/utils"

	"github.com/gin-gonic/gin"
)

func CommonController(r *gin.RouterGroup) {
	// 文件上传接口
	r.POST("/upload", HandlerFunc(uploadFile))
	r.GET("/ws/:id", WebSocket)
}

// 文件上传
func uploadFile(c *gin.Context) Result {
	file, err := c.FormFile("file")
	if err != nil {
		return R.Fail("上传失败:" + err.Error())
	}
	return R.Success(utils.UploadAliyunOss(file))
}

func WebSocket(ctx *gin.Context) {
	utils.Ws.Socket(ctx)
}
