package service

import (
	"tiktok/controller"
	"tiktok/dao"
	"tiktok/middleware/jwt"
	"tiktok/model"
)

type Gets struct {
	Response model.Response
	Token    *string `json:"token"`   // 用户鉴权token
	UserID   *int64  `json:"user_id"` // 用户id
}

func Handlelogin(messge controller.LoginMessge) (*Gets, error) {
	userLoginInfo, err := dao.JudgeUserPassword(messge) //验证登录信息，并返回登录信息
	if err != nil {
		return &Gets{
			Response: model.Response(struct {
				StatusCode int32
				StatusMsg  string
			}{StatusCode: 0, StatusMsg: "登录失败，用户名或密码错误"}),
			Token:  nil,
			UserID: nil,
		}, err
	}
	userInfo, err := dao.GetUserInfoById(userLoginInfo.UserInfoID) //获取用户信息
	if err != nil {
		return &Gets{
			Response: model.Response(struct {
				StatusCode int32
				StatusMsg  string
			}{StatusCode: 0, StatusMsg: "登录失败，无法找到用户信息"}),
			Token:  nil,
			UserID: nil,
		}, err
	}
	newToken, err := jwt.NewToken(userInfo) //使用用户信息生成一个jwt验证
	if err != nil {
		return &Gets{
			Response: model.Response(struct {
				StatusCode int32
				StatusMsg  string
			}{StatusCode: 1, StatusMsg: "登录失败，错误的jwt生成请求"}),
			Token:  nil,
			UserID: nil,
		}, err
	}
	return &Gets{
		Response: model.Response(struct {
			StatusCode int32
			StatusMsg  string
		}{StatusCode: 0, StatusMsg: "登录成功"}),
		Token:  &newToken,
		UserID: &userInfo.ID,
	}, nil
}
