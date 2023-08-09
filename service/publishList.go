package service

import (
	"errors"
	"fmt"
	"tiktok/dao"
	"tiktok/middleware/redis"
	"tiktok/util"
)

// GetPublishList 获取发布列表
func GetPublishList(userId int64) (*[]dao.Video, error) {
	userInfo, videos, err := PublishListCheck(userId)
	if err != nil {
		return nil, err
	}

	// 为视频列表添加作者信息
	util.PrintLog(fmt.Sprintf("调用了service的获取发布列表函数，userId:%d", userId))
	for i := range *videos {
		(*videos)[i].Author = *userInfo
		(*videos)[i].IsFavorite = redis.GetFavorateState(userId, (*videos)[i].ID)
		//(*videos)[i].IsFavorite = dao.GetIsFavorite(userId, (*videos)[i].ID)
	}

	return videos, nil
}

// PublishListCheck 发布列表检查
func PublishListCheck(userId int64) (*dao.UserInfo, *[]dao.Video, error) {
	if err := CheckIsExistByID(userId); err != nil {
		return nil, nil, err
	}

	userInfo, err := GetUserInfoByUserId(userId)
	if err != nil {
		return nil, nil, err
	}

	videos, err := GetVideoListByUserId(userId)
	if err != nil {
		return nil, nil, err
	}

	return userInfo, videos, nil
}

// CheckIsExistByID 根据用户id检查用户是否存在
func CheckIsExistByID(userId int64) error {
	if !dao.CheckIsExistByID(userId) {
		return errors.New("该用户不存在")
	}
	return nil
}

// GetVideoListByUserId 根据用户id获取视频列表
func GetVideoListByUserId(userId int64) (*[]dao.Video, error) {
	videos, err := dao.GetVideoListByUserId(userId)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

// GetUserInfoByUserId 根据用户id获取用户信息
func GetUserInfoByUserId(userId int64) (*dao.UserInfo, error) {
	userInfo, err := dao.GetUserInfoById(userId)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
