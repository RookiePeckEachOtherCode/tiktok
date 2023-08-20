package controller

import (
	"fmt"
	"log"
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
	Response dao.Response // 用户鉴权token
	USerInfo dao.UserInfo `json:"user"` // 用户id
}

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

func UserInfoController(c *gin.Context) {
	util.PrintLog("在调用GetUserInfoById方法")
	_userid, ok := c.Get("user_id")
	if !ok {
		log.Println("对方id获取失败")
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "用户id获取失败",
		})
		return
	}
	USerId, okk := _userid.(int64)
	if !okk {
		log.Println("断言失败")
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "断言失败",
		})
		return
	}
	USerInfo, err := dao.GetUserInfoById(USerId)
	if err != nil {
		log.Println("用户信息获取失败")
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
		})
		return
	}

	util.PrintLog(fmt.Sprintf("user_name: %v, user_id: %v,favorite_count:%v,have_favorite_count:%v,work_count:%v", USerInfo.Name, USerInfo.ID, USerInfo.TotalFavorite, USerInfo.TotalFavorite, USerInfo.WorkCount))
	c.JSON(http.StatusOK, GetUserInfoResponse{
		Response: dao.Response{
			StatusCode: 0,
			StatusMsg:  "用户信息获取成功",
		},
		USerInfo: *USerInfo,
	})
}
