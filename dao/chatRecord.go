package dao

import "time"

type ChatRecord struct {
	ID         int64     `json:"id"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
	UserInfo   *UserInfo
	TargetInfo *UserInfo
}
