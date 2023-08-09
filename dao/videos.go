package dao

import (
	"errors"
	"tiktok/configs"
	"time"

	"gorm.io/gorm"
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
func GetVideoListByLastTime(lastTime time.Time) ([]*Video, error) {
	videos := make([]*Video, 0, configs.MAX_VIDEO_CNT)

	err := DB.Model(&Video{}).Where("created_at<?", lastTime).Order("created_at ASC").Limit(configs.MAX_VIDEO_CNT).Select([]string{"id", "user_info_id", "play_url", "cover_url", "favorite_count", "comment_count", "is_favorite", "title", "created_at", "updated_at"}).Find(&videos).Error

	return videos, err
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
func FavoriteVedio(v *Video, act int64) error { //更新点赞数
	if act == 1 {
		v.FavoriteCount++
	} else if act == 0 {
		v.FavoriteCount--
	}
	err := DB.Save(v).Error
	return err
}

//	func VideoFavPlus(userId, videoId int64) error {
//		return DB.Transaction(func(tx *gorm.DB) error {
//			if err := tx.Exec("UPDATE videos SET favorite_count=favorite_count+1 WHERE id = ?", videoId).Error; err != nil {
//				return err
//			}
//			if err := tx.Exec("INSERT INTO `user_favor_videos` (`user_info_id`,`video_id`) VALUES (?,?)", userId, videoId).Error; err != nil {
//				return err
//			}
//			return nil
//		})
//	}
func VideoFavPlus(userId, videoId int64) error {

	tx := DB.Begin()

	// 视频点赞数 +1
	if err := tx.Model(&Video{}).Where("id = ?", videoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count + 1")).Error; err != nil {
		tx.Rollback()
		return err
	}

	userFavoriteVideo := UserFavoriteVideo{
		UserID:  userId,
		VideoID: videoId,
	}

	// 添加用户点赞视频记录
	if err := tx.Create(&userFavoriteVideo).Error; err != nil {
		tx.Rollback()
		return err
	}
	//redis.SetFavorateState(userId, videoId, true)

	return tx.Commit().Error
}

// func VideoFavCancel(userId, videoId int64) error {
// 	return DB.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Exec("UPDATE videos SET favorite_count=favorite_count-1 WHERE id = ? AND favorite_count>0", videoId).Error; err != nil {
// 			return err
// 		}
// 		if err := tx.Exec("DELETE FROM `user_favor_videos`  WHERE `user_info_id` = ? AND `video_id` = ?", userId, videoId).Error; err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }

func VideoFavCancel(userId, videoId int64) error {

	tx := DB.Begin()

	// 视频点赞数-1
	if err := tx.Model(&Video{}).Where("id = ? AND favorite_count > 0", videoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除点赞记录
	if err := tx.Where("user_info_id = ? AND video_id = ?", userId, videoId).Delete(&UserFavoriteVideo{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//redis.SetFavorateState(userId, videoId, false)
	return tx.Commit().Error
}
