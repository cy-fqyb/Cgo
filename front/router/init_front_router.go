package router

import (
	"Cgo/front/controller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup) {
	controller.UserController(r.Group("/user"))
}
