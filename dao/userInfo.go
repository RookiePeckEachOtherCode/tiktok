package dao

import (
	"errors"

	"gorm.io/gorm"
)

type UserInfo struct {
	ID            int64       `json:"id" gorm:"id,omitempty"`                         //用户id
	Name          string      `json:"name" gorm:"name,omitempty"`                     //用户名称
	FollowCount   int64       `json:"follow_count" gorm:"follow_count,omitempty"`     //关注数
	FollowerCount int64       `json:"follower_count" gorm:"follower_count,omitempty"` //粉丝总数
	IsFollow      bool        `json:"is_follow" gorm:"is_follow,omitempty"`           //当前登录用户是否关注该用户,true-已关注,false-未关注
	UserLoginInfo *UserLogin  `json:"-"`                                              //用户与登录信息的一对一
	Videos        []*Video    `json:"-"`                                              //用户与视频的一对多
	Follows       []*UserInfo `json:"-" gorm:"many2many:user_relations;"`             //用户与关注用户之间的多对多
	FavorVideos   []*Video    `json:"-" gorm:"many2many:user_favor_videos;"`          //用户与喜欢视频之间的多对多
	Comments      []*Comment  `json:"-"`                                              //用户与评论的一对
	FavorCount    int64       `json:"favorite_count" gorm:"-"`
	WorkCount     int64       `json:"work_count" gorm:"-"`
}

// GetUserInfoById 根据用户id获取用户信息
func GetUserInfoById(userId int64) (*UserInfo, error) {
	var userInfo UserInfo

	DB.Where("id=?", userId).Select([]string{"id", "name", "follow_count", "follower_count", "is_follow"}).First(&userInfo)

	if userInfo.ID == 0 {
		return nil, errors.New("该用户不存在")
	}

	userInfo.GetFavorCount()
	userInfo.GetWorkCount()

	return &userInfo, nil
}
func (u *UserInfo) GetFavorCount() {
	u.FavorCount = DB.Model(u).Association("user_favor_videos").Count()
}

func (u *UserInfo) GetWorkCount() {
	u.WorkCount = DB.Model(u).Association("videos").Count()
}

// AddUserInfo 保存用户信息到数据库
func AddUserInfo(user *UserInfo) error {
	if user == nil {
		return errors.New("user is nil")
	}
	return DB.Create(user).Error
}

// 通过名字查询用户是否存在
func CheckIsExistByName(name string) bool {
	var userInfo UserInfo
	DB.Where("name=?", name).Select([]string{"id"}).First(&userInfo)

	return userInfo.ID != 0
}

// 通过id查询用户是否存在
func CheckIsExistByID(id int64) bool {
	var userInfo UserInfo
	DB.Where("id=?", id).Select([]string{"id"}).First(&userInfo)

	return userInfo.ID != 0
}

func (u *UserInfo) GetIsFavorite(videoId int64) bool {
	count := DB.Model(u).Where("video_id = ?", videoId).Association("FavorVideos").Count()
	return count > 0
}

// FavoriteVideo 给视频点赞
func (u *UserInfo) ToFavoriteVideo(video *Video) error {

	tx := DB.Begin()

	if err := tx.Model(video).UpdateColumn("favorite_count", gorm.Expr("favorite_count + 1")).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(u).Association("FavorVideos").Append(video); err != nil {
		tx.Rollback()
		return err
	}
	//redis.SetFavorateState(userId, videoId, true)
	return tx.Commit().Error
}

// CancelFavorite 取消点赞
func (u *UserInfo) ToCancelFavorite(video *Video) error {

	tx := DB.Begin()

	if err := tx.Model(video).Where("favorite_count > 0").UpdateColumn("favorite_count", gorm.Expr("favorite_count - 1")).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(u).Association("FavorVideos").Delete(video); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
// 通过id获取用户喜欢的视频列表
func GetFavList(id int64) ([]*Video, error) {
	var uinfo UserInfo
	err := DB.Preload("FavorVideos").First(&uinfo, "id=?", id).Error
	if err != nil {
		return nil, err
	} else {
		return uinfo.FavorVideos, nil
	}
}
}
