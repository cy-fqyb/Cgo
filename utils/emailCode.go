package utils

import (
	"Cgo/global"
	"context"
	"crypto/tls"
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
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
	code := GenerateCode()
	global.Rdb.Set(context.Background(), uemail, code, 5*time.Minute)

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("Cgo <%s>", global.ConfigViper.GetString("email.user")))
	m.SetHeader("To", uemail)
	m.SetHeader("Subject", "欢迎注册CgoIm在线通信平台")
	m.SetBody("text/html", fmt.Sprintf(`
		<h1>欢迎注册CgoIm在线通信平台</h1>
		<h3>您的验证码为:%s，验证码有效时间为5分钟，请尽快完成注册</h3>
	`, code))

	d := gomail.NewDialer(
		global.ConfigViper.GetString("email.host"),
		465, // 使用465端口
		global.ConfigViper.GetString("email.user"),
		global.ConfigViper.GetString("email.password"),
	)

	// 设置TLS配置
	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: true, // 在生产环境中请设置为 false
		ServerName:         global.ConfigViper.GetString("email.host"),
	}

	if err := d.DialAndSend(m); err != nil {
		global.Logger.Error("发送邮件失败:", err.Error())
		return
	}

	global.Logger.Info("邮件发送成功")
}
