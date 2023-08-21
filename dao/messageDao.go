package dao

import (
	"encoding/json"
	"fmt"
	"tiktok/middleware/redis"
	tiktokLog "tiktok/util/log"
	"time"
)

type ChatRecord struct {
	ID          int64  `json:"id" gorm:"primaryKey;column:id"`
	ToUserId    int64  `json:"to_user_id" gorm:"column:to_user_id"`
	FromUserId  int64  `json:"from_user_id" gorm:"column:user_id"`
	Content     string `json:"content" gorm:"column:content"`
	CreatedTime int64  `json:"create_time" gorm:"created_time"`
}

// GetChatRecordList 获取聊天记
func GetChatRecordList(userId, ToUserId int64) ([]ChatRecord, error) {
	tx := DB.Begin()
	var messageList []ChatRecord
	if err := tx.Where("user_id = ? AND to_user_id = ? ", userId, ToUserId).
		Or("user_id = ? AND to_user_id = ?", ToUserId, userId).
		Order("created_time asc").Find(&messageList).Error; err != nil {
		tiktokLog.Error(fmt.Sprintf("获取聊天记录失败, userId: %d, ToUserId: %d, Error: %v", userId, ToUserId, err))
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return messageList, nil
}

// NewMessage 保存消息到数据库
func NewMessage(userId, toUserId int64, content string) error {
	message := &ChatRecord{
		FromUserId:  userId,
		ToUserId:    toUserId,
		Content:     content,
		CreatedTime: time.Now().Unix(),
	}
	if err := DB.Create(message).Error; err != nil {
		tiktokLog.Error(fmt.Sprintf("保存消息到数据库失败, userId: %d, toUserId: %d, content: %s, Error: %v", userId, toUserId, content, err))
		return err
	}

	return nil
}

// AddMessageListInRedis 将消息添加到redis
func AddMessageListInRedis(userId, toUserId int64, message []ChatRecord) {
	msgName := fmt.Sprintf("%d-%d", userId, toUserId)
	for _, message := range message {
		bytes, _ := json.Marshal(message)
		redis.New(redis.MSGS).AddAllMessage(msgName, bytes, message.CreatedTime)
	}
}

// ParesMessageListFromRedis 从redis中解析消息
func ParesMessageListFromRedis(uerId, toUserId, msgTime int64) ([]ChatRecord, error) {
	var messageList []ChatRecord
	msgName := fmt.Sprintf("%d-%d", uerId, toUserId)
	for {
		bytes, _ := redis.New(redis.MSGS).GetMessage(msgName)
		if bytes == "" {
			break
		}
		message := ChatRecord{}
		err := json.Unmarshal([]byte(bytes), &message)
		if err != nil {
			tiktokLog.Error(fmt.Sprintf("解析消息失败, msgName: %s, Error: %v", msgName, err))
			return nil, err
		}
		// 如果消息时间大于等于传入的时间，说明是新消息，直接跳过
		if message.CreatedTime >= msgTime {
			continue
		}
		messageList = append(messageList, message)
	}
	return messageList, nil
}
