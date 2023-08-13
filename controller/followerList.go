package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"tiktok/dao"
	"tiktok/model"
	"tiktok/service"
	"tiktok/util"

	"github.com/gin-gonic/gin"
)

type GetFollowerListResponse struct {
	model.Response
	UserList []*dao.UserInfo `json:"user_list"`
}

func GetFollowerList(c *gin.Context) {
	_userId := c.Query("user_id")
	userId, err := strconv.ParseInt(_userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户id解析失败",
		})
		return
	}

	util.PrintLog(fmt.Sprintf("user_id:%v", userId))

	userList, err := service.GetFollowerList(userId)

	if err != nil {
		util.PrintLog(fmt.Sprintf("获取关注列表失败，err:%v", err))
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	util.PrintLog(fmt.Sprintf("获取关注列表成功，userList:%v", userList))
	c.JSON(http.StatusOK, GetFollowerListResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		UserList: userList,
	})
}
