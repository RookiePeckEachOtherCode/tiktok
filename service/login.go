// package service
package service

import (
	"errors"
	"fmt"
	"tiktok/dao"
	"tiktok/middleware/jwt"
)

func Handlelogin(userName, password string) (string, int64, error) {
	if err := check(userName, password); err != nil {
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

func check(name, password string) error {
	if name == "" {
		return errors.New("用户名不能为空")
	}
	if password == "" {
		return errors.New("密码不能为空")
	}
	if len(name) > 32 {
		return errors.New("用户名不能超过32位")
	}
	if len(password) > 32 {
		return errors.New("密码不能超过32位")
	}
	if !dao.CheckIsExistByName(name) {
		return errors.New("该用户不存在")
	}

	return nil
}
