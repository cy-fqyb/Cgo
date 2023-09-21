package controller

import (
	"reggie/utils"

	"github.com/gin-gonic/gin"
)

func CommonController(r *gin.RouterGroup) {
	// 文件上传接口
	r.POST("/upload", HandlerFunc(uploadFile))
}
func uploadFile(c *gin.Context) Result {
	file, err := c.FormFile("file")
	if err != nil {
		return R.Fail("上传失败:" + err.Error())
	}
	return R.Success(utils.UploadAliyunOss(file))
}
