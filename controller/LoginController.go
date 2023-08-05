package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/model"
	"tiktok/service"
)

type Gets struct {
	Response model.Response
	Token    *string `json:"token"`   // 用户鉴权token
	UserID   *int64  `json:"user_id"` // 用户id
}
type LoginMessge struct { //Query 参数
	Password string `json:"password"`
	Username string `json:"username"`
}

func login(c *gin.Context) { //处理登录请求
	var mes LoginMessge
	err := c.ShouldBindJSON(&mes) //获取登录信息
	if err != nil {
		c.JSON(http.StatusBadRequest, Gets{
			Response: model.Response{
				StatusMsg:  "获取登录信息失败",
				StatusCode: 1,
			},
			Token:  nil,
			UserID: nil,
		})
		return
	}
	res, err := service.Handlelogin(mes)
	if err != nil {
		c.JSON(http.StatusBadRequest, Gets{
			Response: model.Response{
				StatusMsg:  "获取登录请求错误",
				StatusCode: 1,
			},
			Token:  nil,
			UserID: nil,
		})
		return
	}
	c.JSON(http.StatusOK, res) //
}
