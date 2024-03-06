package router

import (
	"Cgo/backend/controller"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup) {
	controller.UserController(r.Group("/user"))
	controller.CommonController(r.Group("/common"))
}
