package dao

import "time"

type Comment struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"id"` // 评论ID
	UserInfoID int64     `gorm:"column:user_info_id" json:"-"`   // 用户信息ID
	VideoID    int64     `gorm:"column:video_id" json:"-"`       // 视频ID
	Content    string    `gorm:"column:content" json:"content"`  // 评论内容
	CreatedAt  time.Time `gorm:"column:created_at" json:"-"`     // 评论创建时间
	CreateDate string    `gorm:"-" json:"create_date"`
	User       UserInfos `gorm:"joinForeignKey:user_info_id;foreignKey:id;references:UserInfoID" json:"user"`
}
