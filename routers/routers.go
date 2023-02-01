package routers

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine {
	server := gin.Default()
	server.Static("static", "./static")

	// 配置路由组
	userGroup := server.Group("/douyin/user")
	{
		userGroup.POST("/register") // 用户注册接口
		userGroup.POST("/login")    // 用户登陆接口
		userGroup.GET("/")          // 请求获取用户信息接口
	}
	publishGroup := server.Group("/douyin/publish")
	{
		publishGroup.POST("/action") // 投稿接口
		publishGroup.GET("list")     // 请求获取发布列表
	}
	favoriteGroup := server.Group("/douyin/favorite")
	{
		favoriteGroup.POST("/action") // 赞操作
		favoriteGroup.GET("/list")    // 请求喜欢列表
	}
	commentGroup := server.Group("/douyin/comment")
	{
		commentGroup.POST("/action") // 评论操作
		commentGroup.GET("/list")    // 请求评论列表
	}
	relationGroup := server.Group("/douyin/relation")
	{
		relationGroup.POST("/action")       // 关注操作
		relationGroup.GET("/follow/list")   // 请求关注列表
		relationGroup.GET("/follower/list") // 请求粉丝列表
		relationGroup.GET("/friend/list")   // 请求好友列表
	}
	messageGroup := server.Group("/douyin/message")
	{
		messageGroup.POST("/action") // 发送消息
		messageGroup.GET("/chat")    // 请求聊天记录
	}
	return server
}
