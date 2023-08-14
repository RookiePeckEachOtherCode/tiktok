package controller

import (
	"log"
	"net/http"
	"tiktok/dao"
	"tiktok/model"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

type FriendListResponse struct {
	model.Response
	UserList []*dao.UserInfo `json:"user_list"`
}

func GetFriendList(c *gin.Context) {
	_userId, _ := c.Get("user_id")
	userId, ok := _userId.(int64)

	if !ok {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户id解析错误",
		})

	}

	userList, err := service.GetFriendList(userId)

	if err != nil {
		log.Println("获取好友列表失败", err)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取好友列表失败" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, FriendListResponse{
		model.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		userList,
	})
}
