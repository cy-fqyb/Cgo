package router

import (
	_ "Cgo/docs"
	"Cgo/global"
	"Cgo/middlewares"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	port := global.ConfigViper.GetInt("server.port")
	if port == 0 {
		port = 8080
	}
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// r.Use(middlewares.GinLogger(global.Logger), middlewares.GinRecovery(global.Logger, true))
	r.Use(middlewares.Cors())
	//静态资源
	InitStatic(r)
	//公共路由组
	InitRouterList(r)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error("http server start error", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.Logger.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error("Server Shutdown:", err)
	}
}
