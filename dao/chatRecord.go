package dao

import (
	"errors"
	"time"
)

type ChatRecord struct {
	ID         int64     `json:"id" gorm:"primaryKey;column:id"`
	Content    string    `json:"content" gorm:"column:content "`
	CreatedAt  time.Time `json:"created_at " gorm:"column:created_at"`
	UserInfo   *UserInfo `json:"user_info" gorm:"column:user_info"`
	TargetInfo *UserInfo `json:"target_info" gorm:"column:target_info"`
}

func CreateMes(uid int64, tid int64, mes string) error {
	userInfo, err := GetUserInfoById(uid)
	if err != nil {
		return errors.New("用户信息获取失败")
	}
	tuserinfo, err := GetUserInfoById(tid)
	if err != nil {
		return errors.New("对方信息获取失败")
	}
	Mes := ChatRecord{
		Content:    mes,
		CreatedAt:  time.Now(),
		UserInfo:   userInfo,
		TargetInfo: tuserinfo,
	}
	tx := DB.Begin()
	if err := tx.Model(&ChatRecord{}).Create(&Mes).Error; err != nil {
		DB.Rollback()
		return err
	}
	return nil
}
func MesList(uid int64, tid int64) (*[]ChatRecord, error) {
	tx := DB.Begin()

	var records []ChatRecord
	if err := tx.Where("user_info = ? AND target_info = ?", uid, tid).
		Or("user_info = ? AND target_info = ?", tid, uid).
		Find(&records).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &records, nil
}
