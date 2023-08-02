package model

import "time"

type Comment struct {
	ID         int64     `gorm:"primaryKey;column:id"` // 评论ID
	UserInfoID int64     `gorm:"column:user_info_id"`  // 用户信息ID
	UserInfo   UserInfo  `gorm:"joinForeignKey:user_info_id;foreignKey:id;references:UserInfoID"`
	VideoID    int64     `gorm:"column:video_id"` // 视频ID
	Video      Video     `gorm:"joinForeignKey:video_id;foreignKey:id;references:VideoID"`
	Content    string    `gorm:"column:content"`    // 评论内容
	CreatedAt  time.Time `gorm:"column:created_at"` // 评论创建时间
}
