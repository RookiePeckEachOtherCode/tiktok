package dao

type UserLogin struct {
	ID         int64  `gorm:"primaryKey;column:id"` // 用户登录ID
	UserInfoID int64  // 用户信息ID
	Username   string `gorm:"primaryKey;column:username"` // 用户名
	Password   string `gorm:"column:password;notnull"`    // 密码
}

func GetUserLoginInfoByName(name string) (UserLogin, error) {
	userLoginInfo := UserLogin{}
	DB.Where("username=?", name).First(&userLoginInfo)

	return userLoginInfo, nil
}

func IsExistUserLoginInfoByName(name string) bool {
	var userLogin UserLogin
	DB.Where("username=?", name).First(&userLogin)

	return userLogin.ID != 0
}
