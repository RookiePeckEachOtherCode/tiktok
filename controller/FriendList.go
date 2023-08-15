package controller

import (
	"log"
	"net/http"
	"strconv"
	"tiktok/dao"
	"tiktok/model"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

type FriendListResponse struct {
	model.Response
	FriendList []*dao.Friend `json:"user_list"`
}

func GetFriendList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

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
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		FriendList: userList,
	})
}
