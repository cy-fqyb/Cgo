package cmd

import (
	"Cgo/config"
	"Cgo/global"
	"Cgo/router"
	"context"
)

func Start() {
	//初始化配置文件
	config.InitConfig()
	//初始化日志组件
	global.Logger = config.InitLogger()
	//初始化自定义验证器
	go config.InitValidator()
	//初始化数据库组件
	global.DB = config.InitDb()
	//初始化Redis
	global.Rdb = config.InitRdb()
	//初始化路由
	router.InitRouter()
}

func End() {
	if result, err := global.Rdb.Del(context.Background(), "online").Result(); err == nil {
		global.Logger.Info("在线用户连接id删除成功，删除个数为:", result)
	}

	global.Logger.Info("服务器停止")
}
