package service

import (
	"tiktok/dao"
)

type MesListRes struct {
	StatusCode  string    `json:"status_code"`
	StatusMsg   string    `json:"status_msg"`
	MessageList []Message `json:"message_list"`
}

type Message struct {
	ID         int64  `json:"id"`
	ToUserID   int64  `json:"to_user_id"`
	FromUserID int64  `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

func HandleMesList(uid int64, tid int64) (*MesListRes, error) {
	list, err := dao.MesList(uid, tid)
	if err != nil {
		return nil, err
	}

	messageList := make([]Message, len(*list))
	for i, record := range *list {
		messageList[i] = Message{
			ID:         record.ID,
			ToUserID:   record.TargetInfo.ID,
			FromUserID: record.UserInfo.ID,
			Content:    record.Content,
			CreateTime: record.CreatedAt.Unix(),
		}
	}

	res := &MesListRes{
		StatusCode:  "200",
		StatusMsg:   "成功",
		MessageList: messageList,
	}

	return res, nil
}
