package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"tiktok/dao"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

// 用户点赞
func FavoriteActionController(c *gin.Context) {
	_userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取用户id失败",
		})
		return
	}
	userId, ok := _userId.(int64)
	if !ok {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取用户id失败",
		})
		return
	}

	_videoId := c.Query("video_id")

	videoId, err := strconv.ParseInt(_videoId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取视频id失败: " + err.Error(),
		})
		return
	}

	_act := c.Query("action_type")

	act, err := strconv.ParseInt(_act, 10, 64)

	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取act失败: " + err.Error(),
		})
		return
	}

	err = service.FavoriteActionService(userId, videoId, act)

	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "操作失败: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, dao.Response{
		StatusCode: 0,
		StatusMsg:  "success",
	})
}

// 获取喜欢列表
func FavoriteListController(c *gin.Context) {
	_uid, _ := c.Get("user_id")
	uid, ok := _uid.(int64)

	if !ok {
		c.JSON(http.StatusBadRequest, service.FavoriteListReponse{
			Response: dao.Response{
				StatusCode: 1,
				StatusMsg:  "用户id解析失败",
			},
		})
		return
	}

	res, err := service.FavoriteListService(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, service.FavoriteListReponse{
			Response: dao.Response{
				StatusCode: 1,
				StatusMsg:  "获取喜欢列表失败",
			},
		})
		fmt.Printf("%v\n", err)
	} else {
		c.JSON(http.StatusOK, res)
	}

}
