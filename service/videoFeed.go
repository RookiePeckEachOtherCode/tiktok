package service

import (
	"fmt"
	"tiktok/dao"
	"tiktok/util"
	"time"
)

// FeedVideoFlow 包含下次请求的最早发布时间和视频列表
type FeedVideoFlow struct {
	NextTime  int64         `json:"next_time"`  //发布最早的时间，作为下次请求时的latest_time
	VideoList *[]*dao.Video `json:"video_list"` //视频列表
}

func Feed(userId int64, latestTime time.Time) (*FeedVideoFlow, error) {
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
