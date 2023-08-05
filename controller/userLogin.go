package controller

import (
	"net/http"
	"tiktok/model"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

type Gets struct {
	Response model.Response
	Token    *string `json:"token"`   // 用户鉴权token
	UserID   *int64  `json:"user_id"` // 用户id
}

func UserLogin(c *gin.Context) { //处理登录请求
	userName := c.Query("username")
	password := c.Query("password")

	token, userID, err := service.Handlelogin(userName, password)

	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Gets{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "登录成功",
		},
		Token:  &token,
		UserID: &userID,
	})

}
