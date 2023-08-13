package service

import (
	"log"
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
	user, err := dao.GetUserInfoById(uid)

	if err != nil {
		log.Println("获取用户信息失败", err)
	}
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
