package model

// Comments [...]
type Comments struct {
	Comment
}

// TableName get sql table name.获取数据库表名
func (m *Comments) TableName() string {
	return "comments"
}

// UserFavorVideos [...]
type UserFavorVideos struct {
	UserFavorVideo
}

// TableName get sql table name.获取数据库表名
func (m *UserFavorVideos) TableName() string {
	return "user_favor_videos"
}

// UserInfos [...]
type UserInfos struct {
	UserInfo
}

// TableName get sql table name.获取数据库表名
func (m *UserInfos) TableName() string {
	return "user_infos"
}

// UserLogins [...]
type UserLogins struct {
	UserLogin
}

// TableName get sql table name.获取数据库表名
func (m *UserLogins) TableName() string {
	return "user_logins"
}

// UserRelations [...]
type UserRelations struct {
	UserRelation
}

// TableName get sql table name.获取数据库表名
func (m *UserRelations) TableName() string {
	return "user_relations"
}

// Videos [...]
type Videos struct {
	Video
}

// TableName get sql table name.获取数据库表名
func (m *Videos) TableName() string {
	return "videos"
}
