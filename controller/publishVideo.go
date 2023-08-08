package controller

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os/exec"
	"path/filepath"
	"tiktok/configs"
	"tiktok/model"
	"tiktok/service"
	"tiktok/util"

	"log"

	"github.com/gin-gonic/gin"
)

type Videoinfo struct {
	file          *multipart.FileHeader
	VideoName     string
	CoverName     string
	VideoSavePath string
	CoverSavePath string
}

func PublishVideo(c *gin.Context) {
	titile := c.PostForm("title")
	_userId, _ := c.Get("user_id")
	userId, ok := _userId.(int64)
	if !ok {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户id 类型错误",
		})
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "视频文件错误",
		})
		return
	}

	videoinfo := GetVideoInfo(data)

	util.PrintLog(fmt.Sprintln("videoinfo: ", videoinfo))

	if err := videoinfo.SaveVideo(c); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("视频保存失败: %s", err.Error()),
		})
		log.Printf("%v\n", err)
		return
	}
	if err := videoinfo.SaveCover(); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("封面保存失败: %s", err.Error()),
		})
		log.Printf("%v\n", err)
		return
	}

	//持久化
	service.PublishVideo(userId, videoinfo.VideoName, videoinfo.CoverName, titile)
}

func GetVideoInfo(file *multipart.FileHeader) *Videoinfo {
	videoinfo := Videoinfo{}
	videoinfo.VideoName = util.NewFileName(file.Filename)
	videoinfo.CoverName = string([]byte(videoinfo.VideoName)[:len(videoinfo.VideoName)-len(filepath.Ext(videoinfo.VideoName))]) + ".jpg"
	videoinfo.VideoSavePath = filepath.Join(configs.VIDEO_SAVE_PATH, videoinfo.VideoName)
	videoinfo.CoverSavePath = filepath.Join(configs.VIDEO_COVER_SAVE_PATH, videoinfo.CoverName)
	videoinfo.file = file
	return &videoinfo
}

func (v Videoinfo) SaveVideo(c *gin.Context) error {
	return c.SaveUploadedFile(v.file, v.VideoSavePath)
}

func (v Videoinfo) SaveCover() error {
	videoPath := v.VideoSavePath
	coverPath := v.CoverSavePath

	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vframes", "1", "-q:v", "2", coverPath)

	// 执行命令
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return nil
}
