package controller

import (
	"errors"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"tiktok/configs"
	"tiktok/dao"
	"tiktok/middleware/jwt"
	"tiktok/service"
	"tiktok/util"
	"tiktok/util/tos"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	dao.Response
	*service.FeedVideoFlow
}
type PublishListResponse struct {
	dao.Response
	VideoList []dao.Video `json:"video_list"`
}

type VideoInfo struct {
	file          *multipart.FileHeader
	VideoName     string
	CoverName     string
	VideoSavePath string
	CoverSavePath string
}

// 获取feed流
func VideoFeedController(c *gin.Context) {
	token, ok := c.GetQuery("token")

	if !ok {
		if err := notLogin(c); err != nil {
			log.Println("feed:", err)
			c.JSON(http.StatusOK, dao.Response{
				StatusCode: 1,
				StatusMsg:  "获取feed失败: " + err.Error(),
			})
		}
		return
	}
	if err := hasLogin(c, token); err != nil {
		log.Println("feed:", err)
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取feed失败: " + err.Error(),
		})
	}
}

// 获取发布列表
func PublishListController(c *gin.Context) {
	_userId, _ := c.Get("user_id")
	userId, ok := _userId.(int64)
	//判断用户id类型是否正确
	if !ok {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "用户id类型错误",
		})

	}
	//获取发布列表
	videoList, err := service.PublishListService(userId)

	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取发布列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, PublishListResponse{
		dao.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		*videoList,
	})

}

// 发布视频
func PublishVideoController(c *gin.Context) {
	_title := c.PostForm("title")
	title, _ := util.FilterDirty(_title)

	_userId, _ := c.Get("user_id")
	userId, ok := _userId.(int64)

	if !ok {
		log.Println("用户id解析失败")
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "用户id解析失败",
		})
		return
	}

	data, err := c.FormFile("data")

	if err != nil {
		log.Println("视频获取失败: ", err)
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "视频获取失败: " + err.Error(),
		})
		return
	}

	videoInfo, err := getVideoInfo(data)

	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "视频获取失败: " + err.Error(),
		})
	}
	//持久化
	if err := service.PublishVideoService(userId, videoInfo.VideoSavePath, videoInfo.CoverSavePath, title); err != nil {
		log.Println("视频持久化失败: ", err)
	}
	c.JSON(http.StatusOK, dao.Response{
		StatusCode: 0,
		StatusMsg:  "视频发布成功",
	})
}

// 已经登陆的情况
func hasLogin(c *gin.Context, token string) error {
	if claims, ok := jwt.ParseToken(token); ok {
		if claims.ExpiresAt < time.Now().Unix() {
			return errors.New("登陆过期")
		}
		_latestTime := c.Query("latest_time")
		var latestTime time.Time
		intTime, err := strconv.ParseInt(_latestTime, 10, 64)
		if err != nil {
			latestTime = time.Unix(0, intTime*1e6)
		}
		videoList, err := service.VideoFeedService(claims.UserId, latestTime)
		if err != nil {
			return err
		}
		c.JSON(http.StatusOK, FeedResponse{
			Response: dao.Response{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			FeedVideoFlow: videoList,
		})
		return nil
	}
	return errors.New("token解析失败")
}

// 未登录的情况
func notLogin(c *gin.Context) error {
	_latestTime := c.Query("latest_time")
	var latestTime time.Time
	intTime, err := strconv.ParseInt(_latestTime, 10, 64)
	if err == nil {
		latestTime = time.Unix(0, intTime*1e6)
	}
	videoList, err := service.VideoFeedService(0, latestTime)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response: dao.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		FeedVideoFlow: videoList,
	})
	return nil
}

func getVideoInfo(file *multipart.FileHeader) (*VideoInfo, error) {
	videoInfo := VideoInfo{file: file}
	videoInfo.file.Filename = util.NewFileName(videoInfo.file.Filename)

	if err := videoInfo.saveVideo(); err != nil {
		return nil, err
	}
	coverName := strings.Split(videoInfo.file.Filename, ".")[0]
	videoInfo.CoverName = "cover" + coverName + ".jpg"
	videoInfo.CoverSavePath = configs.CoverBucketUrl + videoInfo.CoverName
	return &videoInfo, nil
}

func (v *VideoInfo) saveVideo() error {
	if v.file == nil {
		return errors.New("视频文件为空")
	}
	path, err := tos.UploadToQiNiu(v.file)
	if err != nil {
		return err
	}
	v.VideoSavePath = *path
	return nil
}
