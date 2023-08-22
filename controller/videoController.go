package controller

import (
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"tiktok/configs"
	"tiktok/dao"
	"tiktok/middleware/jwt"
	"tiktok/service"
	"tiktok/util"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	dao.Response
	VideoList []*dao.Video `json:"video_list"`
	NextTime  int64        `json:"next_time"`
}
type PublishListResponse struct {
	dao.Response
	VideoList []dao.Video `json:"video_list"`
}

type VideoInfos struct {
	file          *multipart.FileHeader
	VideoName     string
	CoverName     string
	VideoSavePath string
	CoverSavePath string
}

// 获取feed流
func VideoFeedController(c *gin.Context) {
	inputTime := c.Query("latest_time")
	var latestTime time.Time
	if len(inputTime) == 0 {
		_latestTime, _ := strconv.ParseInt(inputTime, 10, 64)
		_latestTime /= 1000
		latestTime = time.Unix(_latestTime, 0)
	} else {
		latestTime = time.Now()
	}
	token := c.Query("token")
	if len(token) == 0 {
		notLogin(latestTime, c)
		return
	}
	hasLogin(token, latestTime, c)
}

// 获取发布列表
func PublishListController(c *gin.Context) {
	_userId, exi := c.Get("user_id")
	if !exi {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 0,
			StatusMsg:  "获取用户id失败",
		})
		return
	}

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

	videoInfo := getVideoInfo(data)

	if err := videoInfo.saveVideo(c); err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "视频保存失败: " + err.Error(),
		})
		log.Println("视频保存失败: ", err)
		return
	}
	if err := videoInfo.saveCover(); err != nil {
		log.Println("封面保存失败: ", err)
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "封面保存失败: " + err.Error(),
		})
		return
	}
	//持久化
	if err := service.PublishVideoService(userId, videoInfo.VideoName, videoInfo.CoverName, title); err != nil {
		log.Println("视频持久化失败: ", err)
	}
	c.JSON(http.StatusOK, dao.Response{
		StatusCode: 0,
		StatusMsg:  "视频发布成功",
	})
}

// 已经登陆的情况
func hasLogin(token string, latestTime time.Time, c *gin.Context) {
	_userId, _ := jwt.ParseToken(token)
	userId := _userId.UserId

	videos, err := service.VideoFeedService(userId, time.Now())
	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取feed流失败: " + err.Error(),
		})

		return
	}

	nextTime, err := dao.GetNextTimeByVideoId((*videos)[0].ID)
	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取nextTime 失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, FeedResponse{
		dao.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		*videos,
		nextTime.UnixNano() / 1e6,
	})

}

// 未登录的情况
func notLogin(latestTime time.Time, c *gin.Context) {
	videos, err := service.VideoFeedService(0, time.Now())
	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取feed流失败: " + err.Error(),
		})
		return
	}
	nextTime, err := dao.GetNextTimeByVideoId((*videos)[0].ID)
	if err != nil {
		c.JSON(http.StatusOK, dao.Response{
			StatusCode: 1,
			StatusMsg:  "获取nextTime失败: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		dao.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		*videos,
		nextTime.UnixNano() / 1e6,
	})
}

func getVideoInfo(file *multipart.FileHeader) *VideoInfos {
	videoInfo := VideoInfos{}
	videoInfo.VideoName = util.NewFileName(file.Filename)
	videoInfo.CoverName = string([]byte(videoInfo.VideoName)[:len(videoInfo.VideoName)-len(filepath.Ext(videoInfo.VideoName))]) + ".jpg"
	videoInfo.VideoSavePath = filepath.Join(configs.VIDEO_SAVE_PATH, videoInfo.VideoName)
	videoInfo.CoverSavePath = filepath.Join(configs.VIDEO_COVER_SAVE_PATH, videoInfo.CoverName)
	videoInfo.file = file
	return &videoInfo
}

func (v VideoInfos) saveVideo(c *gin.Context) error {
	return c.SaveUploadedFile(v.file, v.VideoSavePath)
}

func (v VideoInfos) saveCover() error {

	coverDir := filepath.Dir(v.CoverSavePath)

	if err := os.MkdirAll(coverDir, os.ModePerm); err != nil {
		return err
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("middleware/ffmpeg/ffmpeg.exe", "-i", v.VideoSavePath, "-vframes", "1", "-q:v", "2", v.CoverSavePath)
	case "linux":
		cmd = exec.Command("ffmpeg", "-i", v.VideoSavePath, "-vframes", "1", "-q:v", "2", v.CoverSavePath)
	default:
		cmd = exec.Command("middleware/ffmpeg/ffmpeg", "-i", v.VideoSavePath, "-vframes", "1", "-q:v", "2", v.CoverSavePath)
	}

	// 改用 exec.Command 的正确用法
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
