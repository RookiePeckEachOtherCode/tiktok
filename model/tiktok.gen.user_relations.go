package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _UserRelationsMgr struct {
	*_BaseMgr
}

// UserRelationsMgr open func
func UserRelationsMgr(db *gorm.DB) *_UserRelationsMgr {
	if db == nil {
		panic(fmt.Errorf("UserRelationsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserRelationsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_relations"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_UserRelationsMgr) Debug() *_UserRelationsMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserRelationsMgr) GetTableName() string {
	return "user_relations"
}

// Reset 重置gorm会话
func (obj *_UserRelationsMgr) Reset() *_UserRelationsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UserRelationsMgr) Get() (result UserRelations, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserRelations{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.FollowID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_UserRelationsMgr) Gets() (results []*UserRelations, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserRelations{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].FollowID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserRelationsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UserRelations{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUserInfoID user_info_id获取 用户信息ID
func (obj *_UserRelationsMgr) WithUserInfoID(userInfoID int64) Option {
	return optionFunc(func(o *options) { o.query["user_info_id"] = userInfoID })
}

// WithFollowID follow_id获取 关注用户信息ID
func (obj *_UserRelationsMgr) WithFollowID(followID int64) Option {
	return optionFunc(func(o *options) { o.query["follow_id"] = followID })
}

// GetByOption 功能选项模式获取
func (obj *_UserRelationsMgr) GetByOption(opts ...Option) (result UserRelations, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserRelations{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.FollowID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserRelationsMgr) GetByOptions(opts ...Option) (results []*UserRelations, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserRelations{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].FollowID).Find(&results[i].UserInfo).Error; err != nil { //
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
func (obj *_UserRelationsMgr) GetFromUserInfoID(userInfoID int64) (results []*UserRelations, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserRelations{}).Where("`user_info_id` = ?", userInfoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].FollowID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserInfoID 批量查找 用户信息ID
func (obj *_UserRelationsMgr) GetBatchFromUserInfoID(userInfoIDs []int64) (results []*UserRelations, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserRelations{}).Where("`user_info_id` IN (?)", userInfoIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].FollowID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromFollowID 通过follow_id获取内容 关注用户信息ID
func (obj *_UserRelationsMgr) GetFromFollowID(followID int64) (results []*UserRelations, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserRelations{}).Where("`follow_id` = ?", followID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].FollowID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromFollowID 批量查找 关注用户信息ID
func (obj *_UserRelationsMgr) GetBatchFromFollowID(followIDs []int64) (results []*UserRelations, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserRelations{}).Where("`follow_id` IN (?)", followIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].FollowID).Find(&results[i].UserInfo).Error; err != nil { //
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
func (obj *_UserRelationsMgr) FetchByPrimaryKey(userInfoID int64, followID int64) (result UserRelations, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserRelations{}).Where("`user_info_id` = ? AND `follow_id` = ?", userInfoID, followID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.FollowID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkUserRelationsFollows  获取多个内容
func (obj *_UserRelationsMgr) FetchIndexByFkUserRelationsFollows(followID int64) (results []*UserRelations, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserRelations{}).Where("`follow_id` = ?", followID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].UserInfoID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("user_infos").Where("id = ?", results[i].FollowID).Find(&results[i].UserInfo).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
