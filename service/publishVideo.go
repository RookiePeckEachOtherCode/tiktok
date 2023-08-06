package service

import (
	"tiktok/configs"
	"tiktok/dao"
	"tiktok/util"
)

type SaveVideoInfo struct {
	userId    int64
	videoPath string
	coverPath string
	titile    string
}

func PublishVideo(userID int64, videoName, coverName, title string) error {
	v := &SaveVideoInfo{userId: userID, titile: title}

	v.videoPath = util.GetFileUrl(videoName, configs.VIDEO_SAVE_PATH)
	v.coverPath = util.GetFileUrl(coverName, configs.VIDEO_COVER_SAVE_PATH)

	err := dao.NewVideo(&dao.Video{
		Title:      v.titile,
		UserInfoID: v.userId,
		PlayURL:    v.videoPath,
		CoverURL:   v.coverPath,
	})

	return err
}
