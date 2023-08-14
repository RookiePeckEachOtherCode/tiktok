package service

import (
	"tiktok/dao"
	"tiktok/model"
)

type FriendListRes struct {
	model.Response
	UserList []*dao.UserInfo `json:"user_list"`
}

func HandleFriendList(uid int64) (*FriendListRes, error) {
	floList, err := dao.GetFloList(uid)
	if err != nil {
		return nil, err
	}

	filteredList := FilterFriends(floList, uid)

	res := &FriendListRes{
		UserList: filteredList,
	}
	res.StatusCode = 1
	res.StatusMsg = "获取好友列表成功"
	return res, nil
}
func FilterFriends(floList []*dao.UserInfo, uid int64) []*dao.UserInfo {
	var filteredList []*dao.UserInfo

	for _, friend := range floList {
		if friend.Follwcheck(uid) {
			filteredList = append(filteredList, friend)
		}
	}

	return filteredList
}
