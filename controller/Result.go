package controller

import (
	"Cgo/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Result 返回结构体
type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Err  error       `json:"errs"`
}

func (r *Result) Success(data any) Result {
	return Result{Code: 1, Msg: "ok", Data: data}
}

func (r *Result) Fail(msg string) Result {
	return Result{Code: 0, Msg: msg, Data: nil}
}

type HandlerFuncWrapper func(ctx *gin.Context) Result

func HandlerFunc(wrapper HandlerFuncWrapper) gin.HandlerFunc {
	// 日志处理
	// 异常处理
	// 此处类似Java中的切面
	return func(ctx *gin.Context) {

		// 实际的代码  GetById(ctx *gin.Context)
		resp := wrapper(ctx)

		// 统一处理异常
		if resp.Err != nil {
			global.Logger.Error(resp.Err.Error())
		}

		// 响应客户端
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
			"data": resp.Data,
		})
	}
}

var R = Result{}
