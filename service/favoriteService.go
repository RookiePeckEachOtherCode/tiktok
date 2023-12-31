package service

import (
	"errors"
	"tiktok/dao"
	tiktokLog "tiktok/util/log"
)

const (
	Fav   = 1
	UnFav = 2
)

type FavoriteListReponse struct {
	dao.Response
	VideoList []*dao.Video `json:"video_list"`
}

func FavoriteListService(uid int64) (res *FavoriteListReponse, err error) {
	if !dao.CheckIsExistByID(uid) {
		tiktokLog.Error("用户不存在,uid: ", uid)
		return nil, errors.New("用户不存在")
	}

	favoriteList, err := dao.GetFavoriteList(uid)
	if err != nil {
		return nil, err
	}

	for i := range favoriteList {
		userInfo, err := dao.GetUserInfoById(favoriteList[i].UserInfoID)
		if err == nil {
			favoriteList[i].Author = *userInfo
		}
		//因为是点赞列表，所有的状态都是点赞状态
		favoriteList[i].IsFavorite = true
	}
	return &FavoriteListReponse{
		Response: dao.Response{
			StatusCode: 0,
			StatusMsg:  "获取喜欢列表成功",
		},
		VideoList: favoriteList,
	}, nil

}

func FavoriteActionService(userId, videoId, act int64) error {
	if err := favoriteActionCheck(userId, videoId, act); err != nil {
		return err
	}

	if act == Fav {
		return favorite(userId, videoId)
	} else if act == UnFav {
		return disFavorite(userId, videoId)
	} else {
		return errors.New("未定义的操作类型")
	}
}

func favoriteActionCheck(userId, videoId, act int64) error {
	if userId <= 0 {
		tiktokLog.Error("用户不存在,userId: ", userId)
		return errors.New("用户不存在")
	}
	if act != Fav && act != UnFav {
		tiktokLog.Error("act参数错误:未定义的操作类型,act: ", act)
		return errors.New("act参数错误:未定义的操作类型")
	}
	if act == Fav && (&dao.UserInfo{ID: userId}).GetIsFavorite(videoId) { //点赞动作，视频已经被点赞，就不用再点赞了
		tiktokLog.Error("视频已经被点赞,videoId: ", videoId)
		return errors.New("视频已经被点赞")
	}
	return nil
}

func favorite(userId, videoId int64) error {
	user := &dao.UserInfo{ID: userId}
	//视频的点赞数+1
	//视频作者的获赞数+1
	if err := (user).ToFavoriteVideo(&dao.Video{ID: videoId}); err != nil {
		return err
	}
	return nil
}
func disFavorite(userId, videoIs int64) error {
	user := &dao.UserInfo{ID: userId}
	//视频的点赞数-1
	//视频作者的获赞数-1
	if err := (user).ToCancelFavorite(&dao.Video{ID: videoIs}); err != nil {
		return err
	}
	return nil
}
