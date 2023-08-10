package dao

import "time"

type Comment struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"id"` // 评论ID
	UserInfoID int64     `gorm:"column:user_info_id" json:"-"`   // 用户信息ID
	VideoID    int64     `gorm:"column:video_id" json:"-"`       // 视频ID
	User       UserInfo  `gorm:"-" json:"user"`
	Content    string    `gorm:"column:content" json:"content"` // 评论内容
	CreatedAt  time.Time `gorm:"column:created_at" json:"-"`    // 评论创建时间
	CreateDate string    `gorm:"-" json:"create_date"`
}

func FindComment(cid string) (*Comment, error) { //通过评论id查询评论
	var comment Comment
	err := DB.Where("id=?", cid).Find(&comment).Error
	return &comment, err
}
