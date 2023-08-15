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

type ChatRecordResponse struct {
	model.Response
	MessageList []dao.ChatRecord `json:"message_list"`
}

func GetChatRecord(c *gin.Context) {
	_toUserId := c.Query("to_user_id")
	if _toUserId == "" {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "to_user_id不能为空",
		})
		return
	}
	toUserId, _ := strconv.ParseInt(_toUserId, 10, 64)

	_userId, _ := c.Get("user_id")
	userId, ok := _userId.(int64)

	if !ok {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户id错误",
		})
		return
	}

	pre_msg_time := c.Query("pre_msg_time")
	preMsgTime, err := strconv.ParseInt(pre_msg_time, 10, 64)

	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "无效的时间戳",
		})
		return
	}

	messagelist, err := service.GetChatRecord(userId, toUserId, preMsgTime)

	for _, v := range messagelist {
		util.PrintLog(fmt.Sprintln("from_user_id:", v.FromUserId, "to_user_id:", v.ToUserId, "content:", v.Content, "create_time:", v.CreatedTime))
	}

	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取聊天记录失败" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ChatRecordResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取聊天记录成功",
		},
		MessageList: messagelist,
	})
}
