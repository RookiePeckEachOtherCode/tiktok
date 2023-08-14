package controller

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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
		log.Println("用户id解析失败")
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户id解析失败",
		})
		return
	}

	data, err := c.FormFile("data")

	if err != nil {
		log.Println("视频获取失败: ", err)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "视频获取失败",
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
		log.Println("视频保存失败: ", err)
		return
	}
	if err := videoinfo.SaveCover(); err != nil {
		log.Println("封面保存失败: ", err)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  fmt.Sprintf("封面保存失败: %s", err.Error()),
		})
		return
	}
	//持久化
	if err := service.PublishVideo(userId, videoinfo.VideoName, videoinfo.CoverName, titile); err != nil {
		log.Println("视频持久化失败: ", err)
	}
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  "视频发布成功",
	})
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

	coverDir := filepath.Dir(v.CoverSavePath)

	if err := os.MkdirAll(coverDir, os.ModePerm); err != nil {
		util.PrintLog(fmt.Sprintf("封面文件夹创建失败: %s", err.Error()))
		return err
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("F:\\ffmpeg-2023-07-19-git-efa6cec759-essentials_build\\bin\\ffmpeg.exe", "-i", v.VideoSavePath, "-vframes", "1", "-q:v", "2", v.CoverSavePath)
	case "linux":
		cmd = exec.Command("middleware/ffmpeg/ffmpeg", "-i", v.VideoSavePath, "-vframes", "1", "-q:v", "2", v.CoverSavePath)
	default:
		cmd = exec.Command("middleware/ffmpeg/ffmpeg.exe", "-i", v.VideoSavePath, "-vframes", "1", "-q:v", "2", v.CoverSavePath)
	}

	// 改用 exec.Command 的正确用法
	err := cmd.Run()

	if err != nil {
		util.PrintLog(fmt.Sprintf("封面保存失败: %s", err.Error()))
		return err
	}

	return nil
}
