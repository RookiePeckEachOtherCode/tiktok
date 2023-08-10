package service

import (
	"errors"
	"tiktok/dao"
	"tiktok/util"
)

func OperateComment(videoid int64, userid int64, act int64, text string, comment *dao.Comment) error { //1发布评论，2删除评论
	if err := CommentActCheck(userid, videoid, act); err != nil {
		util.PrintLog(err.Error())
		return err
	}
	if act == 1 {
		return Act1(userid, videoid, text, comment)
	} else if act == 2 {
		return Act2(userid, videoid, text)
	} else {
		return errors.New("未定义的操作类型")
	}
	return nil
}
func CommentActCheck(userId, videoId, act int64) error {
	if userId <= 0 {
		return errors.New("用户不存在")
	}
	if act != 1 && act != 2 {
		return errors.New("act参数错误:未定义的操作类型")
	}
	return nil
}
func Act1(userId, videoId int64, text string, comment *dao.Comment) error {
	user := &dao.UserInfo{ID: userId}
	//视频的评论数+1
	if err := (user).PostComment(text, &dao.Video{ID: videoId}, comment); err != nil {
		return err
	}
	return nil
}
func Act2(userId, videoIs int64, text string) error {
	user := &dao.UserInfo{ID: userId}
	//视频的评论数-1
	if err := (user).ToCancelFavorite(&dao.Video{ID: videoIs}); err != nil {
		return err
	}
	return nil
}
