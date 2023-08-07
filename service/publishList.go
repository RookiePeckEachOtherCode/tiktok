package service

import (
	"errors"
	"fmt"
	"tiktok/dao"
	"tiktok/middleware/redis"
)

func GetPublishList(userId int64) (*[]dao.Video, error) {
	var videos []dao.Video
	var userInfo dao.UserInfo

	if err := Check(userId, &videos, &userInfo); err != nil {
		return nil, fmt.Errorf("Check: %v", err)
	}

	for i := range videos {
		videos[i].Author = userInfo
		videos[i].IsFavorite = redis.NewProxyIndexMap().GetFavorateState(userId, videos[i].ID)
	}

	return &videos, nil
}

func Check(userId int64, videos *[]dao.Video, userInfo *dao.UserInfo) error {
	if err := CheckIsExistByID(userId); err != nil {
		return err
	}
	if err := GetVideoListByUserId(userId, videos); err != nil {
		return err
	}
	if err := GetUserInfoByUserId(userId, userInfo); err != nil {
		return err
	}
	return nil
}

func CheckIsExistByID(userId int64) error {
	if !dao.CheckIsExistByID(userId) {
		return errors.New("该用户不存在")
	}
	return nil
}

func GetVideoListByUserId(userId int64, videos *[]dao.Video) error {
	var err error
	videos, err = dao.GetVideoListByUserId(userId)
	return err
}
func GetUserInfoByUserId(userId int64, userInfo *dao.UserInfo) error {
	var err error
	userInfo, err = dao.GetUserInfoById(userId)
	return err
}
