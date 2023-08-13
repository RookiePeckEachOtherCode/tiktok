package service

import (
	"errors"
	"tiktok/dao"
	"tiktok/middleware/redis"
)

func GetFollowerList(userId int64) ([]*dao.UserInfo, error) {
	if !dao.CheckIsExistByID(userId) {
		return nil, errors.New("该用户不存在")
	}

	userList, err := dao.GetFollowerList(userId)

	if err != nil {
		return nil, err
	}

	for _, v := range userList {
		v.IsFollow = redis.New().GetUserRelation(userId, v.ID)
	}

	return userList, nil
}
