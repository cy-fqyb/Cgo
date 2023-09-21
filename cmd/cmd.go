package cmd

import (
	"Cgo/config"
	"Cgo/global"
	"Cgo/router"
)

func Start() {
	//初始化配置文件
	config.InitConfig()
	//初始化日志组件
	global.Logger = config.InitLogger()
	//初始化数据库组件
	global.DB = config.InitDb()
	//初始化Redis
	global.Rdb = config.InitRdb()
	//初始化路由
	router.InitRouter()
}

func End() {
	global.Logger.Info("服务器停止")
}
