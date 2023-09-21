package router

import "github.com/gin-gonic/gin"

func InitStatic(r *gin.Engine) {
	// 前端项目静态资源
	r.Static("/img", "./static/img")
	r.StaticFile("/favicon.ico", "./static/dist/favicon.png")
}
