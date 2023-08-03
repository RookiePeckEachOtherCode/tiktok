package main

import (
	"tiktok/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	apiGroup := r.Group("/douyin")

	apiGroup.GET("/feed", controller.Feed)
}
