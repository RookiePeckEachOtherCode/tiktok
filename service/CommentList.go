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

	//添加评论信息
	//因为数据库没有存储UserInfo,和CreatedDate
	for _, v := range clist {
		userInfo, err := dao.GetUserInfoById(v.UserInfoID)
		if err != nil {
			return nil, err
		}
		v.User = *userInfo
		v.CreatedDate = v.CreatedAt.Format("01-02")
	}

	res = &ComListRes{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取评论列表成功",
		},
		CommentList: clist,
	}

	return res, nil
}
