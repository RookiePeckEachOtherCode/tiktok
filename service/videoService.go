package service

import (
	"errors"
	"fmt"
	"tiktok/dao"
	"tiktok/middleware/redis"
	"tiktok/util"
	"time"
)

// FeedVideoFlow 包含下次请求的最早发布时间和视频列表
type FeedVideoFlow struct {
	NextTime  int64         `json:"next_time,omitempty"`  //发布最早的时间，作为下次请求时的latest_time
	VideoList *[]*dao.Video `json:"video_list,omitempty"` //视频列表
}

func VideoFeedService(userId int64, latestTime time.Time) (*FeedVideoFlow, error) {
	if latestTime.IsZero() {
		latestTime = time.Now()
	}
	videoList, err := dao.GetVideoListByLastTime(latestTime)
	if err != nil {
		return nil, fmt.Errorf("获取视频fedd失败：%v", err)
	}
	_latestTime, _ := util.UpdateVideoInfo(userId, videoList)
	var nextTime int64
	if _latestTime != nil {
		nextTime = (*_latestTime).UnixNano() / 1e6
	} else {
		nextTime = time.Now().Unix() / 1e6
	}
	return &FeedVideoFlow{
		NextTime:  nextTime,
		VideoList: videoList,
	}, nil
}

// PublishListService 获取发布列表
func PublishListService(userId int64) (*[]dao.Video, error) {
	userInfo, videos, err := publishListCheck(userId)
	if err != nil {
		return nil, err
	}

	// 为视频列表添加作者信息
	util.PrintLog(fmt.Sprintf("调用了service的获取发布列表函数，userId:%d", userId))
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
