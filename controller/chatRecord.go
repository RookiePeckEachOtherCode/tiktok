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
	MsgList []dao.ChatRecord `json:"message_list"`
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
	}
	//now   1692058913
	//need  1692059113053
	//nano  1692059223836730265
	//micro 1692059223836745
	//milli 1692059223836
	//unix  1692059223

	msgList, err := service.GetChatRecord(userId, toUserId, preMsgTime)

	for _, v := range msgList {
		util.PrintLog(fmt.Sprintf("content:%v created_at: %v", v.Content, v.CreatedAt))
	}

	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取聊天记录失败" + err.Error(),
		})
	}

	c.JSON(http.StatusOK, ChatRecordResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取聊天记录成功",
		},
		MsgList: msgList,
	})
}
