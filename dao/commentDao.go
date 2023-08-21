package dao

import (
	"fmt"
	tiktokLog "tiktok/util/log"
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

// FindCommentByCommentId 通过评论id查询评论
func FindCommentByCommentId(cid string) (*Comment, error) {
	var comment Comment
	err := DB.Where("id=?", cid).Find(&comment).Error
	if err != nil {
		tiktokLog.Error(fmt.Sprintf("通过评论id查询评论失败, cid: %s, Error: %v", cid, err))
	}
	return &comment, err
}

// PostComment 发布评论
func (com *Comment) PostComment() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&com).Error; err != nil {
			tiktokLog.Error("保存评论到数据库失败: ", err.Error(), "comment: ", com)
			return err
		}
		if err := tx.Model(&Video{}).Where("id = ?", com.VideoID).UpdateColumn("comment_count", gorm.Expr("comment_count + 1")).Error; err != nil {
			tiktokLog.Error("更新视频评论数失败: ", err.Error(), "videoId: ", com.VideoID)
			return err
		}
		return nil
	})
}

// 删除评论
func (com *Comment) DeleteComment() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&com).Error; err != nil {
			tiktokLog.Error("删除评论失败: ", err.Error(), "comment: ", com)
			return err
		}
		if err := tx.Model(&Video{}).Where("id = ? AND comment_count > 0", com.VideoID).UpdateColumn("comment_count", gorm.Expr("comment_count - 1")).Error; err != nil {
			tiktokLog.Error("更新视频评论数失败: ", err.Error(), "videoId: ", com.VideoID)
			return err
		}
		return nil
	})
}
