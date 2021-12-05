package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"strconv"
	"time"
)

func addComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("name")
	Name := iUsername.(string)

	context := ctx.PostForm("context")
	postIdString := ctx.PostForm("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "文章id有误")
		return
	}

	comment := model.Comment{
		PostId:      postId,
		Context:     context,
		Name:        Name,
		CommentTime: time.Now(),
	}
	err = service.AddComment(comment)
	if err != nil {
		fmt.Println("add comment err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}
