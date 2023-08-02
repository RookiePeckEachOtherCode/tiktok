package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _CommentsMgr struct {
	*_BaseMgr
}

// CommentsMgr open func
func CommentsMgr(db *gorm.DB) *_CommentsMgr {
	if db == nil {
		panic(fmt.Errorf("CommentsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CommentsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("comments"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_CommentsMgr) Debug() *_CommentsMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CommentsMgr) GetTableName() string {
	return "comments"
}

// Reset 重置gorm会话
func (obj *_CommentsMgr) Reset() *_CommentsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CommentsMgr) Get() (result Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("videos").Where("id = ?", result.VideoID).Find(&result.Video).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_CommentsMgr) Gets() (results []*Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("videos").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CommentsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Comments{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 评论ID
func (obj *_CommentsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserInfoID user_info_id获取 用户信息ID
func (obj *_CommentsMgr) WithUserInfoID(userInfoID int64) Option {
	return optionFunc(func(o *options) { o.query["user_info_id"] = userInfoID })
}

// WithVideoID video_id获取 视频ID
func (obj *_CommentsMgr) WithVideoID(videoID int64) Option {
	return optionFunc(func(o *options) { o.query["video_id"] = videoID })
}

// WithContent content获取 评论内容
func (obj *_CommentsMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithCreatedAt created_at获取 评论创建时间
func (obj *_CommentsMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// GetByOption 功能选项模式获取
func (obj *_CommentsMgr) GetByOption(opts ...Option) (result Comments, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("videos").Where("id = ?", result.VideoID).Find(&result.Video).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CommentsMgr) GetByOptions(opts ...Option) (results []*Comments, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("videos").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 评论ID
func (obj *_CommentsMgr) GetFromID(id int64) (result Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("videos").Where("id = ?", result.VideoID).Find(&result.Video).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找 评论ID
func (obj *_CommentsMgr) GetBatchFromID(ids []int64) (results []*Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("videos").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUserInfoID 通过user_info_id获取内容 用户信息ID
func (obj *_CommentsMgr) GetFromUserInfoID(userInfoID int64) (results []*Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where("`user_info_id` = ?", userInfoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("videos").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserInfoID 批量查找 用户信息ID
func (obj *_CommentsMgr) GetBatchFromUserInfoID(userInfoIDs []int64) (results []*Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where("`user_info_id` IN (?)", userInfoIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("videos").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVideoID 通过video_id获取内容 视频ID
func (obj *_CommentsMgr) GetFromVideoID(videoID int64) (results []*Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where("`video_id` = ?", videoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("videos").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVideoID 批量查找 视频ID
func (obj *_CommentsMgr) GetBatchFromVideoID(videoIDs []int64) (results []*Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where("`video_id` IN (?)", videoIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("videos").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromContent 通过content获取内容 评论内容
func (obj *_CommentsMgr) GetFromContent(content string) (results []*Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where("`content` = ?", content).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("videos").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromContent 批量查找 评论内容
func (obj *_CommentsMgr) GetBatchFromContent(contents []string) (results []*Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where("`content` IN (?)", contents).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("videos").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedAt 通过created_at获取内容 评论创建时间
func (obj *_CommentsMgr) GetFromCreatedAt(createdAt time.Time) (results []*Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where("`created_at` = ?", createdAt).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("videos").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedAt 批量查找 评论创建时间
func (obj *_CommentsMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where("`created_at` IN (?)", createdAts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("videos").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
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
func (obj *_CommentsMgr) FetchByPrimaryKey(id int64) (result Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("videos").Where("id = ?", result.VideoID).Find(&result.Video).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkUserInfosComments  获取多个内容
func (obj *_CommentsMgr) FetchIndexByFkUserInfosComments(userInfoID int64) (results []*Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where("`user_info_id` = ?", userInfoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("videos").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByFkVideosComments  获取多个内容
func (obj *_CommentsMgr) FetchIndexByFkVideosComments(videoID int64) (results []*Comments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Comments{}).Where("`video_id` = ?", videoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("videos").Where("id = ?", results[i].VideoID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
