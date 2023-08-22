package util

import (
	"errors"
	"fmt"
	"path/filepath"
	"tiktok/configs"
	"tiktok/dao"
	"tiktok/middleware/redis"

	"github.com/importcjj/sensitive"
	uuid "github.com/satori/go.uuid"
)

// UpdateVideoInfo 更新视频信息
// userId: 用户ID
// videos: 视频列表
// 返回值:
// - latestTime: 最新视频的创建时间
// - error: 错误信息
func UpdateVideoInfo(userId int64, videos *[]*dao.Video) error {
	if videos == nil {
		return errors.New("[UpdateVideoInfo] video is nil")
	}

	for i := range *videos {
		userInfo, err := dao.GetUserInfoById((*videos)[i].UserInfoID)
		if err != nil {
			continue
		}
		(*videos)[i].IsFavorite = false
		(*videos)[i].Author = *userInfo

		if userId > 0 {
			userInfo.IsFollow = redis.New(redis.RELATION).GetUserRelation(userId, userInfo.ID)
			(*videos)[i].IsFavorite = redis.New(redis.FAVORITE).GetFavoriteState(userId, (*videos)[i].ID)
		}
	}
	return nil
}

func NewFileName(fileName string) string {
	return uuid.NewV4().String() + filepath.Ext(fileName)
}

func GetFileUrl(name, ty string) string {
	return fmt.Sprintf("http://%v:%v/%v/%v", configs.LAN_IP, configs.GIN_PORT, ty, name)
}

func FilterDirty(str string) (string, error) {
	filter := sensitive.New()
	if err := filter.LoadWordDict(configs.GetDictAbsPath()); err != nil {
		panic(err)
	}
	return filter.Replace(str, '*'), nil
}
