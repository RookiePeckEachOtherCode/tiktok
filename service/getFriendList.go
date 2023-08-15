package service

import (
	"log"
	"tiktok/dao"
)

func GetFriendList(userId int64) ([]*dao.Friend, error) {
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
			UserInfo: userInfo,
			Message:  message,
			MsgType:  msgType,
		})
	}

	return FriendList, nil
}
