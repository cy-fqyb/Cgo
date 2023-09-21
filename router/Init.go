package router

import (
	"reggie/controller"
	"reggie/middlewares"

	"github.com/gin-gonic/gin"
)

// 创建并初始化路由分组
func InitRouterList(r *gin.Engine) {
	controller.CommonController(r.Group("/common"))
	controller.UserController(r.Group("/user", middlewares.JWTAuth()))
}
