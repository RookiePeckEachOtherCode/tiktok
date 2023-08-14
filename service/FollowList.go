package service

import (
	"tiktok/dao"
	"tiktok/model"
)

type FollowListres struct {
	model.Response
	UserList []*dao.UserInfo `json:"user_list"`
}

func HandleFollowList(uid int64) (*FollowListres, error) {
	list, err := dao.GetFloList(uid)
	if err != nil {
		return nil, err
	}
	res := &FollowListres{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserList: list,
	}
	return res, nil
}
