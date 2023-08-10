package controller

import (
	"net/http"
	"tiktok/dao"
	"tiktok/model"
	"tiktok/util"

	"github.com/gin-gonic/gin"
)

type GetUserInfoResponse struct {
	Response model.Response // 用户鉴权token
	USerInfo dao.UserInfo   `json:"user"` // 用户id
}

func GetUserInfo(c *gin.Context) {
	util.PrintLog("在调用GetUserInfoById方法")
	_userid, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户id获取失败",
		})
		return
	}
	USerId, okk := _userid.(int64)
	if !okk {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "断言失败",
		})
		return
	}
	USerInfo, err := dao.GetUserInfoById(USerId)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户信息获取失败",
		})
		return
	}
	c.JSON(http.StatusOK, GetUserInfoResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "用户信息获取成功",
		},
		USerInfo: *USerInfo,
	})
}
