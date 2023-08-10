package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/dao"
	"tiktok/model"
	"tiktok/service"
	"time"
)

type get struct {
	Response model.Response
	comment  dao.Comment
}

func CommentAct(c *gin.Context) {
	var comment *dao.Comment
	_userid, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取用户id失败",
		})
	}
	userid, ok1 := _userid.(int64)
	user, errr := dao.GetUserInfoById(userid)
	if errr != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户信息获取失败",
		})
	}
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
			StatusCode: 0,
			StatusMsg:  "获取act失败",
		})
	}
	if Actiontype == 1 {
		Commenttext := c.Query("comment_text")
		comment = &dao.Comment{UserInfoID: userid, VideoID: Videoid, User: *user, Content: Commenttext, CreatedAt: time.Now(), CreateDate: time.Now().Format("2006-01-02")}
		err2 := service.OperateComment(Videoid, userid, Actiontype, Commenttext, comment)
		if err2 != nil {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: 1,
				StatusMsg:  "评论发布失败",
			})
			fmt.Printf("%v", err2)
			return
		}
		c.JSON(http.StatusOK, get{
			Response: model.Response{
				StatusCode: 0,
				StatusMsg:  "发布评论成功",
			},
			comment: *comment,
		})
	} else {
		Commentid := c.Query("comment_id")
		err3 := service.OperateComment(Videoid, userid, Actiontype, Commentid, comment)
		if err3 != nil {
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
