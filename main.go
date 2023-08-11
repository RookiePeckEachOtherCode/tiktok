package main

import (
	"fmt"
	"tiktok/configs"
	"tiktok/dao"
	"tiktok/router"
)

func main() {
	// 添加东方神秘力量
	configs.Bless()
	// 初始化数据库
	dao.InitDb()
	// 初始化路由
	InitGin()

}

// 初始化数据库和路由
func InitGin() {
	r := router.Init()

	r.Run(fmt.Sprintf(":%d", configs.GIN_PORT))
}
