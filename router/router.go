package router

import (
	"io"
	"os"
	"tiktok/controller"
	"tiktok/middleware/hash"
	"tiktok/middleware/jwt"

	"github.com/gin-gonic/gin"
)

// 初始化路由

func Init() *gin.Engine {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 记录到文件。
	f, _ := os.Create("tiktok.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	r.Static("static", "./static")

	apiGroup := r.Group("/douyin")

	//注册 视频流 路由
	apiGroup.GET("/feed", controller.VideoFeedController)
	//注册 登陆 路由
	apiGroup.POST("/user/login/", hash.CheckUserName(), hash.CheckAndHashPassword(), controller.UserLoginController)
	//注册 注册 路由
	apiGroup.POST("/user/register/", hash.CheckUserName(), hash.CheckAndHashPassword(), controller.UserRegisterController)
	//注册 获取用户信息 路由
	apiGroup.GET("/user", jwt.Auth(), controller.UserInfoController)
	//注册 发布视频 路由
	apiGroup.POST("/publish/action/", jwt.Auth(), controller.PublishVideoController)
	//注册 获取发布列表 路由
	apiGroup.GET("/publish/list/", jwt.Auth(), controller.PublishListController)
	//注册 赞操作 路由
	apiGroup.POST("/favorite/action/", jwt.Auth(), controller.FavoriteActionController)
	//注册 获取喜欢列表 路由
	apiGroup.GET("/favorite/list/", jwt.Auth(), controller.FavoriteListController)
	//注册 评论 路由
	apiGroup.POST("/comment/action/", jwt.Auth(), controller.CommentActionController)
	//注册 获取评论列表 路由
	apiGroup.GET("/comment/list/", jwt.Auth(), controller.CommentListController)
	//注册 获取关注者列表 路由
	apiGroup.GET("/relation/follower/list", jwt.Auth(), controller.FollowerListController)
	//注册 关注操作 路由
	apiGroup.POST("/relation/action/", jwt.Auth(), controller.FollowActionController)
	//注册 关注的人列表 路由
	apiGroup.GET("/relation/follow/list/", jwt.Auth(), controller.FollowListController)
	//注册发送消息路由
	apiGroup.POST("/message/action/", jwt.Auth(), jwt.FilterDirtyMessage(), controller.ChatActionController)
	// //注册消息列表路由
	apiGroup.GET("/message/chat/", jwt.Auth(), controller.ChatRecordListController)
	//注册好友列表路由
	apiGroup.GET("/relation/friend/list/", jwt.Auth(), controller.FriendListController)

	return r
}
