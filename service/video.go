package service

import (
	"tiktok/dao"
	"tiktok/util"
	"time"
)

type Respons struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	FeedVideoFlow
}

type FeedVideoFlow struct {
	NextTime  int64        `json:"next_time"`  //发布最早的时间，作为下次请求时的latest_time
	VideoList []*dao.Video `json:"video_list"` //视频列表
}

type GetFeedVideoListFlow struct {
	latestTime time.Time
	userId     int64
	videos     []*dao.Video
	feedVideo  *FeedVideoFlow
	nextTime   int64
}

func Feed(lastTime time.Time, userId int64) (*FeedVideoFlow, error) {
	getFlow := GetFeedVideoListFlow{latestTime: lastTime, userId: userId}
	return getFlow.Do()
}

// 流程处理方式
func (g *GetFeedVideoListFlow) Do() (*FeedVideoFlow, error) {
	if err := g.Init(); err != nil {
		return nil, err
	}
	if err := g.Pack(); err != nil {
		return nil, err
	}
	return g.feedVideo, nil
}

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

func (g *GetFeedVideoListFlow) Pack() error {
	g.feedVideo = &FeedVideoFlow{
		NextTime:  g.nextTime,
		VideoList: g.videos,
	}
	return nil
}
