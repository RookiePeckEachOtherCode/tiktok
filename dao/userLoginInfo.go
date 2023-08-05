package dao

type UserLoginInfo struct {
	ID         int64 `gorm:"primary_key"` //用户登录ID
	UserInfoID int64
	Username   string `gorm:"primary_key"`      ////用户名
	Password   string `gorm:"size:200;notnull"` //密码
}
