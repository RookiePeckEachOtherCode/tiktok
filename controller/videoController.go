package controller

import (
	"net/http"
	"strconv"
	"tiktok/service"
	"time"

	"github.com/gin-gonic/gin"
)

// Response 结构体定义了响应的状态码和状态信息
type Response struct {
	StatusCode int32  `json:"status_code"`          // 状态码
	StatusMsg  string `json:"status_msg,omitempty"` // 状态信息
}

// FeedResponse 结构体定义了响应的状态码、状态信息和 FeedVideoFlow
type FeedResponse struct {
	Response
	FeedFlow service.FeedVideoFlow // 视频流
}

// Feed 函数处理获取视频流的请求
func Feed(c *gin.Context) {
	inputTime := ParseInputTime(c) // 解析请求中的时间参数
	userId := GetUserId(c)         // 获取请求中的用户 ID

	// 调用 service 层的 Feed 函数获取视频流
	vidoFlow, err := service.Feed(inputTime, userId)

	if err != nil { // 如果出现错误
		// 返回错误响应
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "获取失败",
			},
		})
		return
	}

	// 返回成功响应和视频流
	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "获取成功",
		},
		FeedFlow: *vidoFlow,
	})
}

// ParseInputTime 函数解析请求中的时间参数
func ParseInputTime(c *gin.Context) time.Time {
	inputTime := c.Query("latest_time") // 获取请求中的时间参数 latest_time

	if len(inputTime) != 0 { // 如果请求中有时间参数
		tempTime, _ := strconv.ParseInt(inputTime, 10, 64) // 将时间参数转换为 int64 类型

		return time.Unix(tempTime, 0) // 返回 Unix 时间
	}
	return time.Now() // 如果请求中没有时间参数，则返回当前时间
}

// GetUserId 函数获取请求中的用户 ID
func GetUserId(c *gin.Context) int64 {
	userId, _ := strconv.ParseInt(c.GetString("userId"), 10, 64) // 获取请求中的用户 ID
	return userId                                                // 返回用户 ID
}
