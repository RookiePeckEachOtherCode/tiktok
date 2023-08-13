package controller

import (
	"net/http"
	"strconv"
	"tiktok/dao"
	"tiktok/model"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

type CommentActResponse struct {
	model.Response
	Comment dao.Comment `json:"comment"`
}

func CommentAct(c *gin.Context) {
	_userid, _ := c.Get("user_id")
	userid, ok1 := _userid.(int64)
	if !ok1 {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取用户信息失败",
		})
	}

	videoid := c.Query("video_id")
	Videoid, err := strconv.ParseInt(videoid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "视频id获取失败",
		})
	}
	actinotype := c.Query("action_type")
	Actiontype, err1 := strconv.ParseInt(actinotype, 10, 64)

	if err1 != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "操作类型获取失败",
		})
	}

	switch Actiontype {
	case 1:
		Commenttext := c.Query("comment_text")
		comment, err := service.PostComment(Videoid, userid, Commenttext)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: 1,
				StatusMsg:  "评论发布失败",
			})
			return
		}

		comment.CreatedDate = comment.CreatedAt.Format("01-02")

		c.JSON(http.StatusOK, CommentActResponse{
			model.Response{
				StatusCode: 0,
				StatusMsg:  "发布评论成功",
			},
			*comment,
		})
	case 2:
		Commentid := c.Query("comment_id")
		err := service.DeleteComment(Commentid)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: 1,
				StatusMsg:  "删除评论失败",
			})
			return
		}
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 0,
			StatusMsg:  "删除评论成功",
		})
	}

}
