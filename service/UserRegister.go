package service

import (
	"errors"
	"fmt"
	"tiktok/configs"
	"tiktok/dao"
	"tiktok/middleware/jwt"
)

// Register 注册
func Register(name, password string) (string, int64, error) {
	// 校验用户名和密码
	if err := RegisterCheck(name, password); err != nil {
		return "", 0, fmt.Errorf("密码校验失败:%v", err)
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
func RegisterCheck(name, password string) error {
	if len(name) == 0 {
		return errors.New("用户名不能为空")
	}

	if len(password) == 0 {
		return errors.New("密码不能为空")
	}

	if len(name) > configs.MAX_NAME_LEN {
		return errors.New("用户名长度不能超过32个字符")
	}

	if len(password) > configs.MAX_PASSWORD_LEN {
		return errors.New("密码长度不能超过32个字符")
	}

	if dao.IsExistUserLoginInfoByName(name) {
		return errors.New("该用户名已被注册")
	}

	return nil
}
