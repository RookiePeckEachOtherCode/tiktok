package service

import (
	"errors"
	"fmt"
	"math/rand"
	"tiktok/configs"
	"tiktok/dao"
	"tiktok/middleware/jwt"
	tiktokLog "tiktok/util/log"
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
		Avatar:        getRandomAvatar(),
	}

	// 生成token
	token, err := jwt.NewToken(userinfo.ID)

	if err != nil {
		return "", 0, fmt.Errorf("NewToken error:%v", err)
	}

	// 保存用户信息到数据库
	err = dao.AddUserInfo(&userinfo)

	if err != nil {
		return "", 0, fmt.Errorf("AddUserInfo error:%v", err)
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

func getRandomAvatar() string {
	//生成一个[1,8]的随机数
	randNum := rand.Intn(9)
	path := fmt.Sprintf("http://%v:%v/%v/%v.jpg", configs.LAN_IP, configs.GIN_PORT, configs.AVATAR_SAVE_PATH, randNum)
	return path
}
