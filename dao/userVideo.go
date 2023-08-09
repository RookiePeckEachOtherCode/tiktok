package dao

type UserFavoriteVideo struct {
	UserID  int64 `gorm:"column:user_info_id"`
	VideoID int64 `gorm:"column:video_id"`
}

func (u UserFavoriteVideo) TableName() string {
	return "user_favor_videos"
}

func GetIsFavorite(userId, videoId int64) bool {
	var count int64
	DB.Model(&UserFavoriteVideo{}).Where("user_info_id = ? AND video_id = ?", userId, videoId).Count(&count)
	return count > 0
}
