package controller

import (
	"fmt"
	"net/http"
	"tiktok/model"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

// UserRegisterRespons 用户注册响应
type UserRegisterRespons struct {
	model.Response
	Token  string `json:"token"`   // 用户鉴权token
	UserID int64  `json:"user_id"` // 用户id
}

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	userName := c.Query("username")
	_password, _ := c.Get("password")

	password, ok := _password.(string)

	if !ok {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "密码解析失败",
		})
	}

	token, Id, err := service.Register(userName, password)

	// 注册失败
	if err != nil {
		c.JSON(http.StatusOK, UserRegisterRespons{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  fmt.Sprintf("注册失败，%s", err.Error()),
			},
		})
		return
	}

	// 注册成功
	c.JSON(http.StatusOK, UserRegisterRespons{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "注册成功",
		},
		Token:  token,
		UserID: Id,
	})
}
