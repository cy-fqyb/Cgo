package middlewares

import (
	"net/http"
	"reggie/global"
	"reggie/utils"

	"github.com/gin-gonic/gin"
)

// token鉴权中间件
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("token")
		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "token不能为空",
			})
			return
		}

		// 验证token
		tokenData, err := utils.ParseToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "token解析失败：" + err.Error(),
			})
			return
		}
		ctx.Set("tokenData", tokenData)
		//验证成功接着往下走
		ctx.Next()
		global.Logger.Info("token验证成功")
		return
	}
}
