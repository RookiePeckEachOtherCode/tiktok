package controller

import (
	"fmt"
	"net/http"
	"tiktok/dao"
	"tiktok/service"
	"tiktok/util"

	"github.com/gin-gonic/gin"
)

type UserLoginResponse struct {
	dao.Response
	Token  *string `json:"token"`   // 用户鉴权token
	UserID *int64  `json:"user_id"` // 用户id
}

type UserRegisterResponse struct {
	dao.Response
	Token  string `json:"token"`   // 用户鉴权token
	UserID int64  `json:"user_id"` // 用户id
}

type GetUserInfoResponse struct {
	dao.Response               // 用户鉴权token
	dao.UserInfo `json:"user"` // 用户id
}

// 用户登录操作
func UserLoginController(c *gin.Context) { //处理登录请求
	userName := c.Query("username")
	_password, _ := c.Get("password")
	password, ok := _password.(string)
	if !ok {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "密码解析失败",
		})
		return
	}

	token, userID, err := service.UserLoginService(userName, password)

	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "登录失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		dao.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		&token,
		&userID,
	})

}

// 用户注册操作
func UserRegisterController(c *gin.Context) {
	_userName := c.Query("username")
	userName, _ := util.FilterDirty(_userName)

	_password, _ := c.Get("password")

	password, ok := _password.(string)

	if !ok {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "密码解析失败",
		})
	}

	token, Id, err := service.UserRegisterService(userName, password)

	// 注册失败
	if err != nil {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: dao.Response{
				StatusCode: 1,
				StatusMsg:  fmt.Sprintf("注册失败，%s", err.Error()),
			},
		})
		return
	}

	// 注册成功
	c.JSON(http.StatusOK, UserRegisterResponse{
		Response: dao.Response{
			StatusCode: 0,
			StatusMsg:  "注册成功",
		},
		Token:  token,
		UserID: Id,
	})
}

// 获取用户信息
func UserInfoController(c *gin.Context) {
	_userid, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "用户id获取失败",
		})
		return
	}
	USerId, okk := _userid.(int64)
	if !okk {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "断言失败",
		})
		return
	}
	USerInfo, err := dao.GetUserInfoById(USerId)
	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
		})
		return
	}

	c.JSON(http.StatusOK, GetUserInfoResponse{
		dao.Response{
			StatusCode: 0,
			StatusMsg:  "用户信息获取成功",
		},
		*USerInfo,
	})
}
