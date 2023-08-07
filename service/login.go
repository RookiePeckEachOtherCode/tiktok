// package service
package service

import (
	"errors"
	"fmt"
	"tiktok/dao"
	"tiktok/middleware/jwt"
)

func Handlelogin(userName, password string) (string, int64, error) {
	if err := check(userName); err != nil {
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

func check(name string) error {
	if !dao.CheckIsExistByName(name) {
		return errors.New("该用户不存在")
	}
	return nil
}
