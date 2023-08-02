package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _UserInfosMgr struct {
	*_BaseMgr
}

// UserInfosMgr open func
func UserInfosMgr(db *gorm.DB) *_UserInfosMgr {
	if db == nil {
		panic(fmt.Errorf("UserInfosMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserInfosMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_infos"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_UserInfosMgr) Debug() *_UserInfosMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserInfosMgr) GetTableName() string {
	return "user_infos"
}

// Reset 重置gorm会话
func (obj *_UserInfosMgr) Reset() *_UserInfosMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_UserInfosMgr) Get() (result UserInfos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_UserInfosMgr) Gets() (results []*UserInfos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserInfosMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 用户信息ID
func (obj *_UserInfosMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 用户名
func (obj *_UserInfosMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithFollowCount follow_count获取 关注数
func (obj *_UserInfosMgr) WithFollowCount(followCount int64) Option {
	return optionFunc(func(o *options) { o.query["follow_count"] = followCount })
}

// WithFollowerCount follower_count获取 粉丝数
func (obj *_UserInfosMgr) WithFollowerCount(followerCount int64) Option {
	return optionFunc(func(o *options) { o.query["follower_count"] = followerCount })
}

// WithIsFollow is_follow获取 是否关注
func (obj *_UserInfosMgr) WithIsFollow(isFollow bool) Option {
	return optionFunc(func(o *options) { o.query["is_follow"] = isFollow })
}


// GetByOption 功能选项模式获取
func (obj *_UserInfosMgr) GetByOption(opts ...Option) (result UserInfos, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserInfosMgr) GetByOptions(opts ...Option) (results []*UserInfos, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Where(options.query).Find(&results).Error
	
	return
}



//////////////////////////enume case ////////////////////////////////////////////


// GetFromID 通过id获取内容 用户信息ID 
func (obj *_UserInfosMgr)  GetFromID(id int64) (result UserInfos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Where("`id` = ?", id).First(&result).Error
	
	return
}

// GetBatchFromID 批量查找 用户信息ID
func (obj *_UserInfosMgr) GetBatchFromID(ids []int64) (results []*UserInfos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Where("`id` IN (?)", ids).Find(&results).Error
	
	return
}
 
// GetFromName 通过name获取内容 用户名 
func (obj *_UserInfosMgr) GetFromName(name string) (results []*UserInfos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Where("`name` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 用户名
func (obj *_UserInfosMgr) GetBatchFromName(names []string) (results []*UserInfos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Where("`name` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromFollowCount 通过follow_count获取内容 关注数 
func (obj *_UserInfosMgr) GetFromFollowCount(followCount int64) (results []*UserInfos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Where("`follow_count` = ?", followCount).Find(&results).Error
	
	return
}

// GetBatchFromFollowCount 批量查找 关注数
func (obj *_UserInfosMgr) GetBatchFromFollowCount(followCounts []int64) (results []*UserInfos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Where("`follow_count` IN (?)", followCounts).Find(&results).Error
	
	return
}
 
// GetFromFollowerCount 通过follower_count获取内容 粉丝数 
func (obj *_UserInfosMgr) GetFromFollowerCount(followerCount int64) (results []*UserInfos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Where("`follower_count` = ?", followerCount).Find(&results).Error
	
	return
}

// GetBatchFromFollowerCount 批量查找 粉丝数
func (obj *_UserInfosMgr) GetBatchFromFollowerCount(followerCounts []int64) (results []*UserInfos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Where("`follower_count` IN (?)", followerCounts).Find(&results).Error
	
	return
}
 
// GetFromIsFollow 通过is_follow获取内容 是否关注 
func (obj *_UserInfosMgr) GetFromIsFollow(isFollow bool) (results []*UserInfos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Where("`is_follow` = ?", isFollow).Find(&results).Error
	
	return
}

// GetBatchFromIsFollow 批量查找 是否关注
func (obj *_UserInfosMgr) GetBatchFromIsFollow(isFollows []bool) (results []*UserInfos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Where("`is_follow` IN (?)", isFollows).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 
 // FetchByPrimaryKey primary or index 获取唯一内容
 func (obj *_UserInfosMgr) FetchByPrimaryKey(id int64 ) (result UserInfos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UserInfos{}).Where("`id` = ?", id).First(&result).Error
	
	return
}
 

 

	

