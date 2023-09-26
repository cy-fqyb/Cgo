package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/sony/sonyflake"
)

func Md5(password string) string {
	has := md5.Sum([]byte(password))
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func IsSame(password, md5Password string) bool {
	password = Md5(password)
	return password == md5Password
}

// 使用雪花算法生成用户的id
func CreateUserId() uint64 {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	if id, err := flake.NextID(); err != nil {
		fmt.Println("Failed to generate ID: %s\n", err)
		return 0
	} else {
		return id
	}
}
