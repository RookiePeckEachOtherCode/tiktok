// package dao
package dao

import (
	"errors"
)

// import "tiktok/controller"

type UserLogin struct {
	ID         int64  `gorm:"primaryKey;column:id"` // 用户登录ID
	UserInfoID int64  // 用户信息ID
	Username   string `gorm:"primaryKey;column:username"` // 用户名
	Password   string `gorm:"column:password;notnull"`    // 密码
}

func JudgeUserPassword(name, password string) (int64, error) { //根据传入的用户名和密码查找数据
	userlogInfo := UserLogin{}
	DB.Where("username=? and password=?", name, password).First(&userlogInfo)

	if userlogInfo.ID == 0 {
		return 0, errors.New("密码错误")
	}
	return userlogInfo.ID, nil
}
