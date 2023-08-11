package service

import (
	"errors"
	"log"
	"tiktok/dao"
	"tiktok/model"
)

type FavListRes struct {
	model.Response
	VideoList []*dao.Video `json:"video_list"`
}

func HandleFavListQuery(uid int64) (res *FavListRes, err error) {
	if !dao.CheckIsExistByID(uid) {
		log.Println("用户不存在")
		return nil, errors.New("用户不存在")
	}

	vlist, err := dao.GetFavList(uid)
	if err != nil {
		log.Println("获取喜欢列表失败:", err)
		return nil, err
	}

	for i := range vlist {
		userInfo, err := dao.GetUserInfoById(vlist[i].UserInfoID)
		if err == nil {
			vlist[i].Author = *userInfo
		}
		//因为是点赞列表，所有的状态都是点赞状态
		vlist[i].IsFavorite = true
	}
	log.Println("获取喜欢列表成功")
	return &FavListRes{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取喜欢列表成功",
		},
		VideoList: vlist,
	}, nil

}
