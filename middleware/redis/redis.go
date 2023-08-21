package redis

import (
	"context"
	"fmt"
	"log"
	"tiktok/configs"

	"github.com/go-redis/redis/v8"
)

const (
	FAVORITE = 1
	RELATION = 2
	LIKED    = 3
	MSGS     = 4
)

var ctx = context.Background()
var fav *redis.Client
var relation *redis.Client
var liked *redis.Client
var msgs *redis.Client

func Init() {
	log.Println("REDIS INIT")
	fav = redis.NewClient( //管理点赞的数据
		&redis.Options{
			Addr:     configs.GetRedisInfo(),
			Password: "",
			DB:       0,
		})
	relation = redis.NewClient( //管理关注的数据
		&redis.Options{
			Addr:     configs.GetRedisInfo(),
			Password: "",
			DB:       1,
		})
	liked = redis.NewClient( //管理收藏的数据
		&redis.Options{
			Addr:     configs.GetRedisInfo(),
			Password: "",
			DB:       2,
		})
	msgs = redis.NewClient( //管理聊天的数据
		&redis.Options{
			Addr:     configs.GetRedisInfo(),
			Password: "",
			DB:       3,
		})

	if _, err := fav.Ping(ctx).Result(); err != nil {
		log.Panicln("redis_fav init error")
	}
	if _, err := relation.Ping(ctx).Result(); err != nil {
		log.Panicln("redis_relation init error")
	}
	if _, err := liked.Ping(ctx).Result(); err != nil {
		log.Panicln("redis_liked init error")
	}
	if _, err := msgs.Ping(ctx).Result(); err != nil {
		log.Panicln("redis_message init error")
	}
}

type Redis struct {
	*redis.Client
}

func New(num int) *Redis {
	switch num {
	case FAVORITE:
		return &Redis{fav}
	case RELATION:
		return &Redis{relation}
	case LIKED:
		return &Redis{liked}
	case MSGS:
		return &Redis{msgs}
	default:
		log.Panicln("redis num error")
		return nil
	}
}

// Favorite
// ==================================================================================================
func (r *Redis) UpdateFavoriteState(uid, vid int64, state bool) {
	key := fmt.Sprintf("%s:%d", "favorite", uid)
	if state {
		r.Client.SAdd(ctx, key, vid)
		return
	}
	r.Client.SRem(ctx, key, vid)
}

func (r *Redis) GetFavoriteState(uid int64, vid int64) bool {
	key := fmt.Sprintf("%s:%d", "favorite", uid)
	ret := r.Client.SIsMember(ctx, key, vid)
	return ret.Val()
}

func (r *Redis) GetUserFavoriteCount(uid int64) int64 {
	key := fmt.Sprintf("%s:%d", "favorite", uid)
	count := r.Client.SCard(ctx, key).Val()
	return count
}

// Relation
// ==================================================================================================
func (r *Redis) UpdateUserRelation(userId int64, followId int64, state bool) {
	key := fmt.Sprintf("%s:%d", "relation", userId)
	if state {
		r.Client.SAdd(ctx, key, followId)
		return
	}
	r.Client.SRem(ctx, key, followId)
}

func (r *Redis) GetUserRelation(userId int64, tid int64) bool {
	key := fmt.Sprintf("%s:%d", "relation", userId)
	ret := r.Client.SIsMember(ctx, key, tid)
	return ret.Val()
}

// Liked
// ==================================================================================================
func (r *Redis) UpdateUserReceivedLikeCount(uid int64, state bool) {
	key := fmt.Sprintf("liked:%d", uid)
	if state {
		r.Client.IncrBy(ctx, key, 1)
	} else {
		r.Client.DecrBy(ctx, key, 1)
	}
}

func (r *Redis) GetUserReceivedLikeCount(uid int64) int64 {
	key := fmt.Sprintf("liked:%d", uid)
	count, _ := r.Client.Get(ctx, key).Int64()
	return count
}

// Msgs
// ==================================================================================================
func (r *Redis) NewMessage(msgName string, bytes []byte) {
	r.Client.RPush(ctx, msgName, bytes)
}

func (r *Redis) AddAllMessage(msgName string, bytes []byte, createTime int64) {
	r.Client.ZAdd(ctx, msgName, &redis.Z{
		Score:  float64(createTime),
		Member: bytes,
	})
}

func (r *Redis) GetMessage(msgName string) (string, error) {
	return r.Client.LPop(ctx, msgName).Result()
}

// flushall

func RedisFlushAll() {
	fav.FlushAll(ctx)
	relation.FlushAll(ctx)
	liked.FlushAll(ctx)
	msgs.FlushAll(ctx)
}
