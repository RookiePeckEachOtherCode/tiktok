package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"tiktok/dao"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

type GetFollowerListResponse struct {
	dao.Response
	UserList []*dao.UserInfo `json:"user_list"`
}

// 用户关注操作
func FollowActionController(c *gin.Context) {
	Tid := c.Query("to_user_id")
	Act := c.Query("action_type")
	_userId, _ := c.Get("user_id")
	userId, ok := _userId.(int64)
	if !ok {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取用户id失败",
		})
		return
	}
	tid, err := strconv.ParseInt(Tid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dao.Response{
			StatusCode: 1,
			StatusMsg:  "对方id获取失败: " + err.Error(),
		})
	}
	act, err := strconv.ParseInt(Act, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dao.Response{
			StatusCode: 1,
			StatusMsg:  "操作类型获取失败:" + err.Error(),
		})
	}
	err = service.FollowActionService(act, tid, userId)
	if err == nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 0,
			StatusMsg:  "关注操作成功",
		})
	} else {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "关注操作失败:" + err.Error(),
		})
		fmt.Printf("%v\n", err)
	}
}

// 获取用户关注列表
func FollowListController(c *gin.Context) {
	Uid := c.Query("user_id")
	uid, err := strconv.ParseInt(Uid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "无法获取用户id: " + err.Error(),
		})
	}
	followList, err := service.FollowListService(uid)
	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "列表获取失败: " + err.Error(),
		})
		fmt.Printf("%v\n", err)
	} else {
		c.JSON(http.StatusOK, followList)
	}
}

// 获取用户粉丝列表
func FollowerListController(c *gin.Context) {
	_userId := c.Query("user_id")
	userId, err := strconv.ParseInt(_userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "用户id解析失败: " + err.Error(),
		})
		return
	}

	userList, err := service.FollowerListService(userId)

	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取关注列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, GetFollowerListResponse{
		Response: dao.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserList: userList,
	})
}
