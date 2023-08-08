package service

import (
	"tiktok/dao"
	"tiktok/model"
)

type FavListRes struct {
	model.Response
	VideoList []Video `json:"video_list"`
}
type Video struct {
	Author        Author `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}
type Author struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	ID              int64  `json:"id"`               // 用户id
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name            string `json:"name"`             // 用户名称
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
}

func HandleFavListQuery(uid int64) (res *FavListRes, err error) {
	vlist, err := dao.GetFavList(uid)
	if err != nil {
		return nil, err
	}
	res = &FavListRes{}
	// 创建 Video 切片并遍历 vlist 添加到 VideoList
	res.VideoList = make([]Video, len(vlist))
	for i, v := range vlist {
		// 转换底层数据结构为 Video 对象
		video := Video{
			Author: Author{
				Avatar:          "", //头象不知道在哪
				BackgroundImage: "", //背景也不知道在哪
				FavoriteCount:   int64(len(v.Author.FavorVideos)),
				FollowCount:     v.Author.FollowCount,
				FollowerCount:   v.Author.FollowerCount,
				ID:              v.Author.ID,
				IsFollow:        false,
				Name:            v.Author.Name,
				Signature:       "", //个人简介
				TotalFavorited:  "", //总获赞数目
				WorkCount:       int64(len(v.Author.Videos)),
			},
			CommentCount:  v.CommentCount,
			CoverURL:      v.CoverURL,
			FavoriteCount: v.FavoriteCount,
			ID:            v.ID,
			IsFavorite:    v.IsFavorite,
			PlayURL:       v.PlayURL,
			Title:         v.Title,
		}
		res.VideoList[i] = video
	}
	res.StatusCode = 0
	res.StatusMsg = "获取列表成功"
	return res, nil
}
