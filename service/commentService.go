package service

import (
	"tiktok/dao"
)

type CommentListResponse struct {
	dao.Response
	CommentList []*dao.Comment `json:"comment_list"`
}

func PostCommentService(vid, uid int64, context string) (*dao.Comment, error) {
	comment := &dao.Comment{
		UserInfoID: uid,
		VideoID:    vid,
		Content:    context,
	}
	if err := comment.PostComment(); err != nil {
		return nil, err
	}
	user, _ := dao.GetUserInfoById(uid)
	comment.User = *user
	return comment, nil
}

func DeleteCommentService(commentId string) error {
	comment, err := dao.FindComment(commentId)
	if err != nil {
		return err
	}
	if err := comment.DeleteComment(); err != nil {
		return err
	}
	return nil
}

func CommentListService(vid int64) (res *CommentListResponse, err error) {
	commentList, err := dao.GetCommentList(vid)
	if err != nil {
		return nil, err
	}

	//添加评论信息
	//因为数据库没有存储UserInfo,和CreatedDate
	for _, v := range commentList {
		userInfo, err := dao.GetUserInfoById(v.UserInfoID)
		if err != nil {
			return nil, err
		}
		v.User = *userInfo
		v.CreatedDate = v.CreatedAt.Format("01-02")
	}

	res = &CommentListResponse{
		Response: dao.Response{
			StatusCode: 0,
			StatusMsg:  "获取评论列表成功",
		},
		CommentList: commentList,
	}

	return res, nil
}
