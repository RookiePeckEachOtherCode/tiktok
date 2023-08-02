package model

type UserFavorVideo struct {
	UserInfoID int64    `gorm:"primaryKey;column:user_info_id"` // 用户信息ID
	UserInfo   UserInfo `gorm:"joinForeignKey:user_info_id;foreignKey:id;references:UserInfoID"`
	VideoID    int64    `gorm:"primaryKey;column:video_id"` // 视频ID
	Video      Video    `gorm:"joinForeignKey:video_id;foreignKey:id;references:VideoID"`
}
