package dao

import (
	"encoding/json"
	"fmt"
	"tiktok/middleware/redis"
	"time"
)

type ChatRecord struct {
	ID          int64  `json:"id" gorm:"primaryKey;column:id"`
	ToUserId    int64  `json:"to_user_id" gorm:"column:to_user_id"`
	FromUserId  int64  `json:"from_user_id" gorm:"column:user_id"`
	Content     string `json:"content" gorm:"column:content"`
	CreatedTime int64  `json:"create_time" gorm:"created_time"`
}

func GetChatRecordList(userId, ToUserId int64) ([]ChatRecord, error) {
	tx := DB.Begin()
	var messageList []ChatRecord
	if err := tx.Where("user_id = ? AND to_user_id = ? ", userId, ToUserId).
		Or("user_id = ? AND to_user_id = ?", ToUserId, userId).
		Order("created_time asc").Find(&messageList).Error; err != nil {

		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return messageList, nil
}

func NewMessage(userId, toUserId int64, content string) error {
	message := &ChatRecord{
		FromUserId:  userId,
		ToUserId:    toUserId,
		Content:     content,
		CreatedTime: time.Now().Unix(),
	}
	if err := DB.Create(message).Error; err != nil {
		return err
	}

	return nil
}

func AddMessageListInRedis(userId, toUserId int64, message []ChatRecord) {
	msgName := fmt.Sprintf("%d-%d", userId, toUserId)
	for _, message := range message {
		bytes, _ := json.Marshal(message)
		redis.New(redis.MSGS).AddAllMessage(msgName, bytes, message.CreatedTime)
	}
}

func ParesMessageListFromRedis(uerId, toUserId, msgTime int64) ([]ChatRecord, error) {
	var messageList []ChatRecord
	msgName := fmt.Sprintf("%d-%d", uerId, toUserId)
	for {
		bytes, _ := redis.New(redis.MSGS).GetMessage(msgName)
		if bytes == "" {
			break
		}
		message := ChatRecord{}
		json.Unmarshal([]byte(bytes), &message)
		// 如果消息时间大于等于传入的时间，说明是新消息，直接跳过
		if message.CreatedTime >= msgTime {
			continue
		}
		messageList = append(messageList, message)
	}
	return messageList, nil
}
