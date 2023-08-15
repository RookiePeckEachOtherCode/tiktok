package service

import (
	"encoding/json"
	"fmt"
	"log"
	"tiktok/dao"
	"tiktok/middleware/redis"
	"time"
)

func PostMessage(userId, toUserId int64, content string) error {
	//TODO 敏感词过滤

	if err := PushRedis(userId, toUserId, content); err != nil {
		log.Println("保存到redis失败 : ", err)
		return fmt.Errorf("保存到redis失败 : %w", err)
	}
	if err := dao.NewMessage(userId, toUserId, content); err != nil {
		log.Println("保存到mysql失败: ", err)
		return fmt.Errorf("保存到mysql失败: %w", err)
	}
	return nil
}
func PushRedis(userId int64, toUserId int64, content string) error {
	timeUnix := time.Now().Unix()
	message := dao.ChatRecord{
		FromUserId:  userId,
		ToUserId:    toUserId,
		Content:     content,
		CreatedTime: timeUnix,
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
