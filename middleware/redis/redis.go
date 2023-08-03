package redis

import (
	"context"
	"fmt"
	"tiktok/configs"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func redisInit() {
	rdb = redis.NewClient(
		&redis.Options{
			Addr:     configs.GetDBInfo(),
			Password: "",
			DB:       0,
		},
	)
}

type ProxyIndexMap struct{}

func NewProxyIndexMap() *ProxyIndexMap {
	return &ProxyIndexMap{}
}

func (p *ProxyIndexMap) GetUserRelation(userId, followId int64) bool {
	key := fmt.Sprintf("%v:%v", userId, followId)
	result := rdb.SIsMember(ctx, key, followId)
	return result.Val()
}

func (p *ProxyIndexMap) GetFavorateState(userId, videoId int64) bool {
	key := fmt.Sprintf("%v:%v", "favor", userId)
	result := rdb.SIsMember(ctx, key, videoId)
	return result.Val()
}
