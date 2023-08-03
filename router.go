package main

import "github.com/gin-gonic/gin"

func initRouter(r *gin.Engine) {
	apiGroup := r.Group("/douyin")

	apiGroup.GET("/feed")
}
