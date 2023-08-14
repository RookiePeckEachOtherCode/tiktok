package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/service"
)

func FriendList(c *gin.Context) {
	Uid := c.Query("user_id")
	uid, err := strconv.ParseInt(Uid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取用户Id失败",
		})
	}
	friendListRes, err := service.HandleFriendList(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: 1,
			StatusMsg:  "好友列表获取失败",
		})
	} else {
		c.JSON(http.StatusOK, friendListRes)
	}
}
