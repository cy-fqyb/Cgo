package router

import (
	backendRouter "Cgo/backend/router"
	"Cgo/common/controller"
	frontRouter "Cgo/front/router"

	"github.com/gin-gonic/gin"
)

// InitRouterList 创建并初始化路由分组
func InitRouterList(r *gin.Engine) {
	controller.CommonController(r.Group("/common"))
	backendRouter.Init(r.Group("/backend"))
	frontRouter.Init(r.Group("/front"))
}
