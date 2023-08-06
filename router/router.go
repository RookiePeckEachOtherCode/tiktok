package router

import (
	"tiktok/controller"

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
	apiGroup.POST("/user/login/", controller.UserLogin)
	//注册 register 路由
	apiGroup.POST("/user/register/", controller.UserRegister)
	apiGroup.POST("/publish/action/", controller.PublishVideo)

	return r
}
