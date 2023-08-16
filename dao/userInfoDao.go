package dao

import (
	"errors"
	"log"
	"tiktok/middleware/redis"

	"gorm.io/gorm"
)

type UserInfo struct {
	ID              int64       `json:"id" gorm:"id,omitempty"`                         //用户id
	Name            string      `json:"name" gorm:"name,omitempty"`                     //用户名称
	FollowCount     int64       `json:"follow_count" gorm:"follow_count,omitempty"`     //关注数
	FollowerCount   int64       `json:"follower_count" gorm:"follower_count,omitempty"` //粉丝总数
	IsFollow        bool        `json:"is_follow" gorm:"is_follow,omitempty"`           //当前登录用户是否关注该用户,true-已关注,false-未关注
	UserLoginInfo   *UserLogin  `json:"-"`                                              //用户与登录信息的一对一
	Videos          []*Video    `json:"-"`                                              //用户与视频的一对多
	Follows         []*UserInfo `json:"-" gorm:"many2many:user_relations;"`             //用户与关注用户之间的多对多
	FavorVideos     []*Video    `json:"-" gorm:"many2many:user_favor_videos;"`          //用户与喜欢视频之间的多对多
	Comments        []*Comment  `json:"-"`                                              //用户与评论的一对多
	TotalFavorite   int64       `json:"total_favorited" gorm:"-"`                       //用户获赞数
	WorkCount       int64       `json:"work_count" gorm:"-"`                            //作品数
	Avatar          string      `json:"avatar" gorm:"avatar,omitempty"`                 //头像
	FavoriteCount   int64       `json:"favorite_count" gorm:"-"`                        //用户喜欢的视频数
	BackgroundImage string      `json:"background_image" gorm:"-"`                      //背景图片
	Signature       string      `json:"signature" gorm:"-"`                             //个性签名
}
type Friend struct {
	UserInfo
	Message string `json:"message"` //消息
	MsgType int8   `json:"msgType"` //消息类型 message信息的类型，0=>请求用户接受信息，1=>当前请求用户发送的信息
}

// GetUserInfoById 根据用户id获取用户信息
func GetUserInfoById(userId int64) (*UserInfo, error) {
	var userInfo UserInfo
	DB.Where("id=?", userId).First(&userInfo)
	if userInfo.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	userInfo.Signature = "这个人很懒，什么都没写"
	userInfo.FavoriteCount = redis.New(redis.FAVORITE).GetUserFavoriteCount(userId)
	userInfo.TotalFavorite = redis.New(redis.LIKED).GetUserReceivedLikeCount(userId)
	return &userInfo, nil
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
	redis.New(redis.FAVORITE).UpdateFavoriteState(u.ID, video.ID, true)
	redis.New(redis.LIKED).UpdateUserReceivedLikeCount(video.UserInfoID, true)

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
	redis.New(redis.FAVORITE).UpdateFavoriteState(u.ID, video.ID, false)
	redis.New(redis.LIKED).UpdateUserReceivedLikeCount(video.UserInfoID, false)
	return tx.Commit().Error
}

// 通过id获取用户喜欢的视频列表
func GetFavoriteList(userId int64) ([]*Video, error) {
	var userInfo UserInfo
	err := DB.Preload("FavorVideos").First(&userInfo, "id=?", userId).Error
	if err != nil {
		return nil, err
	} else {
		return userInfo.FavorVideos, nil
	}
}

// 用户获赞数+1
func (u *UserInfo) PlusFavCount() error {
	tx := DB.Begin()
	if err := tx.Model(u).UpdateColumn("total_favorited", gorm.Expr("total_favorited + 1")).Error; err != nil {
		tx.Rollback()
		return err
	}
	redis.New(redis.LIKED).UpdateUserReceivedLikeCount(u.ID, true)
	return tx.Commit().Error
}

func (u *UserInfo) FollowAct(tu *UserInfo) error {
	tx := DB.Begin()
	tx.Model(u)
	if err := tx.Model(u).Association("Follows").Append(tu); err != nil { //将对方添加到关注列表
		tx.Rollback()
		return err
	}
	if err := tx.Model(u).UpdateColumn("follow_count", gorm.Expr("follow_count +1")).Error; err != nil { //自己关注数+1
		tx.Rollback()
		return err
	}
	if err := tx.Model(tu).UpdateColumn("follower_count", gorm.Expr("follower_count +1")).Error; err != nil { //对方粉丝数+1
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 取关
func (u *UserInfo) UnFollowAct(tu *UserInfo) error {
	tx := DB.Begin()
	if err := tx.Model(u).Association("Follows").Delete(tu); err != nil { //从自己的关注列表移除
		tx.Rollback()
		return err
	}
	if err := tx.Model(u).UpdateColumn("follow_count", gorm.Expr("follow_count -1")).Error; err != nil { //自己关注数-1
		tx.Rollback()
		return err
	}
	if err := tx.Model(tu).UpdateColumn("follower_count", gorm.Expr("follower_count -1")).Error; err != nil { //对方粉丝数-1
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// 获取用户关注列表
func GetFloList(uid int64) ([]*UserInfo, error) {
	var userInfo UserInfo
	err := DB.Preload("Follows").First(&userInfo, "id=?", uid).Error
	if err != nil {
		return nil, err
	} else {
		return userInfo.Follows, nil
	}
}

func GetFollowerList(uid int64) ([]*UserInfo, error) {
	var follower []*UserInfo
	if err := DB.Model(&UserInfo{}).Where("id in (select user_info_id from user_relations where follow_id = ?)", uid).Find(&follower).Error; err != nil {
		return nil, err
	}
	return follower, nil
}

// GetUserRelation 判断两个用户之间是否存在关注关系
func GetUserRelation(uid, tid int64) bool { //uid是否关注tid
	tx := DB.Begin()
	var userInfo UserInfo
	if err := tx.Preload("Follows").First(&userInfo, uid).Error; err != nil {
		tx.Rollback()
		return false
	}
	count := tx.Model(&userInfo).Where("id=?", tid).Association("Follows").Count()
	return count > 0
}
func GetMutualFriendListById(userId int64) ([]*UserInfo, error) {
	var mutualFriendList []int64
	var Friends []*UserInfo

	if err := DB.Raw("SELECT a.follow_id FROM user_relations a JOIN user_relations b ON a.user_info_id = b.follow_id AND a.follow_id = b.user_info_id WHERE a.user_info_id = ?", userId).Scan(&mutualFriendList).Error; err != nil {
		return nil, err
	}

	for _, id := range mutualFriendList {
		userInfo, err := GetUserInfoById(id)
		if err != nil {
			continue //忽略错误
		}
		Friends = append(Friends, userInfo)
	}
	if len(Friends) == 0 {
		return nil, errors.New("没有共同好友")
	}
	if len(Friends) < len(mutualFriendList) {
		return Friends, errors.New("部分好友不存在")
	}
	return Friends, nil
}

// messageType消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息

func GetNewestMessageByUserIdAndToUserID(userId int64, toUserId int64) (string, int8, error) {
	message := ChatRecord{}
	result := DB.Where("user_id = ? AND to_user_id = ? ", userId, toUserId).
		Or("user_id = ? AND to_user_id = ? ", toUserId, userId).
		Order("created_time desc").
		Limit(1).Find(&message)
	if result.Error != nil {
		log.Println("查询最新消息失败", result.Error.Error())
		return "", -1, result.Error
	}

	if userId == message.FromUserId {
		log.Println("当前用户发送的消息", message.Content)
		return message.Content, 1, nil
	} else {
		log.Println("当前用户接收的消息", message.Content)
		return message.Content, 0, nil
	}
}