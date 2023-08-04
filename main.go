package main

import (
	"tiktok/controller"
	"tiktok/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	Init()
}

// 初始化数据库和路由
func Init() {
	//初始化数据库
	dao.InitDb()

	r := gin.Default()

	apiGroup := r.Group("/douyin")

	// 注册 feed 路由
	apiGroup.GET("/feed", controller.Feed)

	r.Run()
}
