package main

import (
	"tiktok/controller"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter(r *gin.Engine) {
	apiGroup := r.Group("/douyin")

	// 注册 feed 路由
	apiGroup.GET("/feed", controller.Feed)
}
