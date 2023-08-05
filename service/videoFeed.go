package service

import (
	"tiktok/dao"
	"tiktok/util"
	"time"
)

// FeedVideoFlow 包含下次请求的最早发布时间和视频列表
type FeedVideoFlow struct {
	NextTime  int64        `json:"next_time"`  //发布最早的时间，作为下次请求时的latest_time
	VideoList []*dao.Video `json:"video_list"` //视频列表
}

// GetFeedVideoListFlow 包含最新时间、用户ID、视频列表、视频流信息和下次请求的最早发布时间
type GetFeedVideoListFlow struct {
	latestTime time.Time
	userId     int64
	videos     []*dao.Video
	feedVideo  *FeedVideoFlow
	nextTime   int64
}

// Feed 返回视频流信息
func Feed(lastTime time.Time, userId int64) (*FeedVideoFlow, error) {
	getFlow := GetFeedVideoListFlow{latestTime: lastTime, userId: userId}
	return getFlow.Do()
}

// Do 处理视频流信息
func (g *GetFeedVideoListFlow) Do() (*FeedVideoFlow, error) {
	if err := g.Init(); err != nil {
		return nil, err
	}
	if err := g.Pack(); err != nil {
		return nil, err
	}
	return g.feedVideo, nil
}

// Init 初始化视频列表和下次请求的最早发布时间
func (g *GetFeedVideoListFlow) Init() error {
	//TODO 给视频添加点赞状态
	//TODO 给登陆的用户 添加视频点赞的状态
	var err error
	g.videos, err = dao.GetVideoListByLastTime(g.latestTime)

	if err != nil {
		return err
	}

	latestTime, _ := util.UpdateVideoInfo(g.userId, &g.videos)

	if latestTime != nil {
		g.nextTime = latestTime.UnixNano() / int64(time.Millisecond)
		return nil
	}

	g.nextTime = time.Now().Unix() / int64(time.Millisecond)

	return nil
}

// Pack 打包视频流信息
func (g *GetFeedVideoListFlow) Pack() error {
	g.feedVideo = &FeedVideoFlow{
		NextTime:  g.nextTime,
		VideoList: g.videos,
	}
	return nil
}
