package routers

import (
	"douyin/handlers"
	"douyin/middlewares"
	"douyin/models"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	models.InitDB()         // 初始化gorm
	server := gin.Default() // 初始化gin服务器
	//server.Use(middlewares.PasswordEncryptionMiddleware()) // 注册加密中间件
	//server.Use(middlewares.JWTMiddleware())                // 注册JWT鉴权中间件
	server.Static("static", "./static")

	// 配置路由组
	userGroup := server.Group("/douyin/user")
	{
		userGroup.POST("/register/", middlewares.PasswordEncryptionMiddleware(), handlers.UserRegisterHandle) // 用户注册接口
		userGroup.POST("/login/", middlewares.PasswordEncryptionMiddleware(), handlers.UserLoginHandle)       // 用户登陆接口
		userGroup.GET("/", middlewares.JWTMiddleware(), handlers.GetUserInfoHandle)                           // 请求获取用户信息接口
	}
	publishGroup := server.Group("/douyin/publish")
	{
		publishGroup.POST("/action/", middlewares.JWTMiddleware(), handlers.PublishVideoHandle) // 投稿接口
		publishGroup.GET("/list/", middlewares.JWTMiddleware(), handlers.GetPublishListHandle)  // 请求获取发布列表
	}
	favoriteGroup := server.Group("/douyin/favorite")
	{
		favoriteGroup.POST("/action/", middlewares.JWTMiddleware(), handlers.FavoriteActionHandle) // 赞操作
		favoriteGroup.GET("/list/", middlewares.JWTMiddleware(), handlers.GetFavoriteListHandle)   // 请求喜欢列表
	}
	commentGroup := server.Group("/douyin/comment")
	{
		commentGroup.POST("/action/", middlewares.JWTMiddleware(), handlers.CommentActionHandle) // 评论操作
		commentGroup.GET("/list/", middlewares.JWTMiddleware(), handlers.GetCommentListHandle)   // 请求评论列表
	}
	relationGroup := server.Group("/douyin/relation")
	{
		relationGroup.POST("/action/", middlewares.JWTMiddleware(), handlers.RelationActionHandle)        // 关注操作
		relationGroup.GET("/follow/list/", middlewares.JWTMiddleware(), handlers.GetFollowListHandle)     // 请求关注列表
		relationGroup.GET("/follower/list/", middlewares.JWTMiddleware(), handlers.GetFollowerListHandle) // 请求粉丝列表
		relationGroup.GET("/friend/list/", middlewares.JWTMiddleware(), handlers.GetFriendListHandle)     // 请求好友列表
	}
	messageGroup := server.Group("/douyin/message")
	{
		messageGroup.POST("/action/", middlewares.JWTMiddleware(), handlers.MessageActionHandle) // 发送消息
		messageGroup.GET("/chat/", middlewares.JWTMiddleware(), handlers.MessageChatHandle)      // 请求聊天记录
	}
	return server
}
