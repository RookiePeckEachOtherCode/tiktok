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

func (r *Redis) GetUserFavoriteCount(uid int64) int64 {
	key := fmt.Sprintf("%s:%d", "favorite", uid)
	count := rdb.SCard(ctx, key).Val()
	return count
}

func (r *Redis) UpdateUserReceivedLikeCount(uid int64, state bool) {
	key := fmt.Sprintf("liked:%d", uid)
	if state {
		rdb.IncrBy(ctx, key, 1)
	} else {
		rdb.DecrBy(ctx, key, 1)
	}
}

func (r *Redis) GetUserReceivedLikeCount(uid int64) int64 {
	key := fmt.Sprintf("liked:%d", uid)
	count, _ := rdb.Get(ctx, key).Int64()
	return count
}
