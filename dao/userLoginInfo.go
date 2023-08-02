package dao

type UserLoginInfo struct {
	ID         int64  `gorm:"primaryKey;column:id"`       // 用户登录ID
	UserInfoID int64  `gorm:"column:user_info_id"`        // 用户信息ID
	Username   string `gorm:"primaryKey;column:username"` // 用户名
	Password   string `gorm:"column:password;notnull"`    // 密码
}
