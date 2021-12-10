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

func addNestedReply(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	Name := iUsername.(string)

	context := ctx.PostForm("context")
	commentIdString := ctx.PostForm("comment_id")
	commentId, err := strconv.Atoi(commentIdString)
	if err != nil {
		fmt.Println("comment id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "文章id有误")
		return
	}

	nestedReply := model.NestedReply{
		CommentId: commentId,
		Context:   context,
		Name:      Name,
		ReplyTime: time.Now(),
	}
	err = service.AddNestedReply(nestedReply)
	if err != nil {
		fmt.Println("add reply err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

func deleteNestedReply(ctx *gin.Context) {
	nestedReplyIdString := ctx.Param("reply_id")
	nestedReplyId, err := strconv.Atoi(nestedReplyIdString)
	replyNameString, _ := ctx.Get("username")
	nameString, _ := service.GetNameById3(nestedReplyId)
	if replyNameString == nameString {
		if err != nil {
			fmt.Println("reply id string to int err: ", err)
			tool.RespErrorWithDate(ctx, "reply_id格式有误")
			return
		}
		err = service.DeleteNestedReply(nestedReplyId)
		if err != nil {
			fmt.Println("delete reply err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessful(ctx)
	} else {
		tool.RespInternalError(ctx)
	}
}

func nestedReplyLikes(ctx *gin.Context) {
	nestedRplyIdString := ctx.Param("reply_id")
	nestedReplyId, err := strconv.Atoi(nestedRplyIdString)
	if err != nil {
		fmt.Println("reply id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "reply_id格式有误")
		return
	}
	err = service.NestedReplyLikes(nestedReplyId)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

func deleteNestedReply0(ctx *gin.Context) {
	nestedReplyIdString := ctx.Param("reply_id")
	nestedReplyId, err := strconv.Atoi(nestedReplyIdString)
	if err != nil {
		fmt.Println("reply id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "reply_id格式有误")
		return
	}
	err = service.DeleteNestedReply(nestedReplyId)
	if err != nil {
		fmt.Println("delete reply err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}
