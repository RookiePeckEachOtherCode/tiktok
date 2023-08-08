package service

import "tiktok/dao"

func HandleFav(Vedioid int64, actiontype int64, uid int64) error {
	video, err := dao.FindVideoByVid(Vedioid) //先找到要修改的视频对象
	if err != nil {
		return err
	}
	err2 := dao.FavoriteVideo(video, actiontype, uid) //增加或减少点赞数目
	if err2 != nil {
		return err2
	}
	return nil
}
