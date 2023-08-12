package dao

import (
	"errors"
	"time"
)

type ChatRecord struct {
	ID        int64     `json:"id" gorm:"primaryKey;column:id"`
	UserID    int64     `json:"user_id" gorm:"column:user_id"`
	TargetID  int64     `json:"target_id" gorm:"column:target_id"`
	Content   string    `json:"content" gorm:"column:content"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	User      *UserInfo `json:"user,omitempty" gorm:"foreignkey:UserID"`
	Target    *UserInfo `json:"target,omitempty" gorm:"foreignkey:TargetID"`
}

func CreateMes(uid int64, tid int64, mes string) error {
	userInfo, err := GetUserInfoById(uid)
	if err != nil {
		return errors.New("用户信息获取失败")
	}
	tuserinfo, err := GetUserInfoById(tid)
	if err != nil {
		errors.New("对方信息获取失败")
	}
	Mes := ChatRecord{
		Content:   mes,
		CreatedAt: time.Now(),
		UserID:    uid,
		TargetID:  tid,
		User:      userInfo,
		Target:    tuserinfo,
	}
	tx := DB.Begin()
	if err := tx.Model(&ChatRecord{}).Create(&Mes).Error; err != nil {
		DB.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func MesList(uid int64, tid int64) (*[]ChatRecord, error) {
	tx := DB.Begin()

	var records []ChatRecord
	if err := tx.Where("user_id = ? AND target_id = ?", uid, tid).
		Or("user_id = ? AND target_id = ?", tid, uid).
		Find(&records).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &records, nil
}
