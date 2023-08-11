package controller

import (
	"errors"
	"net/http"
	"strconv"
	"tiktok/middleware/jwt"
	"tiktok/model"
	"tiktok/service"
	"time"

	"github.com/gin-gonic/gin"
)

// FeedResponse 结构体定义了响应的状态码、状态信息和 FeedVideoFlow
type FeedResponse struct {
	model.Response
	*service.FeedVideoFlow
}

func Feed(c *gin.Context) {
	token, ok := c.GetQuery("token")

	var userId int64 = 0
	if ok {
		if _userId, err := AlreadlyLogin(token); err == nil {
			userId = _userId
		} else {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
		}
	}

	var latestTime time.Time
	_latestTime, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)

	if err != nil {
		latestTime = time.Unix(0, _latestTime*1e6)
	}

	feedVideos, err := service.Feed(userId, latestTime)

	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		FeedVideoFlow: feedVideos,
	})
}

func AlreadlyLogin(token string) (int64, error) {
	claims, ok := jwt.ParseToken(token)
	if ok {
		if claims.ExpiresAt < time.Now().Unix() {
			return 0, errors.New("登陆过期,请重新登陆")
		}
		return claims.UserId, nil
	}
	return 0, errors.New("token解析失败")
}

func ParseLatestTime(c *gin.Context) (time.Time, error) {
	latestTime, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)

	if err != nil {
		return time.Time{}, errors.New("latest_time解析失败")
	}
	return time.Unix(0, latestTime*int64(time.Millisecond)), nil
}
