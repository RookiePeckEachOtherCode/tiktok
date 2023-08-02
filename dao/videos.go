package dao

import "time"

type Video struct {
	ID            int64       `gorm:"primaryKey;column:id"  json:"id"`             // 视频ID
	Title         string      `gorm:"column:title" json:"title"`                   // 视频标题
	UserInfoID    int64       `gorm:"column:user_info_id"   json:"-"`              // 用户信息ID
	Author        UserInfo    `json:"author" gorm:"-"`                             // 视频作者信息
	FavoriteCount int64       `gorm:"column:favorite_count" json:"favorite_count"` // 点赞数
	IsFavorite    bool        `gorm:"column:is_favorite"    json:"is_favorite"`    // 是否点赞true-已点赞，false-未点赞
	CommentCount  int64       `json:"comment_count" json:"comment_count" `         // 视频的评论总数
	PlayURL       string      `gorm:"column:play_url" json:"play_url"`             // 播放地址
	CoverURL      string      `gorm:"column:cover_url" json:"cover_url"`           // 封面地址
	Users         []*UserInfo `gorm:"many2many:user_favor_videos" json:"-"`        //点赞的用户
	Comments      []*Comment  `json:"-"`                                           //评论列表
	CreatedAt     time.Time   `json:"-"`                                           //上传时间
	UpdatedAt     time.Time   `json:"-"`
}

func GetVideoListByLastTime(limit int, lastTime time.Time) ([]*Video, error) {
	videos := make([]*Video, 0, limit)

	err := DB.Model(&Video{}).Where("created_at<?", lastTime).Order("created_at ASC").Limit(limit).Select([]string{"id", "user_info_id", "play_url", "cover_url", "favorite_count", "comment_count", "is_favorite", "title", "created_at", "updated_at"}).Find(&videos).Error

	return videos, err
}
