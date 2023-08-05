package dao

import "errors"

type UserInfo struct {
  ID            int64       `json:"id" gorm:"id,omitempty"`                         //用户id
	Name          string      `json:"name" gorm:"name,omitempty"`                     //用户名称
	FollowCount   int64       `json:"follow_count" gorm:"follow_count,omitempty"`     //关注数
	FollowerCount int64       `json:"follower_count" gorm:"follower_count,omitempty"` //粉丝总数
	IsFollow      bool        `json:"is_follow" gorm:"is_follow,omitempty"`           //true-已关注，false-未关注
	UserLoginInfo *UserLogin  `json:"-"`                                              //用户与登录信息的一对一
	Videos        []*Video    `json:"-"`                                              //用户与视频的一对多
	Follows       []*UserInfo `json:"-" gorm:"many2many:user_relations;"`             //用户与关注用户之间的多对多
	FavorVideos   []*Video    `json:"-" gorm:"many2many:user_favor_videos;"`          //用户与喜欢视频之间的多对多
	Comments      []*Comment  `json:"-"`                                              //用户与评论的一对
}

// GetUserInfoById 根据用户id获取用户信息
func GetUserInfoById(userId int64) (UserInfo, error) {
	var userInfo UserInfo

	DB.Where("id=?", userId).Select([]string{"id", "name", "follow_count", "follower_count", "is_follow"}).First(&userInfo)

	if userInfo.ID == 0 {
		return UserInfo{}, errors.New("该用户不存在")
	}
	return userInfo, nil
}
// AddUserInfo 保存用户信息到数据库
func AddUserInfo(user *UserInfo) error {
	if user == nil {
		return errors.New("user is nil")
	}
	return DB.Create(user).Error
}
//通过名字查询用户是否存在
func CheckIsExistByName(name string) bool {
	var userInfo UserInfo
	DB.Where("name=?", name).Select([]string{"id"}).First(&userInfo)
	if userInfo.ID == 0 {
		return false
	}
	return true
}