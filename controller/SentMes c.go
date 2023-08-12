package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/service"
)

func SentMes(c *gin.Context) {
	Tid := c.Query("to_user_id")
	Act := c.Query("action_type")
	content := c.Query("content")
	_userId, _ := c.Get("user_id")
	userId, ok := _userId.(int64)
	if !ok {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取用户id失败",
		})
		return
	}
	tid, err := strconv.ParseInt(Tid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取对象id失败",
		})
	}
	act, err := strconv.ParseInt(Act, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获操作类型失败",
		})
	}
	err = service.HandleSentMes(act, content, userId, tid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "发送信息失败",
		})
	} else {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 0,
			StatusMsg:  "发送成功",
		})
	}

}
