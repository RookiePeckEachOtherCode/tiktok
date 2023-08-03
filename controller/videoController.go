package controller

import (
	"net/http"
	"strconv"
	"tiktok/service"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type FeedResponse struct {
	Response
	FeedFlow service.FeedVideoFlow
}

func Feed(c *gin.Context) {
	inputTime := ParseInputTime(c)
	userId := GetUserId(c)

	vidoFlow, err := service.Feed(inputTime, userId)

	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "获取失败",
			},
		})
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "获取成功",
		},
		FeedFlow: *vidoFlow,
	})
}

func ParseInputTime(c *gin.Context) time.Time {
	inputTime := c.Query("latest_time")

	if len(inputTime) != 0 {
		tempTime, _ := strconv.ParseInt(inputTime, 10, 64)

		return time.Unix(tempTime, 0)
	}
	return time.Now()
}
func GetUserId(c *gin.Context) int64 {
	userId, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
	return userId
}
