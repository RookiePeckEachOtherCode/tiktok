package redis

import (
	"context"
	"fmt"
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

// ProxyIndexMap结构体用于封装redis操作
type ProxyIndexMap struct{}

// NewProxyIndexMap函数用于创建ProxyIndexMap实例
func NewProxyIndexMap() *ProxyIndexMap {
	return &ProxyIndexMap{}
}

// GetUserRelation函数用于获取用户关系
func (p *ProxyIndexMap) GetUserRelation(userId, followId int64) bool {
	key := fmt.Sprintf("%v:%v", userId, followId)
	result := rdb.SIsMember(ctx, key, followId)
	return result.Val()
}

// GetFavorateState函数用于获取用户收藏状态
func (p *ProxyIndexMap) GetFavorateState(userId, videoId int64) bool {
	key := fmt.Sprintf("%v:%v", "favor", userId)
	result := rdb.SIsMember(ctx, key, videoId)
	return result.Val()
}
