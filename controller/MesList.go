package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/service"
)

func MesList(c *gin.Context) {
	Tuid := c.Query("to_user_id")
	tuid, err := strconv.ParseInt(Tuid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取对象Id失败",
		})
	}
	_userId, _ := c.Get("user_id")
	userId, ok := _userId.(int64)
	if !ok {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取用户id失败",
		})
		return
	}
	mesListRes, err := service.HandleMesList(userId, tuid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取用户id失败",
		})
		fmt.Printf("%v\n", err)
	} else {
		c.JSON(http.StatusOK, mesListRes)
	}
}
