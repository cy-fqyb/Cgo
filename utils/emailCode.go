package utils

import (
	"Cgo/global"
	"context"
	"fmt"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

func GenerateCode() string {
	rand.NewSource(time.Now().UnixNano())
	charset := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	code := make([]byte, global.ConfigViper.GetInt("email.codelength"))
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
func SendEmail(uemail string) {
	e := email.NewEmail()
	code := GenerateCode()
	global.Rdb.Set(context.Background(), uemail, code, 5*time.Minute)
	e.From = fmt.Sprintf("Cgo<%s>", global.ConfigViper.GetString("email.user"))
	//设置接收方的邮箱
	e.To = []string{uemail}
	//设置主题 = "这是主题"
	e.Subject = "这是主题"
	e.HTML = []byte(`
		<h1>欢迎注册CgoIm在线通信平台</h1>
		<h3>您的验证码为:` + code + `验证码有效时间为5分钟，请尽快完成注册</h3>
	`)
	//设置服务器相关的配置
	rootEmail := global.ConfigViper.GetString("email.user")
	rootPass := global.ConfigViper.GetString("email.password")
	err := e.Send(global.ConfigViper.GetString("email.host"), smtp.PlainAuth("", rootEmail, rootPass, "smtp.163.com"))
	if err != nil {
		global.Logger.Error("发送邮件失败:", err.Error())
	}
}
