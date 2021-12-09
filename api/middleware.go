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
	if cookie, err := ctx.Cookie("admin"); err == nil {
		if cookie == "袁鑫浩" {
			ctx.Next()
		}
	}
	// 返回错误
	tool.RespErrorWithDate(ctx, "你好,你不是管理员")
	ctx.Abort()
	return
}
