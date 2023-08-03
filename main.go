package main

import (
	"tiktok/controller"
	"tiktok/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	Init()

}
func Init() {
	//初始化数据库
	dao.InitDb()

	r := gin.Default()

	apiGroup := r.Group("/douyin")

	apiGroup.GET("/feed", controller.Feed)

	r.Run()
}
