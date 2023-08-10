package router

import (
	"tiktok/controller"
	"tiktok/middleware/hash"
	"tiktok/middleware/jwt"

	"github.com/gin-gonic/gin"
)

// 初始化路由

func Init() *gin.Engine {
	r := gin.Default()

	r.Static("static", "./static")

	apiGroup := r.Group("/douyin")

	//注册 视频流 路由
	apiGroup.GET("/feed", controller.Feed)
	//注册 登陆 路由
	apiGroup.POST("/user/login/", hash.CheckUserName(), hash.CheckAndHashPassword(), controller.UserLogin)
	//注册 注册 路由
	apiGroup.POST("/user/register/", hash.CheckUserName(), hash.CheckAndHashPassword(), controller.UserRegister)
	//注册 获取用户信息 路由
	apiGroup.GET("/user", jwt.Auth(), controller.GetUserInfo)
	//注册 发布视频 路由
	apiGroup.POST("/publish/action/", jwt.Auth(), controller.PublishVideo)
	//注册 获取发布列表 路由
	apiGroup.GET("/publish/list/", jwt.Auth(), controller.PublishList)
	//注册 赞操作 路由
	apiGroup.POST("/favorite/action/", jwt.Auth(), controller.FavoriteAct)
	//注册 获取喜欢列表 路由
	apiGroup.GET("/favorite/list/", jwt.Auth(), controller.RecFavList)
	//注册 评论 路由
	apiGroup.POST("/comment/action/", jwt.Auth(), controller.CommentAct)
	//注册 获取评论列表 路由
	apiGroup.GET("/comment/list/", jwt.Auth(), controller.RecComList)
	//注册 获取关注列表 路由
	apiGroup.GET("/relation/follower/list", jwt.Auth(), controller.GetFollowerList)

	return r
}
