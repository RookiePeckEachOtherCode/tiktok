package util

import (
	"errors"
	"fmt"
	"path/filepath"
	"tiktok/configs"
	"tiktok/dao"
	"tiktok/middleware/redis"
	"time"

	uuid "github.com/satori/go.uuid"
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
	latestTime := (*videos)[videoSize-1].CreatedAt
	for i := 0; i < videoSize; i++ {
		userInfo, err := dao.GetUserInfoById((*videos)[i].UserInfoID)
		if err != nil {
			continue
		}
		userInfo.IsFollow = redis.New().GetUserRelation(userId, userInfo.ID)
		(*videos)[i].Author = *userInfo
		if userId > 0 {
			if redis.IsInit {
				(*videos)[i].IsFavorite = redis.New().GetFavoriteState(userId, (*videos)[i].ID)
			} else {
				_user := dao.UserInfo{ID: userId}
				(*videos)[i].IsFavorite = _user.GetIsFavorite((*videos)[i].ID)
			}
		}
	}
	return &latestTime, nil
}

func NewFileName(fileName string) string {
	return uuid.NewV4().String() + filepath.Ext(fileName)
}

func GetFileUrl(name, ty string) string {
	return fmt.Sprintf("http://%v:%v/%v/%v", configs.LAN_IP, configs.GIN_PORT, ty, name)
}
