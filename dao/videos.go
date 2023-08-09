package dao

import (
	"github.com/pkg/errors"
	"tiktok/configs"
	"time"
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
func FavoriteVideo(v *Video, act int64, uid int64) error { //更新喜欢操作，包含对用户喜欢列表
	userInfo, err := GetUserInfoById(uid)
	if err != nil {
		return errors.Wrap(err, "获取用户信息失败")
	}
	if act == 1 {
		v.FavoriteCount++                   //喜欢数++
		v.Users = append(v.Users, userInfo) //添加喜欢用户信息
		userInfo.FavorVideos = append(userInfo.FavorVideos, v)
		err := DB.Save(v).Error
		if err != nil {
			return errors.Wrap(err, "保存视频信息失败")
		}
		err1 := DB.Save(userInfo).Error
		if err != nil {
			return errors.Wrap(err, "保存用户信息失败")
		}
		if err1 != nil {
			return errors.Wrap(err, "保存视频信息失败")
		}

	} else if act == 2 {
		v.FavoriteCount--
		var duser UserInfo
		result := DB.First(&duser, uid) //查找对应的用户对象
		if result.Error != nil {
			return errors.New("无法查询到用户")
		}
		for i, u := range v.Users {
			if u.ID == duser.ID {
				v.Users = append(v.Users[:i], v.Users[i+1:]...) //从当前视频移除对应用户
				break
			}
		}
		for i, vid := range duser.FavorVideos {
			if v.ID == vid.ID {
				duser.FavorVideos = append(duser.FavorVideos[:i], duser.FavorVideos[i+1:]...) //从对象用户中移除视频
				break
			}
		}
		err := DB.Save(duser).Error
		if err != nil {
			return errors.Wrap(err, "保存用户信息失败")
		}
	}
	err = DB.Save(v).Error
	if err != nil {
		return errors.Wrap(err, "保存视频信息失败")
	}
	DB.Model(v).Association("Users").Replace(v.Users) //刷新数据库，使移除喜欢的视频不会回滚到喜欢列表中
	return nil
}

// 	userFavoriteVideo := UserFavoriteVideo{
// 		UserID:  userId,
// 		VideoID: videoId,
// 	}

// 	// 添加用户点赞视频记录
// 	if err := tx.Create(&userFavoriteVideo).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	return tx.Commit().Error
// }

// func VideoFavCancel(userId, videoId int64) error {

// 	tx := DB.Begin()

// 	// 视频点赞数-1
// 	if err := tx.Model(&Video{}).Where("id = ? AND favorite_count > 0", videoId).UpdateColumn("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	// 删除点赞记录
// 	if err := tx.Where("user_info_id = ? AND video_id = ?", userId, videoId).Delete(&UserFavoriteVideo{}).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	//redis.SetFavorateState(userId, videoId, false)
// 	return tx.Commit().Error
// }
