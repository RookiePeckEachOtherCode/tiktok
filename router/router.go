package router

import (
	"tiktok/controller"
	"tiktok/middleware/hash"

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
	apiGroup.POST("/publish/action/", controller.PublishVideo)

	return r
}
