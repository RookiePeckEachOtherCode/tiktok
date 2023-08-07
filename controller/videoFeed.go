package controller

import (
	"errors"
	"net/http"
	"strconv"
	"tiktok/dao"
	"tiktok/middleware/jwt"
	"tiktok/model"
	"tiktok/service"
	"time"

	"github.com/gin-gonic/gin"
)

// FeedResponse 结构体定义了响应的状态码、状态信息和 FeedVideoFlow
type FeedResponse struct {
	model.Response
	NextTime  int64        `json:"next_time"`  //发布最早的时间，作为下次请求时的latest_time
	VideoList []*dao.Video `json:"video_list"` //视频列表
}

func Feed(c *gin.Context) {
	token, ok := c.GetQuery("token")
	var userId int64 = 0

	if ok { //已经登陆
		var err error
		userId, err = AlreadlyLogin(token)
		if err != nil {
			c.JSON(200, FeedResponse{
				Response: model.Response{
					StatusCode: 400,
					StatusMsg:  err.Error(),
				},
			})
			return
		}
	}

	latestTime, err := ParseLatestTime(c)
	if err != nil {
		c.JSON(200, FeedResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	videos, nextTime, err := service.Feed(userId, latestTime)

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
		NextTime:  nextTime,
		VideoList: videos,
	})
}

func AlreadlyLogin(token string) (int64, error) {
	claims, ok := jwt.ParseToken(token)
	if ok {
		if claims.ExpiresAt > time.Now().Unix() {
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
