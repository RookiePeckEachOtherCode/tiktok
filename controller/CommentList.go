package controller

import (
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

func RecComList(c *gin.Context) {
	Vid := c.Query("video_id")
	vid, err := strconv.ParseInt(Vid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "无法获取视频id",
		})
	}
	res, err := service.HandleComListQuery(vid)
	if err != nil {
		c.JSON(http.StatusBadRequest, service.FavListRes{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "获取评论列表失败",
			},
		})
	} else {
		c.JSON(http.StatusOK, res)
	}
}
