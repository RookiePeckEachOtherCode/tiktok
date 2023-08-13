package controller

import (
	"net/http"
	"tiktok/model"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

type UserLoginResponse struct {
	model.Response
	Token  *string `json:"token"`   // 用户鉴权token
	UserID *int64  `json:"user_id"` // 用户id
}

func UserLogin(c *gin.Context) { //处理登录请求
	userName := c.Query("username")
	_password, _ := c.Get("password")
	password, ok := _password.(string)
	if !ok {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "密码解析失败",
		})
		return
	}

	token, userID, err := service.Handlelogin(userName, password)

	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		model.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		&token,
		&userID,
	})

}
