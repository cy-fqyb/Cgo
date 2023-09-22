package utils

import (
	"Cgo/global"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	uuid "github.com/satori/go.uuid"
)

type multipartFileWrapper struct {
	*os.File
}

func (w *multipartFileWrapper) Read(p []byte) (n int, err error) {
	return w.File.Read(p)
}

func (w *multipartFileWrapper) Seek(offset int64, whence int) (int64, error) {
	return w.File.Seek(offset, whence)
}
func UploadAliyunOss(file *multipart.FileHeader) string {
	client, err := oss.New(
		global.ConfigViper.GetString("aliyunOss.endpoint"),
		global.ConfigViper.GetString("aliyunOss.accessKeyId"),
		global.ConfigViper.GetString("aliyunOss.accessKeySecret"))
	if err != nil {
		global.Logger.Info(fmt.Sprintf("Failed to create OSS client, with error: %s", err.Error()))
		os.Exit(-1)
	}
	// src, err := file.Open()
	src, err := CompressImage(file)
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

func CompressImage(fileHeader *multipart.FileHeader) (multipart.File, error) {
	file, err := fileHeader.Open()
	if err != nil {
		global.Logger.Info(fmt.Sprintf("Failed to open file, with error: %s", err.Error()))
		os.Exit(-1)
	}
	defer file.Close()
	originalImage, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	// 创建一个临时文件来保存压缩后的图像
	tempFile, err := os.CreateTemp("", "compressed_image_*.jpg")
	if err != nil {
		return nil, err
	}

	//压缩并保存图像到临时文件
	if err = jpeg.Encode(tempFile, originalImage, &jpeg.Options{Quality: global.ConfigViper.GetInt("image.quality")}); err != nil {
		return nil, err
	}
	// 创建一个模拟的*multipart.FileHeader
	// compressedFileHeader := &multipart.FileHeader{
	// 	Filename: fileHeader.Filename,
	// 	Size:     os.Stat(tempFile),
	// 	Header:   make(map[string][]string),
	// }
	return &multipartFileWrapper{
		tempFile,
	}, err
}
