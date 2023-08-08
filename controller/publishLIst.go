package controller

import (
	"net/http"
	"tiktok/dao"
	"tiktok/model"
	"tiktok/service"
	"tiktok/util"

	"github.com/gin-gonic/gin"
)

type PublishListResponse struct {
	Response  model.Response
	VideoList []dao.Video `json:"video_list"`
}

func PublishList(c *gin.Context) {
	_userId, _ := c.Get("user_id")
	userId, ok := _userId.(int64)
	//判断用户id类型是否正确
	if !ok {
		util.PrintLog("用户id类型错误")
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户id类型错误",
		})

	}
	//获取发布列表
	videoList, err := service.GetPublishList(userId)

	if err != nil {
		util.PrintLog("获取发布列表失败")
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, PublishListResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: *videoList,
	})

	util.PrintLog("获取发布列表成功")
}
