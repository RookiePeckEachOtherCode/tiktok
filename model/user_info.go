package model

type UserInfo struct {
	ID            int64  `gorm:"primaryKey;column:id"`  // 用户信息ID
	Name          string `gorm:"column:name"`           // 用户名
	FollowCount   int64  `gorm:"column:follow_count"`   // 关注数
	FollowerCount int64  `gorm:"column:follower_count"` // 粉丝数
	IsFollow      bool   `gorm:"column:is_follow"`      // 是否关注
}
