package main

import (
	"fmt"
	"log"
	"tiktok/configs"
	"tiktok/dao"
	"tiktok/middleware/redis"
	"tiktok/router"
)

func main() {
	// 添加东方神秘力量
	configs.Bless()
	// 初始化数据库
	dao.InitDb()
	// 初始化redis
	redis.Init()
	// 初始化路由
	InitGin()

}

func InitGin() {
	r := router.Init()

	err := r.Run(fmt.Sprintf(":%d", configs.GIN_PORT))
	if err != nil {
		log.Panicln(err)
	}
}
