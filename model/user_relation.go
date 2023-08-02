package model

type UserRelation struct {
	UserInfoID int64    `gorm:"primaryKey;column:user_info_id"` // 用户信息ID
	UserInfo   UserInfo `gorm:"joinForeignKey:user_info_id;foreignKey:id;references:UserInfoID"`
	FollowID   int64    `gorm:"primaryKey;column:follow_id"` // 关注用户信息ID
}
