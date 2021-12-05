package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"strconv"
	"time"
)

func postDetail(ctx *gin.Context) {
	postIdString := ctx.Param("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "post_id格式有误")
		return
	}

	//根据postId拿到post
	post, err := service.GetPostById(postId)
	if err != nil {
		fmt.Println("get post by id err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	//找到它的评论
	comments, err := service.GetPostComments(postId)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("get post comments err: ", err)
			tool.RespInternalError(ctx)
			return
		}
	}

	var postDetail model.PostDetail
	postDetail.Post = post
	postDetail.Comments = comments

	fmt.Println("123")
	tool.RespSuccessfulWithDate(ctx, postDetail)
}

func briefPosts(ctx *gin.Context) {
	posts, err := service.GetPosts()
	if err != nil {
		fmt.Println("get posts err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessfulWithDate(ctx, posts)
}

func addPost(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	name := iUsername.(string)

	context := ctx.PostForm("txt")

	post := model.Post{
		Context:    context,
		Name:       name,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}

	err := service.AddPost(post)
	if err != nil {
		fmt.Println("add post err: ", err)
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}
