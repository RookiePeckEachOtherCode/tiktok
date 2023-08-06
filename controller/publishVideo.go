package controller

import (
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"tiktok/configs"
	"tiktok/model"
	"tiktok/service"
	"tiktok/util"

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

	if err := videoinfo.SaveVideo(c); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "视频文件保存失败",
		})
		return
	}
	if err := videoinfo.SaveCover(c); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "视频封面保存失败",
		})
		return
	}

	//持久化
	service.PublishVideo(userId, videoinfo.VideoName, videoinfo.CoverName, titile)
}

func GetVideoInfo(file *multipart.FileHeader) *Videoinfo {
	videoinfo := Videoinfo{}
	videoinfo.VideoName = util.NewFileName(file.Filename)
	videoinfo.CoverName = util.NewFileName(file.Filename)
	videoinfo.VideoSavePath = filepath.Join(configs.VIDEO_SAVE_PATH, videoinfo.VideoName)
	videoinfo.CoverSavePath = filepath.Join(configs.VIDEO_COVER_SAVE_PATH, videoinfo.CoverName)
	videoinfo.file = file
	return &videoinfo
}

// SaveVideo
func (v *Videoinfo) SaveVideo(c *gin.Context) error {

	// 保存到临时目录
	tmpPath := filepath.Join(os.TempDir(), v.VideoName)

	if err := c.SaveUploadedFile(v.file, tmpPath); err != nil {
		return err
	}

	// 移动到正式目录
	err := os.Rename(tmpPath, v.VideoSavePath)
	if err != nil {
		return err
	}

	return nil
}

// SaveCover
func (v *Videoinfo) SaveCover(c *gin.Context) error {

	// 创建临时目录
	tmpCoverPath := filepath.Join(os.TempDir(), v.CoverName)

	cmd := exec.Command("ffmpeg", "-i", v.VideoSavePath,
		"-ss", "00:00:01", "-vframes", "1", tmpCoverPath)

	err := cmd.Run()
	if err != nil {
		os.Remove(tmpCoverPath) // 移除临时文件
		return err
	}

	// 移动到正式目录
	err = os.Rename(tmpCoverPath, v.CoverSavePath)
	if err != nil {
		return err
	}

	return nil
}
