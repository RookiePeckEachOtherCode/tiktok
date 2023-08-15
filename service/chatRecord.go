package service

import (
	"fmt"
	"tiktok/dao"
)

func GetChatRecord(userId, toUserId int64, preMsgTime int64) ([]dao.ChatRecord, error) {
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

	messagelist, err := dao.ParesMessageListFromRedis(userId, toUserId, preMsgTime)

	if err != nil {
		return nil, err
	}
	return messagelist, nil
}
