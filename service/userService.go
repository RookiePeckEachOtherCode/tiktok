package service

import (
	"errors"
	"fmt"
	"tiktok/dao"
	"tiktok/middleware/jwt"
	tiktokLog "tiktok/util/log"
	"tiktok/util/oss"
)

type UserInfoResponse struct {
	Response dao.Response // 用户鉴权token
	USerInfo dao.UserInfo `json:"user_info"` // 用户id
}

func UserLoginService(userName, password string) (string, int64, error) {
	if err := userLoginCheck(userName); err != nil {
		return "", 0, err
	}

	userId, err := dao.JudgeUserPassword(userName, password)

	if err != nil {
		return "", 0, err
	}

	token, err := jwt.NewToken(userId)

	if err != nil {
		return "", 0, fmt.Errorf("生成token失败: %v", err)
	}
	return token, userId, nil
}

func UserRegisterService(name, password string) (string, int64, error) {

	if err := registerCheck(name); err != nil {
		return "", 0, err
	}
	// 保存用户登录信息
	userLogin := dao.UserLogin{
		Username: name,
		Password: password,
	}
	// 保存用户登录信息
	userinfo := dao.UserInfo{
		UserLoginInfo: &userLogin,
		Name:          name,
		Avatar:        oss.GetRandomAvatar(),
	}

	// 保存用户信息到数据库
	if err := dao.AddUserInfo(&userinfo); err != nil {
		tiktokLog.Error("AddUserInfo error:%v", err)
		return "", 0, fmt.Errorf("AddUserInfo error:%v", err)
	}

	// 生成token
	token, err := jwt.NewToken(userinfo.ID)

	if err != nil {
		return "", 0, fmt.Errorf("NewToken error:%v", err)
	}

	return token, userinfo.ID, nil
}

// Check 校验用户名和密码

func userLoginCheck(name string) error {
	if !dao.CheckIsExistByName(name) {
		tiktokLog.Error("该用户不存在: userName: ", name)
		return errors.New("该用户不存在")
	}
	return nil
}

func registerCheck(name string) error {
	if dao.IsExistUserLoginInfoByName(name) {
		tiktokLog.Error("该用户名已被注册: userName: ", name)
		return errors.New("该用户名已被注册")
	}
	return nil
}
