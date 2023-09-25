package middlewares

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			//允许所有站点访问
			return true
		},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodHead, http.MethodPatch},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "token"},
	})
}
