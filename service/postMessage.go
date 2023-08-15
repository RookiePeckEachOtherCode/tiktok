package service

import (
	"encoding/json"
	"fmt"
	"log"
	"tiktok/dao"
	"tiktok/middleware/redis"
)

func PostMessage(userId, toUserId int64, content string) error {
	var time int64
	var err error

	if time, err = dao.NewMessage(userId, toUserId, content); err != nil {
		log.Println("保存到mysql失败: ", err)
		return fmt.Errorf("保存到mysql失败: %w", err)
	}
	// 将消息添加到redis
	if err := PushRedis(userId, toUserId, content, time); err != nil {
		log.Println("保存到redis失败 : ", err)
		return fmt.Errorf("保存到redis失败 : %w", err)
	}
	return nil
}
func PushRedis(userId int64, toUserId int64, content string, time int64) error {
	message := dao.ChatRecord{
		UserId:    userId,
		ToUserId:  toUserId,
		Content:   content,
		CreatedAt: time,
	}
	msgName := fmt.Sprintf("%d-%d", userId, toUserId)

	// 序列化
	bytes, err := json.Marshal(message)
	if err != nil {
		log.Println("序列化失败: ", err)
		return err
	}
	redis.New(redis.MSGS).NewMessage(msgName, bytes)
	return nil
}
