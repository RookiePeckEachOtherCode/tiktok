package main

import (
	"tiktok/dao"
	"tiktok/router"
)

func main() {
	dao.InitDb()
	InitGin()

}

// 初始化数据库和路由
func InitGin() {
	r := router.Init()

	r.Run(":8080")
}
