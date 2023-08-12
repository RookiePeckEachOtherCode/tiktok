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
	if err := FavoriteActCheck(userId, videoId, act); err != nil {
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

func FavoriteActCheck(userId, videoId, act int64) error {
	if userId <= 0 {
		return errors.New("用户不存在")
	}
	if act != Fav && act != UnFav {
		return errors.New("act参数错误:未定义的操作类型")
	}
	if act == Fav && (&dao.UserInfo{ID: userId}).GetIsFavorite(videoId) { //点赞动作，视频已经被点赞，就不用再点赞了
		return errors.New("视频已经被点赞")
	}
	return nil
}

func ActFav(userId, videoId int64) error {
	user := &dao.UserInfo{ID: userId}
	//视频的点赞数+1
	//视频作者的获赞数+1
	if err := (user).ToFavoriteVideo(&dao.Video{ID: videoId}); err != nil {
		return err
	}
	return nil
}
func ActUnFav(userId, videoIs int64) error {
	user := &dao.UserInfo{ID: userId}
	//视频的点赞数-1
	//视频作者的获赞数-1
	if err := (user).ToCancelFavorite(&dao.Video{ID: videoIs}); err != nil {
		return err
	}
	return nil
}
