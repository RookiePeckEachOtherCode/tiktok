package dao

import (
	"tiktok/configs"
	"time"

	"github.com/pkg/errors"
)

type Video struct {
	ID            int64       `gorm:"primaryKey;column:id"  json:"id"`             // 视频ID
	Title         string      `gorm:"column:title" json:"title"`                   // 视频标题
	UserInfoID    int64       `gorm:"column:user_info_id"   json:"-"`              // 用户信息ID
	Author        UserInfo    `json:"author" gorm:"-"`                             // 视频作者信息
	FavoriteCount int64       `gorm:"column:favorite_count" json:"favorite_count"` // 点赞数
	IsFavorite    bool        `gorm:"column:is_favorite"    json:"is_favorite"`    // 是否点赞true-已点赞，false-未点赞
	CommentCount  int64       `json:"comment_count" `                              // 视频的评论总数
	PlayURL       string      `gorm:"column:play_url" json:"play_url"`             // 播放地址
	CoverURL      string      `gorm:"column:cover_url" json:"cover_url"`           // 封面地址
	Users         []*UserInfo `gorm:"many2many:user_favor_videos" json:"-"`        //点赞的用户
	Comments      []*Comment  `json:"-"`                                           //评论列表
	CreatedAt     time.Time   `json:"-"`                                           //上传时间
	UpdatedAt     time.Time   `json:"-"`
}

// GetVideoListByLastTime 根据上传时间获取视频列表
func GetVideoListByLastTime(lastTime time.Time) (*[]*Video, error) {
	var videos []*Video
	err := DB.Model(&Video{}).Where("created_at<?", lastTime).
		Order("created_at DESC").Limit(configs.MAX_VIDEO_CNT).
		Select([]string{"id", "user_info_id", "play_url", "cover_url", "favorite_count", "comment_count", "is_favorite", "title", "created_at", "updated_at"}).
		Find(&videos).Error
	return &videos, err
}

func NewVideo(v *Video) error { //上传视频
	if v == nil {
		return errors.New("[NewVideo] video为空")
	}
	return DB.Create(v).Error
}

func GetVideoListByUserId(userId int64) (*[]Video, error) { //通过用户id查询用户视频列表
	var videoList []Video

	err := DB.Where("user_info_id=?", userId).Select([]string{"id", "user_info_id", "play_url", "cover_url", "favorite_count", "comment_count", "is_favorite", "title"}).Find(&videoList).Error

	return &videoList, err
}
func FindVideoByVid(vid int64) (*Video, error) { //通过视频id查询视频
	var vd Video

	err := DB.Where("id=?", vid).Find(&vd).Error

	return &vd, err
}

// 通过视频id获取评论列表
func GetCommentList(id int64) ([]*Comment, error) {
	var comment []*Comment

	if err := DB.Model(&Comment{}).Where("video_id=?", id).Find(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}
