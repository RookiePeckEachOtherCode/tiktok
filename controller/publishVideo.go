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
	_userId, exsist := c.Get("user_id")
	if !exsist {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户id 类型错误",
		})
		return
	}
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
		fmt.Printf("%v\n", err)
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

	coverDir := filepath.Dir(v.CoverSavePath)

	if err := os.MkdirAll(coverDir, os.ModePerm); err != nil {
		util.PrintLog(fmt.Sprintf("封面文件夹创建失败: %s", err.Error()))
		return err
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("middleware/ffmpeg/ffmpeg.exe", "-i", v.VideoSavePath, "-vframes", "1", "-q:v", "2", v.CoverSavePath)
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
