package main

import (
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

	initRouter(r)

	r.Run()
}
