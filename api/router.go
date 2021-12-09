package api

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	engine := gin.Default()

	engine.POST("/register", register) //注册
	engine.POST("/login", login)       //登陆

	userGroup := engine.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/password", changePassword) //修改密码
	}

	postGroup := engine.Group("/post")
	{
		{
			postGroup.Use(auth)
			postGroup.POST("/", addPost)              //发布新留言
			postGroup.POST("/:post_id", changePost)   //修改留言
			postGroup.DELETE("/:post_id", deletePost) //删除留言

			postGroup.GET("/", briefPosts)         //查看全部留言概略
			postGroup.GET("/:post_id", postDetail) //查看一条留言详细信息和其下属评论

			engine.MaxMultipartMemory = 8 << 20
			postGroup.POST("/file", uploadFile) ////上传文章功能

			postGroup.POST("/:post_id/likes", postLikes)
		}
		{
			postGroup.Use(admin)
			postGroup.DELETE("/:post_id0", deletePost0) //管理员删除,无视用户名
		}
	}

	commentGroup := engine.Group("/comment")
	{
		{
			commentGroup.Use(auth)
			commentGroup.POST("/", addComment)                 //发送评论
			commentGroup.DELETE("/:comment_id", deleteComment) //删除评论

			commentGroup.POST("/:comment_id/likes", commentLikes)
		}
		{
			commentGroup.Use(admin)
			commentGroup.DELETE("/:comment_id0", deleteComment0) //管理员删除,无视用户名
		}
	}

	err := engine.Run()
	if err != nil {
		return
	}
}
