package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/service"
	"tiktok/util"

	"github.com/gin-gonic/gin"
)

func FavoriteAct(c *gin.Context) {
	_userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取用户id失败1",
		})
		return
	}
	userId, ok := _userId.(int64)
	if !ok {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取用户id失败2",
		})
		return
	}

	_videoId := c.Query("video_id")

	videoId, err := strconv.ParseInt(_videoId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取视频id失败",
		})
		return
	}

	_act := c.Query("action_type")

	act, err := strconv.ParseInt(_act, 10, 64)

	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取act失败",
		})
		return
	}

	util.PrintLog(fmt.Sprintf("userId:%d,videoId:%d,act:%d", userId, videoId, act))

	err = service.HandleFav(userId, videoId, act)

	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  fmt.Errorf("操作失败: %w", err).Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  "success",
	})
}