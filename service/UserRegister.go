package service

import (
	"errors"
	"fmt"
	"tiktok/dao"
	"tiktok/middleware/jwt"
)

// Register 注册
func Register(name, password string) (string, int64, error) {
	if err := Check(name); err != nil {
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
func Check(name string) error {

	if dao.IsExistUserLoginInfoByName(name) {
		return errors.New("该用户名已被注册")
	}

	return nil
}
