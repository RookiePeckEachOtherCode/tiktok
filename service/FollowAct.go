package service

import (
	"errors"
	"tiktok/dao"
)

func HandleFollowAct(act int64, tid int64, uid int64) error {
	if act == 1 {
		b := dao.GetUserRelation(uid, tid)
		if b == false {
			err := (&dao.UserInfo{ID: uid}).FollowAct(&dao.UserInfo{ID: tid})
			return err
		} else {
			return errors.New("当前用户已关注")
		}
	}
	if act == 2 {
		err := (&dao.UserInfo{ID: uid}).UnFollowAct(&dao.UserInfo{ID: tid})
		return err
	}
	return errors.New("未处理关注请求")
}
