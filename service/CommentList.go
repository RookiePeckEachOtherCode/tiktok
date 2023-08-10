package service

import (
	"tiktok/dao"
	"tiktok/model"
)

type ComListRes struct {
	model.Response
	CommentList []*dao.Comment `json:"comment_list"`
}

func HandleComListQuery(vid int64) (res *ComListRes, err error) {
	clist, err := dao.GetCommentList(vid)
	if err != nil {
		return nil, err
	}
	res = &ComListRes{}
	res.CommentList = clist
	res.StatusCode = 0
	res.StatusMsg = "获取列表成功"
	return res, nil
}
