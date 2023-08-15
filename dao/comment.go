package dao

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID          int64     `gorm:"primaryKey;column:id" json:"id"` // 评论ID
	UserInfoID  int64     `gorm:"column:user_info_id" json:"-"`   // 用户信息ID
	VideoID     int64     `gorm:"column:video_id" json:"-"`       // 视频ID
	User        UserInfo  `gorm:"-" json:"user"`
	Content     string    `gorm:"column:content" json:"content"` // 评论内容
	CreatedAt   time.Time `gorm:"column:created_at" json:"-"`    // 评论创建时间(存储到数据库)
	CreatedDate string    `gorm:"-" json:"create_date"`          // 评论创建时间(不存储到数据库, 用于返回给前端)
}

func FindComment(cid string) (*Comment, error) { //通过评论id查询评论
	var comment Comment
	err := DB.Where("id=?", cid).Find(&comment).Error
	return &comment, err
}

func (com *Comment) PostComment() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&com).Error; err != nil {
			return err
		}
		if err := tx.Model(&Video{}).Where("id = ?", com.VideoID).UpdateColumn("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
			return err
		}
		return nil
	})
}

func (com *Comment) DeleteComment() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&com).Error; err != nil {
			return err
		}
		if err := tx.Model(&Video{}).Where("id = ? AND comment_count > 0", com.VideoID).UpdateColumn("comment_count", gorm.Expr("comment_count - 1")).Error; err != nil {
			return err
		}
		return nil
	})
}
