package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _UserLoginsMgr struct {
	*_BaseMgr
}

// UserLoginsMgr open func
func UserLoginsMgr(db *gorm.DB) *_UserLoginsMgr {
	if db == nil {
		panic(fmt.Errorf("UserLoginsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserLoginsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_logins"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_UserLoginsMgr) Debug() *_UserLoginsMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserLoginsMgr) GetTableName() string {
	return "user_logins"
}

// Reset 重置gorm会话
func (obj *_UserLoginsMgr) Reset() *_UserLoginsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UserLoginsMgr) Get() (result UserLogins, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserLogins{}).First(&result).Error
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
func (obj *_UserLoginsMgr) Gets() (results []*UserLogins, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserLogins{}).Find(&results).Error
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
func (obj *_UserLoginsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UserLogins{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 用户登录ID
func (obj *_UserLoginsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserInfoID user_info_id获取 用户信息ID
func (obj *_UserLoginsMgr) WithUserInfoID(userInfoID int64) Option {
	return optionFunc(func(o *options) { o.query["user_info_id"] = userInfoID })
}

// WithUsername username获取 用户名
func (obj *_UserLoginsMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取 密码
func (obj *_UserLoginsMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// GetByOption 功能选项模式获取
func (obj *_UserLoginsMgr) GetByOption(opts ...Option) (result UserLogins, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserLogins{}).Where(options.query).First(&result).Error
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
func (obj *_UserLoginsMgr) GetByOptions(opts ...Option) (results []*UserLogins, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserLogins{}).Where(options.query).Find(&results).Error
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

// GetFromID 通过id获取内容 用户登录ID
func (obj *_UserLoginsMgr) GetFromID(id int64) (results []*UserLogins, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserLogins{}).Where("`id` = ?", id).Find(&results).Error
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

// GetBatchFromID 批量查找 用户登录ID
func (obj *_UserLoginsMgr) GetBatchFromID(ids []int64) (results []*UserLogins, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserLogins{}).Where("`id` IN (?)", ids).Find(&results).Error
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
func (obj *_UserLoginsMgr) GetFromUserInfoID(userInfoID int64) (results []*UserLogins, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserLogins{}).Where("`user_info_id` = ?", userInfoID).Find(&results).Error
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
func (obj *_UserLoginsMgr) GetBatchFromUserInfoID(userInfoIDs []int64) (results []*UserLogins, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserLogins{}).Where("`user_info_id` IN (?)", userInfoIDs).Find(&results).Error
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

// GetFromUsername 通过username获取内容 用户名
func (obj *_UserLoginsMgr) GetFromUsername(username string) (results []*UserLogins, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserLogins{}).Where("`username` = ?", username).Find(&results).Error
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

// GetBatchFromUsername 批量查找 用户名
func (obj *_UserLoginsMgr) GetBatchFromUsername(usernames []string) (results []*UserLogins, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserLogins{}).Where("`username` IN (?)", usernames).Find(&results).Error
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

// GetFromPassword 通过password获取内容 密码
func (obj *_UserLoginsMgr) GetFromPassword(password string) (results []*UserLogins, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserLogins{}).Where("`password` = ?", password).Find(&results).Error
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

// GetBatchFromPassword 批量查找 密码
func (obj *_UserLoginsMgr) GetBatchFromPassword(passwords []string) (results []*UserLogins, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserLogins{}).Where("`password` IN (?)", passwords).Find(&results).Error
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
func (obj *_UserLoginsMgr) FetchByPrimaryKey(id int64, username string) (result UserLogins, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserLogins{}).Where("`id` = ? AND `username` = ?", id, username).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user_infos").Where("id = ?", result.UserInfoID).Find(&result.UserInfo).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkUserInfoUser  获取多个内容
func (obj *_UserLoginsMgr) FetchIndexByFkUserInfoUser(userInfoID int64) (results []*UserLogins, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserLogins{}).Where("`user_info_id` = ?", userInfoID).Find(&results).Error
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
