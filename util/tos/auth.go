package tos

import (
	"context"
	"fmt"
	"math/rand"
	"mime/multipart"
	"tiktok/configs"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var (
	AccessKey = configs.AccessKey
	SecretKey = configs.SecretKey
)
var ( //video
	VideoBucketName = configs.VideoBucketName
	VideoBucketUrl  = configs.VideoBucketUrl
)
var ( //cover
	CoverBucketName = configs.CoverBucketName
	CoverBucketUrl  = configs.CoverBucketUrl
)
var ( //avatar
	AvatarBucketName = configs.AvatarBucketName
	AvatarBucketUrl  = configs.AvatarBucketUrl
)

func UploadToQiNiu(file *multipart.FileHeader) (*string, error) {

	data, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer data.Close()

	putPlicy := storage.PutPolicy{
		Scope: VideoBucketName,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)

	upToken := putPlicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan, // 华南区
		UseCdnDomains: false,
		UseHTTPS:      false, // 非https
	}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	key := "video/" + file.Filename
	err = formUploader.Put(context.Background(), &ret, upToken, key, data, file.Size, &putExtra)

	if err != nil {
		return nil, err
	}

	url := VideoBucketUrl + ret.Key
	return &url, nil
}
func GetRandomAvatar() string {
	//生成一个[1,8]的随机数
	randNum := rand.Intn(9)
	return fmt.Sprintf("%s%d.jpg", AvatarBucketUrl, randNum)
}
