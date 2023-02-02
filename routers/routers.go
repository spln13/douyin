package routers

import (
	"douyin/dao"
	"douyin/handlers"
	"douyin/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	dao.InitDB()                                           // 初始化gorm
	server := gin.Default()                                // 初始化gin服务器
	server.Use(middlewares.PasswordEncryptionMiddleware()) // 注册加密中间件
	server.Static("static", "./static")

	// 配置路由组
	userGroup := server.Group("/douyin/user")
	{
		userGroup.POST("/register", middlewares.PasswordEncryptionMiddleware(), handlers.UserRegisterHandle) // 用户注册接口
		userGroup.POST("/login", middlewares.PasswordEncryptionMiddleware(), handlers.UserLoginHandle)       // 用户登陆接口
		userGroup.GET("/", handlers.GetUserInfoHandle)                                                       // 请求获取用户信息接口
	}
	publishGroup := server.Group("/douyin/publish")
	{
		publishGroup.POST("/action", handlers.PublishActionHandle) // 投稿接口
		publishGroup.GET("/list", handlers.GetPublishListHandle)   // 请求获取发布列表
	}
	favoriteGroup := server.Group("/douyin/favorite")
	{
		favoriteGroup.POST("/action", handlers.FavoriteActionHandle) // 赞操作
		favoriteGroup.GET("/list", handlers.GetFavoriteListHandle)   // 请求喜欢列表
	}
	commentGroup := server.Group("/douyin/comment")
	{
		commentGroup.POST("/action", handlers.CommentActionHandle) // 评论操作
		commentGroup.GET("/list", handlers.GetCommentListHandle)   // 请求评论列表
	}
	relationGroup := server.Group("/douyin/relation")
	{
		relationGroup.POST("/action", handlers.RelationActionHandle)        // 关注操作
		relationGroup.GET("/follow/list", handlers.GetFollowListHandle)     // 请求关注列表
		relationGroup.GET("/follower/list", handlers.GetFollowerListHandle) // 请求粉丝列表
		relationGroup.GET("/friend/list", handlers.GetFriendListHandle)     // 请求好友列表
	}
	messageGroup := server.Group("/douyin/message")
	{
		messageGroup.POST("/action", handlers.MessageActionHandle) // 发送消息
		messageGroup.GET("/chat", handlers.MessageChatHandle)      // 请求聊天记录
	}
	return server
}
