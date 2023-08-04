package dao

import "errors"

type UserInfo struct {
	ID              int64          `json:"id"`               // 用户id
	Name            string         `json:"name"`             // 用户名称
	FavoriteCount   int64          `json:"favorite_count"`   // 喜欢数
	FollowCount     int64          `json:"follow_count"`     // 关注总数
	FollowerCount   int64          `json:"follower_count"`   // 粉丝总数
	IsFollow        bool           `json:"is_follow"`        // true-已关注，false-未关注
	Signature       string         `json:"signature"`        // 个人简介
	TotalFavorited  string         `json:"total_favorited"`  // 获赞数量
	WorkCount       int64          `json:"work_count"`       // 作品数
	Avatar          string         `json:"avatar"`           // 用户头像
	BackgroundImage string         `json:"background_image"` // 用户个人页顶部大图
	UserLoginInfo   *UserLoginInfo // 用户登录信息
	Videos          []*Video       // 用户发布的视频
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
