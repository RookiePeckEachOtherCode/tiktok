package controller

import (
	"fmt"
	"log"
	"net/http"
	"tiktok/model"
	"tiktok/service"

	"github.com/gin-gonic/gin"
)

func RecFavList(c *gin.Context) {
	_uid, _ := c.Get("user_id")
	uid, ok := _uid.(int64)

	if !ok {
		log.Println("用户id解析失败")
		c.JSON(http.StatusBadRequest, service.FavListRes{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "用户id解析失败",
			},
		})
		return
	}

	res, err := service.HandleFavListQuery(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, service.FavListRes{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "获取喜欢列表失败",
			},
		})
		fmt.Printf("%v\n", err)
	} else {
		c.JSON(http.StatusOK, res)
		println("已经上传了喜欢列表")
	}

}
