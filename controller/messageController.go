package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tiktok/dao"
	"tiktok/service"
	"tiktok/util"

	"github.com/gin-gonic/gin"
)

type ChatRecordListResponse struct {
	dao.Response
	MessageList []dao.ChatRecord `json:"message_list"`
}

type FriendListResponse struct {
	dao.Response
	FriendList []*dao.Friend `json:"user_list"`
}

func ChatActionController(c *gin.Context) {
	//get user_id
	_userId, _ := c.Get("user_id")
	userId, ok := _userId.(int64)
	if !ok {
		log.Println("用户id错误")
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "用户id错误",
		})
		return
	}
	//get to_user_id
	_toUserId := c.Query("to_user_id")
	toUserId, _ := strconv.ParseInt(_toUserId, 10, 64)
	//get content
	_content, _ := c.Get("content")
	content, _ := _content.(string)

	//get action_type
	actionType := c.Query("action_type")

	if actionType != "1" {
		log.Println("action_type错误")
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "action_type错误",
		})
		return
	}

	err := service.ChatActionService(userId, toUserId, content)

	if err != nil {
		log.Println("发送消息失败" + err.Error())
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "发送消息失败" + err.Error(),
		})
		return
	}
	log.Println("发送消息成功")
	c.JSON(http.StatusOK, dao.Response{
		StatusCode: 0,
		StatusMsg:  "发送消息成功",
	})

}

func ChatRecordListController(c *gin.Context) {
	_toUserId := c.Query("to_user_id")
	if _toUserId == "" {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "to_user_id不能为空",
		})
		return
	}
	toUserId, _ := strconv.ParseInt(_toUserId, 10, 64)

	_userId, _ := c.Get("user_id")
	userId, ok := _userId.(int64)

	if !ok {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "用户id错误",
		})
		return
	}

	preMsgTime, err := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)

	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "无效的时间戳",
		})
		return
	}

	messageList, err := service.ChatRecordService(userId, toUserId, preMsgTime)

	for _, v := range messageList {
		util.PrintLog(fmt.Sprintln("from_user_id:", v.FromUserId, "to_user_id:", v.ToUserId, "content:", v.Content, "create_time:", v.CreatedTime))
	}

	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取聊天记录失败" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ChatRecordListResponse{
		Response: dao.Response{
			StatusCode: 0,
			StatusMsg:  "获取聊天记录成功",
		},
		MessageList: messageList,
	})
}

func FriendListController(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	userList, err := service.FriendListService(userId)

	if err != nil {
		log.Println("获取好友列表失败", err)
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取好友列表失败" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, FriendListResponse{
		Response: dao.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		FriendList: userList,
	})
}
