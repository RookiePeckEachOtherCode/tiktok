package controller

import (
	"net/http"
	"strconv"
	"tiktok/dao"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

type CommentActResponse struct {
	dao.Response
	Comment dao.Comment `json:"comment"`
}

// 用户评论操作
func CommentActionController(c *gin.Context) {
	_userid, _ := c.Get("user_id")
	userid, ok1 := _userid.(int64)
	if !ok1 {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取用户信息失败",
		})
	}
	_videoId := c.Query("video_id")
	videoId, err := strconv.ParseInt(_videoId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "视频id获取失败: " + err.Error(),
		})
	}
	_actionType := c.Query("action_type")
	actionType, err1 := strconv.ParseInt(_actionType, 10, 64)

	if err1 != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "操作类型获取失败: " + err.Error(),
		})
	}

	switch actionType {
	case 1:
		_comment := c.Query("comment_text")
		comment, err := service.PostCommentService(videoId, userid, _comment)
		if err != nil {
			c.JSON(http.StatusOK, dao.Response{
				StatusCode: 1,
				StatusMsg:  "评论发布失败: " + err.Error(),
			})
			return
		}

		comment.CreatedDate = comment.CreatedAt.Format("01-02")

		c.JSON(http.StatusOK, CommentActResponse{
			dao.Response{
				StatusCode: 0,
				StatusMsg:  "发布评论成功",
			},
			*comment,
		})
	case 2:
		commentId := c.Query("comment_id")
		err := service.DeleteCommentService(commentId)
		if err != nil {
			c.JSON(http.StatusOK, dao.Response{
				StatusCode: 1,
				StatusMsg:  "删除评论失败: " + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 0,
			StatusMsg:  "删除评论成功",
		})
	}
}

// 获取评论列表
func CommentListController(c *gin.Context) {
	Vid := c.Query("video_id")
	vid, err := strconv.ParseInt(Vid, 9, 64)
	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 0,
			StatusMsg:  "无法获取视频id: " + err.Error(),
		})
	}
	res, err := service.CommentListService(vid)
	if err != nil {
		c.JSON(http.StatusBadRequest, service.FavoriteListReponse{
			Response: dao.Response{
				StatusCode: 0,
				StatusMsg:  "获取评论列表失败: " + err.Error(),
			},
		})
	} else {
		c.JSON(http.StatusOK, res)
	}
}
