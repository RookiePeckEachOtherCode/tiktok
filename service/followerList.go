package service

import (
	"errors"
	"tiktok/dao"
)

func GetFollowerList(userId int64) ([]*dao.UserInfo, error) {
	if !dao.CheckIsExistByID(userId) {
		return nil, errors.New("该用户不存在")
	}

	userList, err := dao.GetFollowerListById(userId)

	if err != nil {
		return nil, err
	}

	for _, v := range userList {
		v.IsFollow = dao.GetUserRelation(userId, v.ID)
	}

	return userList, nil
}