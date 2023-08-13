package controller

import (
	"fmt"
	"log"
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
		log.Println("对方id获取失败")
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户id获取失败",
		})
		return
	}
	USerId, okk := _userid.(int64)
	if !okk {
		log.Println("断言失败")
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "断言失败",
		})
		return
	}
	USerInfo, err := dao.GetUserInfoById(USerId)
	if err != nil {
		log.Println("用户信息获取失败")
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户信息获取失败",
		})
		return
	}

	util.PrintLog(fmt.Sprintf("user_name: %v, user_id: %v,favorite_count:%v,have_favorite_count:%v,work_count:%v", USerInfo.Name, USerInfo.ID, USerInfo.TotalFavorited, USerInfo.TotalFavorited, USerInfo.WorkCount))
	c.JSON(http.StatusOK, GetUserInfoResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "用户信息获取成功",
		},
		USerInfo: *USerInfo,
	})
}
