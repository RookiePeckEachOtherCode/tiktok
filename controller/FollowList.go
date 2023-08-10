package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/service"
)

func FollowList(c *gin.Context) {
	Uid := c.Query("user_id")
	uid, err := strconv.ParseInt(Uid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "无法获取用户id",
		})
	}
	respon, err := service.HandleFollowList(uid)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "列表获取失败",
		})
		fmt.Printf("%v\n", err)
	} else {
		c.JSON(http.StatusOK, respon)
	}
}
