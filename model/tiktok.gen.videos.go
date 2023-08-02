package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _VideoMgr struct {
	*_BaseMgr
}

// VideoMgr open func
func VideoMgr(db *gorm.DB) *_VideoMgr {
	if db == nil {
		panic(fmt.Errorf("VideoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_VideoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("Video"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_VideoMgr) Debug() *_VideoMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_VideoMgr) GetTableName() string {
	return "Video"
}

// Reset 重置gorm会话
func (obj *_VideoMgr) Reset() *_VideoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_VideoMgr) Get() (result Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_VideoMgr) Gets() (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_VideoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Video{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 视频ID
func (obj *_VideoMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserInfoID user_info_id获取 用户信息ID
func (obj *_VideoMgr) WithUserInfoID(userInfoID int64) Option {
	return optionFunc(func(o *options) { o.query["user_info_id"] = userInfoID })
}

// WithPlayURL play_url获取 播放链接
func (obj *_VideoMgr) WithPlayURL(playURL string) Option {
	return optionFunc(func(o *options) { o.query["play_url"] = playURL })
}

// WithCoverURL cover_url获取 封面链接
func (obj *_VideoMgr) WithCoverURL(coverURL string) Option {
	return optionFunc(func(o *options) { o.query["cover_url"] = coverURL })
}

// WithFavoriteCount favorite_count获取 收藏数
func (obj *_VideoMgr) WithFavoriteCount(favoriteCount int64) Option {
	return optionFunc(func(o *options) { o.query["favorite_count"] = favoriteCount })
}

// WithCommentCount comment_count获取 评论数
func (obj *_VideoMgr) WithCommentCount(commentCount int64) Option {
	return optionFunc(func(o *options) { o.query["comment_count"] = commentCount })
}

// WithIsFavorite is_favorite获取 是否收藏
func (obj *_VideoMgr) WithIsFavorite(isFavorite bool) Option {
	return optionFunc(func(o *options) { o.query["is_favorite"] = isFavorite })
}

// WithTitle title获取 视频标题
func (obj *_VideoMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithCreatedAt created_at获取 视频创建时间
func (obj *_VideoMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 视频更新时间
func (obj *_VideoMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// GetByOption 功能选项模式获取
func (obj *_VideoMgr) GetByOption(opts ...Option) (result Video, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_VideoMgr) GetByOptions(opts ...Option) (results []*Video, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 视频ID
func (obj *_VideoMgr) GetFromID(id int64) (result Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找 视频ID
func (obj *_VideoMgr) GetBatchFromID(ids []int64) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUserInfoID 通过user_info_id获取内容 用户信息ID
func (obj *_VideoMgr) GetFromUserInfoID(userInfoID int64) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`user_info_id` = ?", userInfoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserInfoID 批量查找 用户信息ID
func (obj *_VideoMgr) GetBatchFromUserInfoID(userInfoIDs []int64) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`user_info_id` IN (?)", userInfoIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPlayURL 通过play_url获取内容 播放链接
func (obj *_VideoMgr) GetFromPlayURL(playURL string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`play_url` = ?", playURL).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPlayURL 批量查找 播放链接
func (obj *_VideoMgr) GetBatchFromPlayURL(playURLs []string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`play_url` IN (?)", playURLs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCoverURL 通过cover_url获取内容 封面链接
func (obj *_VideoMgr) GetFromCoverURL(coverURL string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`cover_url` = ?", coverURL).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCoverURL 批量查找 封面链接
func (obj *_VideoMgr) GetBatchFromCoverURL(coverURLs []string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`cover_url` IN (?)", coverURLs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromFavoriteCount 通过favorite_count获取内容 收藏数
func (obj *_VideoMgr) GetFromFavoriteCount(favoriteCount int64) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`favorite_count` = ?", favoriteCount).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromFavoriteCount 批量查找 收藏数
func (obj *_VideoMgr) GetBatchFromFavoriteCount(favoriteCounts []int64) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`favorite_count` IN (?)", favoriteCounts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCommentCount 通过comment_count获取内容 评论数
func (obj *_VideoMgr) GetFromCommentCount(commentCount int64) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`comment_count` = ?", commentCount).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCommentCount 批量查找 评论数
func (obj *_VideoMgr) GetBatchFromCommentCount(commentCounts []int64) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`comment_count` IN (?)", commentCounts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIsFavorite 通过is_favorite获取内容 是否收藏
func (obj *_VideoMgr) GetFromIsFavorite(isFavorite bool) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`is_favorite` = ?", isFavorite).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIsFavorite 批量查找 是否收藏
func (obj *_VideoMgr) GetBatchFromIsFavorite(isFavorites []bool) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`is_favorite` IN (?)", isFavorites).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTitle 通过title获取内容 视频标题
func (obj *_VideoMgr) GetFromTitle(title string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`title` = ?", title).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTitle 批量查找 视频标题
func (obj *_VideoMgr) GetBatchFromTitle(titles []string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`title` IN (?)", titles).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedAt 通过created_at获取内容 视频创建时间
func (obj *_VideoMgr) GetFromCreatedAt(createdAt time.Time) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`created_at` = ?", createdAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedAt 批量查找 视频创建时间
func (obj *_VideoMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedAt 通过updated_at获取内容 视频更新时间
func (obj *_VideoMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`updated_at` = ?", updatedAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedAt 批量查找 视频更新时间
func (obj *_VideoMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`updated_at` IN (?)", updatedAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
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
func (obj *_VideoMgr) FetchByPrimaryKey(id int64) (result Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkUserInfoVideo  获取多个内容
func (obj *_VideoMgr) FetchIndexByFkUserInfoVideo(userInfoID int64) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`user_info_id` = ?", userInfoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
