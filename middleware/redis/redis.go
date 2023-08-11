package redis

import (
	"context"
	"fmt"
	"tiktok/configs"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client
var IsInit = false

func init() {
	rdb = redis.NewClient(
		&redis.Options{
			Addr:     configs.GetRedisInfo(),
			Password: "",
			DB:       0,
		})
	IsInit = true
}

type Redis struct{}

func New() *Redis {
	return &Redis{}
}

func (r *Redis) UpdateFavoriteState(uid, vid int64, state bool) {
	key := fmt.Sprintf("%s:%d", "favorite", uid)
	if state {
		rdb.SAdd(ctx, key, vid)
		return
	}
	rdb.SRem(ctx, key, vid)
}

func (r *Redis) GetFavoriteState(uid int64, vid int64) bool {
	key := fmt.Sprintf("%s:%d", "favorite", uid)
	ret := rdb.SIsMember(ctx, key, vid)
	return ret.Val()
}

func (r *Redis) UpdateUserRelation(userId int64, followId int64, state bool) {
	key := fmt.Sprintf("%s:%d", "relation", userId)
	if state {
		rdb.SAdd(ctx, key, followId)
		return
	}
	rdb.SRem(ctx, key, followId)
}

func (r *Redis) GetUserRelation(userId int64, tid int64) bool {
	key := fmt.Sprintf("%s:%d", "relation", userId)
	ret := rdb.SIsMember(ctx, key, tid)
	return ret.Val()
}
