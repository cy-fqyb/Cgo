package router

import (
	"Cgo/controller"
	"Cgo/middlewares"

	"github.com/gin-gonic/gin"
)

// InitRouterList 创建并初始化路由分组
func InitRouterList(r *gin.Engine) {
	controller.CommonController(r.Group("/common"))
	controller.UserController(r.Group("/user", middlewares.JWTAuth()))
}
