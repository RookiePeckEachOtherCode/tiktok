package service

import (
	"errors"
	"log"
	"tiktok/dao"
	"tiktok/middleware/redis"
)

type FollowerResponse struct {
	dao.Response
	UserList []*dao.UserInfo `json:"user_list"`
}

func FollowActionService(act int64, tid int64, uid int64) error {

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
