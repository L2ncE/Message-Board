package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"net/http"
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

	context := ctx.PostForm("context")

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

func deletePost(ctx *gin.Context) {
	postIdString := ctx.Param("post_id")
	postId, err := strconv.Atoi(postIdString)
	postNameString, _ := ctx.Get("username")
	nameString, _ := service.GetNameById(postId)
	if postNameString == nameString {
		if err != nil {
			fmt.Println("post id string to int err: ", err)
			tool.RespErrorWithDate(ctx, "post_id格式有误")
			return
		}
		err = service.DeletePost(postId)
		if err != nil {
			fmt.Println("delete post err: ", err)
			tool.RespInternalError(ctx)
			return
		}
		tool.RespSuccessful(ctx)
	} else {
		tool.RespInternalError(ctx)
	}
}

func changePost(ctx *gin.Context) {
	newPost := ctx.PostForm("newPost")
	ipostId := ctx.PostForm("post_id")
	postId, err := strconv.Atoi(ipostId)
	UpdateTime := time.Now()
	postNameString, _ := ctx.Get("username")
	nameString, _ := service.GetNameById(postId)
	if postNameString == nameString {
		if err != nil {
			fmt.Println("post id string to int err: ", err)
			tool.RespErrorWithDate(ctx, "post_id格式有误")
			return
		} else {
			err := service.ChangePost(postId, newPost, UpdateTime)
			if err != nil {
				fmt.Println("change password err: ", err)
				tool.RespInternalError(ctx)
				return
			}

			tool.RespSuccessful(ctx)
		}
	}
	tool.RespErrorWithDate(ctx, "无法更改他人留言")
}

func uploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(500, "上传文章出错")
	}
	err = ctx.SaveUploadedFile(file, file.Filename)
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, file.Filename)
}

func postLikes(ctx *gin.Context) {
	postIdString := ctx.Param("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "post_id格式有误")
		return
	}
	err = service.PostLikes(postId)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}

	tool.RespSuccessful(ctx)
}

func deletePost0(ctx *gin.Context) {
	postIdString := ctx.Param("post_id")
	postId, err := strconv.Atoi(postIdString)

	if err != nil {
		fmt.Println("post id string to int err: ", err)
		tool.RespErrorWithDate(ctx, "post_id格式有误")
		return
	}
	err = service.DeletePost(postId)
	if err != nil {
		fmt.Println("delete post err: ", err)
		tool.RespInternalError(ctx)
		return
	}
	tool.RespSuccessful(ctx)
}
