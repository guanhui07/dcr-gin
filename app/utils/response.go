package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 通用的返回
func Response(ctx *gin.Context, code int, msg string, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// SuccessApi Success 成功的请求
func SuccessApi(ctx *gin.Context, data any, msg string) {
	if msg == "" {
		msg = "请求成功"
	}
	Response(ctx, 100, msg, data)
	return
}

// Success 成功的请求
func Success(ctx *gin.Context, data any, msg string) {
	if msg == "" {
		msg = "请求成功"
	}
	Response(ctx, 100, msg, data)
	return
}

// Fail 失败的请求
func Fail(ctx *gin.Context, msg string) {
	Response(ctx, 104, msg, nil)
	return
}

// FailCode Fail 失败的请求
func FailCode(ctx *gin.Context, msg string, code int) {
	if code == 0 {
		code = 104
	}
	Response(ctx, code, msg, nil)
	return
}

// FailMiddleWare 中间件失败的请求
func FailMiddleWare(ctx *gin.Context, msg string) {
	Response(ctx, 104, msg, nil)
	ctx.Abort()
	return
}
