package service

import (
	"tiktok/dao"
	"tiktok/model"
)

type FavListRes struct {
	model.Response
	VideoList []*dao.Video `json:"video_list"`
}

func HandleFavListQuery(uid int64) (res *FavListRes, err error) {
	vlist, err := dao.GetFavList(uid)
	if err != nil {
		return nil, err
	}
	res = &FavListRes{}
	res.VideoList = vlist
	res.StatusCode = 0
	res.StatusMsg = "获取列表成功"
	return res, nil
}
