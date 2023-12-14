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
	r.GET("/email", common.HandlerFunc(getEmailCode))
	r.GET("/ws/:id", WebSocket)
}

// @Summary 文件上传接口
// @Description 文件上传接口
// @Tags 前端通用接口
// @Accept json
// @Produce json
// @Param file formData file true "文件"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /common/upload [post]
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
// @Tags 前端通用接口
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
	}
	return common.R.Success(data)
}

// @Summary 发送邮箱验证码
// @Description 发送邮箱验证码
// @Tags 前端通用接口
// @Accept json
// @Produce json
// @Param email query string true "邮箱"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /common/email [get]
func getEmailCode(c *gin.Context) common.Result {
	email := c.Query("email")
	utils.SendEmail(email)
	return common.R.Success("发送成功")
}
