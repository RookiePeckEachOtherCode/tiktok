package model

type UserLogin struct {
	ID         int64    `gorm:"primaryKey;column:id"` // 用户登录ID
	UserInfoID int64    `gorm:"column:user_info_id"`  // 用户信息ID
	UserInfo   UserInfo `gorm:"joinForeignKey:user_info_id;foreignKey:id;references:UserInfoID"`
	Username   string   `gorm:"primaryKey;column:username"` // 用户名
	Password   string   `gorm:"column:password"`            // 密码
}
