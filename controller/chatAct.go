package controller

import (
	"log"
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

func ChatActino(c *gin.Context) {
	//get user_id
	_userId, _ := c.Get("user_id")
	userId, ok := _userId.(int64)
	if !ok {
		log.Println("用户id错误")
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户id错误",
		})
		return
	}
	//get to_user_id
	_toUserId := c.Query("to_user_id")
	toUserId, _ := strconv.ParseInt(_toUserId, 10, 64)
	//get content
	content := c.Query("content")
	//get action_type
	actionType := c.Query("action_type")

	if actionType != "1" {
		log.Println("action_type错误")
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "action_type错误",
		})
		return
	}

	err := service.PostMessage(userId, toUserId, content)

	if err != nil {
		log.Println("发送消息失败" + err.Error())
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "发送消息失败" + err.Error(),
		})
		return
	}
	log.Println("发送消息成功")
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  "发送消息成功",
	})

}
