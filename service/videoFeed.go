package service

import (
	"fmt"
	"tiktok/dao"
	"tiktok/util"
	"time"
)

// FeedVideoFlow 包含下次请求的最早发布时间和视频列表
type FeedVideoFlow struct {
	NextTime  int64        `json:"next_time"`  //发布最早的时间，作为下次请求时的latest_time
	VideoList []*dao.Video `json:"video_list"` //视频列表
}

func Feed(userId int64, lastTime time.Time) ([]*dao.Video, int64, error) {
	videos, err := dao.GetVideoListByLastTime(lastTime)
	if err != nil {
		return nil, 0, fmt.Errorf("dao.GetVideoListByLastTime error: %v", err)
	}
	lastestTime, _ := util.UpdateVideoInfo(userId, &videos)
	if lastestTime != nil {
		//latestTime不为空表示还有更早的视频,使用其时间戳获取下一页
		nextTime := (lastestTime.UnixNano() / int64(time.Millisecond))
		return videos, nextTime, nil
	}
	//latestTime为空表示已取到最新视频,需要使用当前时间获取下一页
	nextTime := (time.Now().UnixNano() / int64(time.Millisecond))

	return videos, nextTime, nil
}
