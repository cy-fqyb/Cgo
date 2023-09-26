package test

import (
	"Cgo/utils"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	uuid "github.com/satori/go.uuid"
)

func TestDemo(t *testing.T) {
	fmt.Println("OSS Go SDK Version: ", oss.Version)
}

func TestUUid(t *testing.T) {
	fileSuffix := path.Ext("demo.jpg")
	name := uuid.NewV4().String() + fileSuffix
	fmt.Println(name)
}

func TestOss(t *testing.T) {
	endpoint := viper.GetString("aliyunOss.endpoint")
	accessKeyId := viper.GetString("aliyunOss.accessKeyId")
	accessKeySecret := viper.GetString("aliyunOss.accessKeySecret")
	bucketName := viper.GetString("aliyunOss.bucket")
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	err = bucket.PutObjectFromFile("go/demo.txt", "D:\\golang\\reggie\\test\\demo.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func TestMdt(t *testing.T) {
	println(utils.Md5("zs"))
}
