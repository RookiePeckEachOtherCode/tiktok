package main

import (
	"tiktok/controller"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/douyin")

	//注册 feed 路由
	apiGroup.GET("/feed", controller.Feed)
	//注册 login 路由
	apiGroup.POST("/user/login/", controller.Login)

	return r
}
