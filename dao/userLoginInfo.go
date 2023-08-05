package dao

import "tiktok/controller"

type UserLoginInfo struct {
	ID         int64  `gorm:"primaryKey;column:id" json:"id"`             // 用户登录ID
	UserInfoID int64  `gorm:"column:user_info_id" json:"user_info_id"`    // 用户信息ID
	Username   string `gorm:"primaryKey;column:username" json:"username"` // 用户名
	Password   string `gorm:"column:password;notnull" json:"password"`    // 密码
}

func JudgeUserPassword(a controller.LoginMessge) (*UserLoginInfo, error) { //根据传入的用户名和密码查找数据
	var info *UserLoginInfo
	result := DB.Where("username=? AND Password=?", a.Username, a.Password).Find(&info)
	if result.Error != nil {
		return nil, result.Error
	}
	return info, nil
}
