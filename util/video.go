package util

import (
	"errors"
	"tiktok/dao"
	"tiktok/middleware/redis"
	"time"
)

func UpdateVideoInfo(userId int64, videos *[]*dao.Video) (*time.Time, error) {
	videoSize := len(*videos)
	if videos == nil || videoSize == 0 {
		return nil, errors.New("[UpdateVideoInfo] video is nil")
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
