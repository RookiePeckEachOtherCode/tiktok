package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _UserFavorVideoMgr struct {
	*_BaseMgr
}

// UserFavorVideoMgr open func
func UserFavorVideoMgr(db *gorm.DB) *_UserFavorVideoMgr {
	if db == nil {
		panic(fmt.Errorf("UserFavorVideoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserFavorVideoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_favor_Video"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_UserFavorVideoMgr) Debug() *_UserFavorVideoMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserFavorVideoMgr) GetTableName() string {
	return "user_favor_Video"
}

// Reset 重置gorm会话
func (obj *_UserFavorVideoMgr) Reset() *_UserFavorVideoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UserFavorVideoMgr) Get() (result UserFavorVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFavorVideo{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("Video").Where("id = ?", result.VideoID).Find(&result.Video).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_UserFavorVideoMgr) Gets() (results []*UserFavorVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFavorVideo{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Video").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserFavorVideoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UserFavorVideo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUserInfoID user_info_id获取 用户信息ID
func (obj *_UserFavorVideoMgr) WithUserInfoID(userInfoID int64) Option {
	return optionFunc(func(o *options) { o.query["user_info_id"] = userInfoID })
}

// WithVideoID video_id获取 视频ID
func (obj *_UserFavorVideoMgr) WithVideoID(videoID int64) Option {
	return optionFunc(func(o *options) { o.query["video_id"] = videoID })
}

// GetByOption 功能选项模式获取
func (obj *_UserFavorVideoMgr) GetByOption(opts ...Option) (result UserFavorVideo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserFavorVideo{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("Video").Where("id = ?", result.VideoID).Find(&result.Video).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserFavorVideoMgr) GetByOptions(opts ...Option) (results []*UserFavorVideo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserFavorVideo{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Video").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromUserInfoID 通过user_info_id获取内容 用户信息ID
func (obj *_UserFavorVideoMgr) GetFromUserInfoID(userInfoID int64) (results []*UserFavorVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFavorVideo{}).Where("`user_info_id` = ?", userInfoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Video").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserInfoID 批量查找 用户信息ID
func (obj *_UserFavorVideoMgr) GetBatchFromUserInfoID(userInfoIDs []int64) (results []*UserFavorVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFavorVideo{}).Where("`user_info_id` IN (?)", userInfoIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Video").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVideoID 通过video_id获取内容 视频ID
func (obj *_UserFavorVideoMgr) GetFromVideoID(videoID int64) (results []*UserFavorVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFavorVideo{}).Where("`video_id` = ?", videoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Video").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVideoID 批量查找 视频ID
func (obj *_UserFavorVideoMgr) GetBatchFromVideoID(videoIDs []int64) (results []*UserFavorVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFavorVideo{}).Where("`video_id` IN (?)", videoIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Video").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserFavorVideoMgr) FetchByPrimaryKey(userInfoID int64, videoID int64) (result UserFavorVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFavorVideo{}).Where("`user_info_id` = ? AND `video_id` = ?", userInfoID, videoID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("Video").Where("id = ?", result.VideoID).Find(&result.Video).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkUserFavorVideoVideo  获取多个内容
func (obj *_UserFavorVideoMgr) FetchIndexByFkUserFavorVideoVideo(videoID int64) (results []*UserFavorVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserFavorVideo{}).Where("`video_id` = ?", videoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("Video").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
