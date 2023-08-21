package service

import (
	"errors"
	"fmt"
	"tiktok/dao"
	"tiktok/middleware/redis"
	tiktokLog "tiktok/util/log"
)

type FollowerResponse struct {
	dao.Response
	UserList []*dao.UserInfo `json:"user_list"`
}

func FollowActionService(act int64, tid int64, uid int64) error {
	if !dao.CheckIsExistByID(tid) {
		tiktokLog.Error(fmt.Sprintf("用户不存在,tid: %d", tid))
		return errors.New("用户不存在")
	}
	if tid == uid {
		tiktokLog.Error("不能关注自己,tid: ", tid)
		return errors.New("不能关注自己")
	}

	if act == 1 {
		if !dao.GetUserRelation(uid, tid) {
			if err := (&dao.UserInfo{ID: uid}).FollowAct(&dao.UserInfo{ID: tid}); err != nil {
				tiktokLog.Error(fmt.Sprintf("关注失败,uid: %d, tid: %d, Error:%v", uid, tid, err))
				return err
			}
			redis.New(redis.RELATION).UpdateUserRelation(uid, tid, true)
			return nil
		} else {
			tiktokLog.Error(fmt.Sprintf("当前用户已关注,uid: %d, tid: %d", uid, tid))
			return errors.New("当前用户已关注")
		}
	}
	if act == 2 {
		if err := (&dao.UserInfo{ID: uid}).UnFollowAct(&dao.UserInfo{ID: tid}); err != nil {
			return err
		}
		redis.New(redis.RELATION).UpdateUserRelation(uid, tid, false)
		return nil
	} else {
		return errors.New("非法操作")
	}
}

func FollowListService(uid int64) (*FollowerResponse, error) {
	list, err := dao.GetFloList(uid)
	if err != nil {
		return nil, err
	}
	res := &FollowerResponse{
		Response: dao.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserList: list,
	}
	return res, nil
}

func FollowerListService(userId int64) ([]*dao.UserInfo, error) {
	if !dao.CheckIsExistByID(userId) {
		tiktokLog.Error(fmt.Sprintf("用户不存在,userId: %d", userId))
		return nil, errors.New("该用户不存在")
	}

	userList, err := dao.GetFollowerList(userId)

	if err != nil {
		return nil, err
	}

	for _, v := range userList {
		v.IsFollow = redis.New(redis.RELATION).GetUserRelation(userId, v.ID)
	}

	return userList, nil
}
