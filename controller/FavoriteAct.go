package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/model"
	"tiktok/service"
)

func RecFavorite(c *gin.Context) {
	Uid, exists := c.Get("user_id")
	if exists == false {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取用户id失败",
		})
	}
	id, erro := Uid.(int64)
	if !erro {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: 1,
			StatusMsg:  "断言id失败",
		})
	}
	vid := c.Query("video_id")                //获取视频id
	Vid, err := strconv.ParseInt(vid, 10, 64) //格式转换以查询
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: 1,
			StatusMsg:  "获取视频id失败",
		})
	}
	act := c.Query("action_type") //获取点赞类型
	Act, err := strconv.ParseInt(act, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: 1,
			StatusMsg:  "点赞失败",
		})
	}
	err = service.HandleFav(Vid, Act, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: 1,
			StatusMsg:  "点赞保存失败",
		})
		fmt.Printf("%v", err)
	} else {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 0,
			StatusMsg:  "点赞操作成功",
		})
	}
}
