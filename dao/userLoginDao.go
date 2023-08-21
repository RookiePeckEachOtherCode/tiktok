// package dao
package dao

import (
	"errors"
)

type UserLogin struct {
	ID         int64  `gorm:"primaryKey;column:id"` // 用户登录ID
	UserInfoID int64  // 用户信息ID
	Username   string `gorm:"primaryKey;column:username"` // 用户名
	Password   string `gorm:"column:password;notnull"`    // 密码
}

// JudgeUserPassword 判断用户密码是否正确，正确返回用户id，错误返回错误信息
func JudgeUserPassword(name, password string) (int64, error) {
	if !IsExistUserLoginInfoByName(name) {
		return 0, errors.New("该用户不存在")
	}

	userlogInfo := UserLogin{}
	DB.Where("username=? and password=?", name, password).First(&userlogInfo)

	if userlogInfo.ID == 0 {
		return 0, errors.New("密码错误")
	}
	return userlogInfo.ID, nil
}

// IsExistUserLoginInfoByName 通过用户名判断用户是否存在
func IsExistUserLoginInfoByName(name string) bool {
	var userLogin UserLogin
	DB.Where("username=?", name).First(&userLogin)

	return userLogin.ID != 0
}
