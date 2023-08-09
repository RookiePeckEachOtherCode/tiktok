package controller

import (
	"errors"
	"fmt"
	"log"
	"tiktok/dao"
	"tiktok/model"

	"github.com/gin-gonic/gin"
)

type UserInfoResponse struct {
	Response model.Response
	User     *dao.UserInfo `json:"user"`
}

func GetUserInfo(c *gin.Context) {
	_userId, ok := c.Get("user_id")

	log.Printf("--------------------------user_id: %v------------------\n", _userId)
	log.Printf("--------------------------ok: %v------------------\n", ok)

	if !ok {
		log.Println("--------------获取用户信息失败----------------")
		c.JSON(200, UserInfoResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "获取用户信息失败",
			},
		})
		return
	}

	userInfo, err := ToGetUserInfo(_userId)

	if err != nil {
		log.Println("--------------获取用户信息失败----------------")
		c.JSON(200, UserInfoResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  fmt.Errorf("获取用户信息失败: %w", err).Error(),
			},
		})
		return
	}

	c.JSON(200, UserInfoResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		User: userInfo,
	})
}
func ToGetUserInfo(_userId interface{}) (*dao.UserInfo, error) {
	userId, ok := _userId.(int64)
	if !ok {
		log.Println("------------userId断言----------------")
		return nil, errors.New("userId类型断言失败")
	}
	userInfo, err := dao.GetUserInfoById(userId)

	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
