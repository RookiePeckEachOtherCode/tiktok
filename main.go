package main

import (
	"fmt"
	"tiktok/configs"
	"tiktok/dao"
	"tiktok/router"
)

func main() {
	configs.Bless()
	dao.InitDb()
	InitGin()

}

// 初始化数据库和路由
func InitGin() {
	r := router.Init()

	r.Run(fmt.Sprintf(":%d", configs.GIN_PORT))
}
