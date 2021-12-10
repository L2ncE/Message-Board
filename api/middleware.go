package api

import (
	"github.com/gin-gonic/gin"
	"message-board/tool"
)

func auth(ctx *gin.Context) {
	username, err := ctx.Cookie("username")
	if err != nil {
		tool.RespErrorWithDate(ctx, "游客你好！没有您的信息,请先登录!")
		ctx.Abort()
	}

	ctx.Set("username", username)
	ctx.Next()
}

func admin(ctx *gin.Context) {
	username, err := ctx.Cookie("username")
	if err != nil {
		tool.RespErrorWithDate(ctx, "游客你好！没有您的信息,请先登录!")
		ctx.Abort()
	}
	iUsername, _ := ctx.Get("username")
	username = iUsername.(string)
	if username != "袁鑫浩" {
		tool.RespErrorWithDate(ctx, "你好,你不是管理员")
		ctx.Abort()
	}
	ctx.Next()
	return
}
