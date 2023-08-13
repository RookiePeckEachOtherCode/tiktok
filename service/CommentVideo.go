package service

import (
	"tiktok/dao"
)

func PostComment(vid, uid int64, context string) (*dao.Comment, error) {
	comment := &dao.Comment{
		UserInfoID: uid,
		VideoID:    vid,
		Content:    context,
	}
	if err := comment.PostComment(); err != nil {
		return nil, err
	}
	user, _ := dao.GetUserInfoById(uid)
	comment.User = *user
	return comment, nil
}

func DeleteComment(commentId string) error {
	comment, err := dao.FindComment(commentId)
	if err != nil {
		return err
	}
	if err := comment.DeleteComment(); err != nil {
		return err
	}
	return nil
}
