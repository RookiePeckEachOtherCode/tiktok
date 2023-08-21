package service

import (
	"encoding/json"
	"fmt"
	"log"
	"tiktok/dao"
	"tiktok/middleware/redis"
	"time"
)

func ChatActionService(userId, toUserId int64, content string) error {

	if err := pushToRedis(userId, toUserId, content); err != nil {
		log.Println("保存到redis失败 : ", err)
		return fmt.Errorf("保存到redis失败 : %w", err)
	}
	if err := dao.NewMessage(userId, toUserId, content); err != nil {
		log.Println("保存到mysql失败: ", err)
		return fmt.Errorf("保存到mysql失败: %w", err)
	}
	return nil
}

func ChatRecordService(userId, toUserId int64, preMsgTime int64) ([]dao.ChatRecord, error) {
	if userId == toUserId {
		return nil, fmt.Errorf("不能和自己聊天")
	}

	if preMsgTime == 0 {
		messageList, err := dao.GetChatRecordList(userId, toUserId)
		if err != nil {
			return nil, fmt.Errorf("获取聊天记录失败: %w", err)
		}
		dao.AddMessageListInRedis(userId, toUserId, messageList)
		return messageList, nil
	}

	messageList, err := dao.ParesMessageListFromRedis(userId, toUserId, preMsgTime)

	if err != nil {
		return nil, err
	}
	return messageList, nil
}

func FriendListService(userId int64) ([]*dao.Friend, error) {
	var FriendList []*dao.Friend
	eachLikeUserInfo, err := dao.GetMutualFriendListById(userId)
	if err != nil {
		return nil, err
	}

	for _, userInfo := range eachLikeUserInfo {
		message, msgType, err := dao.GetNewestMessageByUserIdAndToUserID(userId, userInfo.ID)

		if err != nil {
			log.Println("获取最新消息失败", err)
			return nil, err
		}

		FriendList = append(FriendList, &dao.Friend{
			UserInfo: *userInfo,
			Message:  message,
			MsgType:  msgType,
		})
	}

	return FriendList, nil
}

func pushToRedis(userId int64, toUserId int64, content string) error {
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
