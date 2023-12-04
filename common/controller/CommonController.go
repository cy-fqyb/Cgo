package controller

import (
	"Cgo/common"
	"Cgo/global"
	"Cgo/utils"

	"github.com/gin-gonic/gin"
)

func CommonController(r *gin.RouterGroup) {
	// 文件上传接口
	r.POST("/upload", common.HandlerFunc(uploadFile))
	r.GET("/test", common.HandlerFunc(Test))
	r.GET("/ws/:id", WebSocket)
}

// 文件上传
func uploadFile(c *gin.Context) common.Result {
	file, err := c.FormFile("file")
	if err != nil {
		return common.R.Fail("上传失败:" + err.Error())
	}
	return common.R.Success(utils.UploadAliyunOss(file))
}

func WebSocket(ctx *gin.Context) {
	Ws.Socket(ctx)
}

// @Summary 测试接口
// @Description 测试接口
// @Tags 前端用户接口
// @Accept json
// @Produce json
// @Param id path int true "用户id"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /common/test [get]
func Test(c *gin.Context) common.Result {
	email := c.Query("email")
	utils.SendEmail(email)
	data := map[string]any{
		"test": "test",
		"user": global.ConfigViper.GetString("email.user"),
		"code": utils.GenerateCode(),
	}
	return common.R.Success(data)
}
