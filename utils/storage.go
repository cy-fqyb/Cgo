package utils

import (
	"Cgo/global"
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io"
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
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			return
		}
	}(src)

	//判断上传的文件类型
	isImg := isImage(file)
	bucket, _ := client.Bucket(global.ConfigViper.GetString("aliyunOss.bucket"))
	fileSuffix := path.Ext(file.Filename)
	//拼接上传文件的名字并指定上传的路径
	name := global.ConfigViper.GetString("aliyunOss.dir") + uuid.NewV4().String() + fileSuffix
	if isImg {
		imgData, err := CompressImage(src)
		if err != nil {
			global.Logger.Info(fmt.Sprintf("Failed to compress image, with error: %s", err.Error()))
		}
		if err = bucket.PutObject(name, bytes.NewReader(imgData)); err != nil {
			global.Logger.Info(fmt.Sprintf("Failed to upload file, with error: %s", err.Error()))
			os.Exit(-1)
		}

	} else {
		//不是图片的话直接上传
		err = bucket.PutObject(name, src)
	}
	if err != nil {
		global.Logger.Info(fmt.Sprintf("Failed to upload file, with error: %s", err.Error()))
		os.Exit(-1)
	}
	urlEndpoint := strings.TrimPrefix(global.ConfigViper.GetString("aliyunOss.endpoint"), "https://")
	global.Logger.Info(fmt.Sprintf("Successfully uploaded file: %s", file.Filename))
	return fmt.Sprintf("https://%s.%s/%s", global.ConfigViper.GetString("aliyunOss.bucket"), urlEndpoint, name)
}

// 压缩图片的函数
func CompressImage(file io.Reader) ([]byte, error) {
	originalImage, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	// 创建一个缓冲区来保存压缩后的图像
	var compressedImageBuffer bytes.Buffer

	// 压缩并保存图像到缓冲区
	if err := jpeg.Encode(&compressedImageBuffer, originalImage, &jpeg.Options{Quality: global.ConfigViper.GetInt("image.quality")}); err != nil {
		return nil, err
	}

	// 获取压缩后的图像数据
	compressedImageData := compressedImageBuffer.Bytes()

	return compressedImageData, nil
}

func isImage(file *multipart.FileHeader) bool {
	// // 获取文件的 MIME 类型
	contentType := file.Header.Get("Content-Type")
	// 判断 MIME 类型是否以 "image/" 开头
	return strings.HasPrefix(contentType, "image/")
}
