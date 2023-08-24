package service

import (
	"errors"
	"fmt"
	"tiktok/dao"
	"tiktok/middleware/redis"
	"tiktok/util"
	tiktokLog "tiktok/util/log"
	"time"
)

// FeedVideoFlow 包含下次请求的最早发布时间和视频列表
type SaveVideoInfo struct {
	userId    int64
	videoPath string
	coverPath string
	title     string
}

func VideoFeedService(userId int64, latestTime time.Time) (*[]*dao.Video, error) {
	videoList, err := dao.GetVideoListByLastTime(latestTime)
	if err != nil {
		tiktokLog.Error(fmt.Sprintf("获取视频fedd失败：%v, 时间: %v, userId: %d", err, latestTime, userId))
		return nil, fmt.Errorf("获取视频fedd失败：%v", err)
	}
	err = util.UpdateVideoInfo(userId, videoList)

	if err != nil {
		return nil, err
	}

	return videoList, nil
}

// PublishListService 获取发布列表
func PublishListService(userId int64) (*[]dao.Video, error) {
	userInfo, videos, err := publishListCheck(userId)
	if err != nil {
		return nil, err
	}

	// 为视频列表添加作者信息
	for i := range *videos {
		(*videos)[i].Author = *userInfo
		(*videos)[i].IsFavorite = redis.New(redis.FAVORITE).GetFavoriteState(userId, (*videos)[i].ID)
	}
	return videos, nil
}

func PublishVideoService(userID int64, videoSavePath, coverSavePath, title string) error {
	err := dao.NewVideo(&dao.Video{
		Title:      title,
		UserInfoID: userID,
		PlayURL:    videoSavePath,
		CoverURL:   coverSavePath,
	})
	return err
}

// publishListCheck 发布列表检查
func publishListCheck(userId int64) (*dao.UserInfo, *[]dao.Video, error) {
	if err := checkIsExistByID(userId); err != nil {
		return nil, nil, err
	}

	userInfo, err := getUserInfoByUserId(userId)
	if err != nil {
		return nil, nil, err
	}

	videos, err := getVideoListByUserId(userId)
	if err != nil {
		return nil, nil, err
	}

	return userInfo, videos, nil
}

// checkIsExistByID 根据用户id检查用户是否存在
func checkIsExistByID(userId int64) error {
	if !dao.CheckIsExistByID(userId) {
		tiktokLog.Error(fmt.Sprintf("用户不存在,userId: %d", userId))
		return errors.New("该用户不存在")
	}
	return nil
}

// getVideoListByUserId 根据用户id获取视频列表
func getVideoListByUserId(userId int64) (*[]dao.Video, error) {
	videos, err := dao.GetVideoListByUserId(userId)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

// getUserInfoByUserId 根据用户id获取用户信息
func getUserInfoByUserId(userId int64) (*dao.UserInfo, error) {
	userInfo, err := dao.GetUserInfoById(userId)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
