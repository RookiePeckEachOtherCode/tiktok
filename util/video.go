package util

import (
	"errors"
	"tiktok/dao"
	"tiktok/middleware/redis"
	"time"
)

// UpdateVideoInfo 更新视频信息
// userId: 用户ID
// videos: 视频列表
// 返回值:
// - latestTime: 最新视频的创建时间
// - error: 错误信息
func UpdateVideoInfo(userId int64, videos *[]*dao.Video) (*time.Time, error) {
	if videos == nil {
		return nil, errors.New("[UpdateVideoInfo] video is nil")
	}

	videoSize := len(*videos)
	if videoSize == 0 {
		return nil, errors.New("[UpdateVideoInfo] video size is 0")
	}

	p := redis.NewProxyIndexMap()
	latestTime := (*videos)[videoSize-1].CreatedAt

	for i := 0; i < videoSize; i++ {
		userInfo, err := dao.GetUserInfoById((*videos)[i].UserInfoID)

		if err != nil {
			continue
		}
		userInfo.IsFollow = p.GetUserRelation(userId, userInfo.ID)

		(*videos)[i].Author = userInfo

		if userId > 0 {
			(*videos)[i].IsFavorite = p.GetFavorateState(userId, (*videos)[i].ID)
		}
	}

	return &latestTime, nil
}
