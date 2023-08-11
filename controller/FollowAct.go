package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

func FollowAct(c *gin.Context) {
	Tid := c.Query("to_user_id")
	Act := c.Query("action_type")
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
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: 1,
			StatusMsg:  "对方id获取失败",
		})
	}
	act, err := strconv.ParseInt(Act, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: 1,
			StatusMsg:  "操作类型获取失败:" + err.Error(),
		})
	}
	err = service.HandleFollowAct(act, tid, userId)
	if err == nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 0,
			StatusMsg:  "关注操作成功",
		})
	} else {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "关注操作失败:" + err.Error(),
		})
		fmt.Printf("%v\n", err)
	}
}
