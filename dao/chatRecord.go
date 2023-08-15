package dao

import (
	"encoding/json"
	"fmt"
	"tiktok/middleware/redis"
	"time"
)

type ChatRecord struct {
	ID        int64  `json:"id" gorm:"primaryKey;column:id"`
	Content   string `json:"content" gorm:"column:content"`
	UserId    int64  `json:"from_user_id" gorm:"column:user_id"`
	ToUserId  int64  `json:"to_user_id " gorm:"column:to_user_id"`
	CreatedAt int64  `json:"created_at"  gorm:"column:created_at"`
}

func GetChatRecordList(userId, ToUserId int64) ([]ChatRecord, error) {
	var messageList []ChatRecord
	if err := DB.Where("user_id = ? AND to_user_id = ? ", userId, ToUserId).
		Or("user_id = ? AND to_user_id = ?", ToUserId, userId).
		Order("created_at asc").Find(&messageList).Error; err != nil {
		return nil, err
	}
	return messageList, nil
}

func NewMessage(userId, toUserId int64, content string) (int64, error) {
	message := &ChatRecord{
		UserId:    userId,
		ToUserId:  toUserId,
		Content:   content,
		CreatedAt: time.Now().UnixMilli(),
	}
	if err := DB.Create(message).Error; err != nil {
		return 0, err
	}

	return message.CreatedAt, nil
}

func AddMessageListInRedis(userId, toUserId int64, message []ChatRecord) error {
	msgName := fmt.Sprintf("%d-%d", userId, toUserId)
	for _, message := range message {
		bytes, _ := json.Marshal(message)
		err := redis.New(redis.MSGS).AddAllMessage(msgName, bytes, message.CreatedAt)
		if err != nil {
			return err
		}
	}
	return nil
}

func ParesMessageListFromRedis(uerId, toUserId, msgTime int64) ([]ChatRecord, error) {
	var messageList []ChatRecord
	msgName := fmt.Sprintf("%d-%d", uerId, toUserId)
	for {
		bytes, err := redis.New(redis.MSGS).GetMessage(msgName)
		if err != nil {
			return nil, err
		}
		if bytes == "" {
			break
		}
		message := ChatRecord{}
		json.Unmarshal([]byte(bytes), &message)
		if message.CreatedAt >= msgTime {
			continue
		}
		messageList = append(messageList, message)
	}
	return messageList, nil
}
