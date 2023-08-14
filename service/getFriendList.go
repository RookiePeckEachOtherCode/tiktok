package service

import "tiktok/dao"

func GetFriendList(userId int64) ([]*dao.UserInfo, error) {
	friendList, err := dao.GetMutualFriendListById(userId)
	if err != nil {
		return nil, err
	}
	return friendList, nil
}
