package service

import (
	"errors"
	"tiktok/dao"
)

func HandleFollowAct(act int64, tid int64, uid int64) error {
	if act == 1 {
		err := (&dao.UserInfo{ID: uid}).FollowAct(&dao.UserInfo{ID: tid})
		return err
	}
	if act == 2 {
		err := (&dao.UserInfo{ID: uid}).UnFollowAct(&dao.UserInfo{ID: tid})
		return err
	}
	return errors.New("未处理关注请求")
}
