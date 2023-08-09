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

	//注册 feed 路由
	apiGroup.GET("/feed", controller.Feed)
	//注册 login 路由
	apiGroup.POST("/user/login/", hash.CheckUserName(), hash.CheckAndHashPassword(), controller.UserLogin)
	//注册 register 路由
	apiGroup.POST("/user/register/", hash.CheckUserName(), hash.CheckAndHashPassword(), controller.UserRegister)
	//注册 user-info 路由
	apiGroup.GET("/user", jwt.Auth(), controller.GetUserInfo)
	//注册 publish-action路由
	apiGroup.POST("/publish/action/", jwt.Auth(), controller.PublishVideo)
	//注册 get-publish-list路由
	apiGroup.GET("/publish/list/", jwt.Auth(), controller.PublishList)
	//注册 FavoriteAct路由
	apiGroup.POST("/favorite/action/", jwt.Auth(), controller.FavoriteAct)
	//注册 FavList路由
	apiGroup.GET("/favorite/list/", jwt.Auth(), controller.RecFavList)
	return r
}
