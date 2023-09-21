package utils

import (
	"Cgo/global"
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	uuid "github.com/satori/go.uuid"
)

func UploadAliyunOss(file *multipart.FileHeader) string {
	client, err := oss.New(
		global.ConfigViper.GetString("aliyunOss.endpoint"),
		global.ConfigViper.GetString("aliyunOss.accessKeyId"),
		global.ConfigViper.GetString("aliyunOss.accessKeySecret"))
	if err != nil {
		global.Logger.Info(fmt.Sprintf("Failed to create OSS client, with error: %s", err.Error()))
		os.Exit(-1)
	}
	src, err := file.Open()
	if err != nil {
		global.Logger.Info(fmt.Sprintf("Failed to open file, with error: %s", err.Error()))
		os.Exit(-1)
	}
	defer src.Close()
	bucket, _ := client.Bucket(global.ConfigViper.GetString("aliyunOss.bucket"))
	fileSuffix := path.Ext(file.Filename)
	//拼接上传文件的名字并指定上传的路径
	name := global.ConfigViper.GetString("aliyunOss.dir") + uuid.NewV4().String() + fileSuffix
	err = bucket.PutObject(name, src)
	if err != nil {
		global.Logger.Info(fmt.Sprintf("Failed to upload file, with error: %s", err.Error()))
		os.Exit(-1)
	}
	urlEndpoint := strings.TrimPrefix(global.ConfigViper.GetString("aliyunOss.endpoint"), "https://")
	global.Logger.Info(fmt.Sprintf("Successfully uploaded file: %s", file.Filename))
	return fmt.Sprintf("https://%s.%s/%s", global.ConfigViper.GetString("aliyunOss.bucket"), urlEndpoint, name)
}
