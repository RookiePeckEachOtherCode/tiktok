package redis

import (
	"context"
	"fmt"
	"strings"
	"tiktok/configs"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

// redisInit函数用于初始化redis客户端
func Init() {
	rdb = redis.NewClient(
		&redis.Options{
			Addr:     configs.GetDBInfo(),
			Password: "",
			DB:       0,
		},
	)
}

// GetUserRelation函数用于获取用户关系
func GetUserRelation(userId, followId int64) bool {
	key := fmt.Sprintf("%v:%v", "relation", followId)
	result := rdb.SIsMember(ctx, key, followId)
	PrintLog(fmt.Sprintf("调用了redis的获取用户关系函数，userId:%d,followId:%d,result:%v", userId, followId, result.Val()))
	return result.Val()
}

// SetUserRelation函数用于设置点赞关系
func SetFavorateState(userId, videoId int64, state bool) {
	key := fmt.Sprintf("%v:%v", "favor", userId)
	if state {
		rdb.SAdd(ctx, key, videoId)
		return
	}
	rdb.SRem(ctx, key, videoId)
}

// GetFavorateState函数用于获取用户收藏状态
func GetFavorateState(userId, videoId int64) bool {
	key := fmt.Sprintf("%v:%v", "favor", userId)
	result := rdb.SIsMember(ctx, key, videoId)
	PrintLog(fmt.Sprintf("调用了redis的获取收藏状态函数，userId:%d,videoId:%d,result:%v", userId, videoId, result.Val()))
	return result.Val()
}

func PrintLog(log string) {
	// 计算日志信息的长度
	length := len(log) + 4

	// 打印上边框
	fmt.Println(strings.Repeat("+", length))

	// 打印日志信息
	fmt.Printf("| %s |\n", log)

	// 打印下边框
	fmt.Println(strings.Repeat("+", length))
}
