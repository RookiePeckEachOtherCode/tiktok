package service

import (
	"errors"
	"tiktok/dao"
	"tiktok/util"
)

// 1-点赞，2-取消点赞
const (
	Fav   = 1
	UnFav = 2
)

func HandleFav(userId, videoId, act int64) error {
	if err := FavoriteActCheck(userId, act); err != nil {
		util.PrintLog(err.Error())
		return err
	}

	if act == Fav {
		return ActFav(userId, videoId)
	} else if act == UnFav {
		return ActUnFav(userId, videoId)
	} else {
		return errors.New("未定义的操作类型")
	}
}

func FavoriteActCheck(userId, act int64) error {
	if userId <= 0 {
		return errors.New("用户不存在")
	}
	if act != Fav && act != UnFav {
		return errors.New("act参数错误:未定义的操作类型")
	}
	return nil
}

func ActFav(userId, videoId int64) error {
	err := (&dao.UserInfo{ID: userId}).ToFavoriteVideo(&dao.Video{ID: videoId})
	if err != nil {
		return err
	}
	return nil
}
func ActUnFav(userId, videoIs int64) error {
	err := (&dao.UserInfo{ID: userId}).ToCancelFavorite(&dao.Video{ID: videoIs})
	if err != nil {
		return err
	}
	return nil

}
