package service

import (
	"log"
	"tiktok/dao"
	"tiktok/util"
)

type CommentListResponse struct {
	dao.Response
	CommentList []*dao.Comment `json:"comment_list"`
}

func PostCommentService(vid, uid int64, context string) (*dao.Comment, error) {
	if ctx, err := util.FilterDirty(context); err != nil {
		return nil, err
	} else {
		context = ctx
	}

	log.Println(context)
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
	comment, err := dao.FindCommentByCommentId(commentId)
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
