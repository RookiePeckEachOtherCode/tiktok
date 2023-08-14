package service

import (
	"errors"
	"log"
	"tiktok/dao"
	"tiktok/middleware/redis"
)

func HandleFollowAct(act int64, tid int64, uid int64) error {

	if !dao.CheckIsExistByID(tid) {
		log.Println("用户不存在")
		return errors.New("用户不存在")
	}
	if tid == uid {
		log.Println("不能关注自己")
		return errors.New("不能关注自己")
	}

	if act == 1 {
		if !dao.GetUserRelation(uid, tid) {
			if err := (&dao.UserInfo{ID: uid}).FollowAct(&dao.UserInfo{ID: tid}); err != nil {
				return err
			}
			redis.New(redis.RELATION).UpdateUserRelation(uid, tid, true)
			return nil
		} else {
			log.Println("当前用户已关注")
			return errors.New("当前用户已关注")
		}
	}
	if act == 2 {
		if err := (&dao.UserInfo{ID: uid}).UnFollowAct(&dao.UserInfo{ID: tid}); err != nil {
			log.Println("取消关注失败")
			return err
		}
		redis.New(redis.RELATION).UpdateUserRelation(uid, tid, false)
		return nil
	} else {
		return errors.New("非法操作")
	}
}
