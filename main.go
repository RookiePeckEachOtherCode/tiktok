package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"tiktok/configs"
	"tiktok/dao"
	"tiktok/middleware/redis"
	"tiktok/router"
	tiktokLog "tiktok/util/log"

	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 添加东方神秘力量
	configs.Bless()
	// 初始化数据库
	dao.InitDb()
	// 初始化redis
	redis.Init()
	// 初始化路由
	InitGin()

	tiktokLog.Normal("服务启动成功")
}

func InitGin() {
	r := router.Init()

	err := r.Run(fmt.Sprintf(":%d", configs.GIN_PORT))
	if err != nil {
		log.Panicln(err)
	}
}
