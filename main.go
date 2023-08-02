package main

import (
	"tiktok/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	Init()

}
func Init() {
	dao.InitDb()

	r := gin.Default()

	apiGroup := r.Group("/douyin")

	apiGroup.GET("/feed")
	r.Run()
}
