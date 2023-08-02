package model

import "time"

type Video struct {
	ID            int64     `gorm:"primaryKey;column:id"` // 视频ID
	UserInfoID    int64     `gorm:"column:user_info_id"`  // 用户信息ID
	UserInfo      UserInfo  `gorm:"joinForeignKey:user_info_id;foreignKey:id;references:UserInfoID"`
	PlayURL       string    `gorm:"column:play_url"`       // 播放链接
	CoverURL      string    `gorm:"column:cover_url"`      // 封面链接
	FavoriteCount int64     `gorm:"column:favorite_count"` // 收藏数
	CommentCount  int64     `gorm:"column:comment_count"`  // 评论数
	IsFavorite    bool      `gorm:"column:is_favorite"`    // 是否收藏
	Title         string    `gorm:"column:title"`          // 视频标题
	CreatedAt     time.Time `gorm:"column:created_at"`     // 视频创建时间
	UpdatedAt     time.Time `gorm:"column:updated_at"`     // 视频更新时间
}
