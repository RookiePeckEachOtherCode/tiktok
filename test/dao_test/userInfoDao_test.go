package dao_test

import (
	"testing"
	"tiktok/dao"
	"tiktok/middleware/redis"

	"github.com/stretchr/testify/assert"
)

// GetUserInfoById 根据用户id获取用户信息
func TestGetUserInfoById(t *testing.T) {
	TestInit()

	users := []dao.UserInfo{
		{
			ID:            6,
			Name:          "TestGetUserInfoById",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}

	type args struct {
		testName string
		ID       int64
		wantUser dao.UserInfo
		wantErr  bool
	}
	tests := []args{
		{ // 正确的用户id
			testName: "TestGetUserInfoById",
			ID:       6,
			wantUser: users[0],
			wantErr:  false,
		},
		{ // 错误的用户id
			testName: "TestGetUserInfoById-nil",
			ID:       0,
			wantUser: dao.UserInfo{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		var Err bool
		t.Run((tt.testName), func(t *testing.T) {
			user, err := dao.GetUserInfoById(tt.ID)
			if err != nil {
				Err = true
			} else {
				if assert.EqualValues(t, tt.wantUser, *user) {
					Err = false
				} else {
					Err = true
				}
			}
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	for _, user := range users {
		dao.DB.Delete(&user)
	}
}

// AddUserInfo 保存用户信息到数据库
func TestAddUserInfo(t *testing.T) {
	TestInit()

	users := []dao.UserInfo{
		{
			ID:            7,
			Name:          "TestAddUserInfo",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
		},
	}

	type args struct {
		testName string
		user     *dao.UserInfo
		wantErr  bool
	}
	tests := []args{
		{ // 正确的用户信息
			testName: "TestAddUserInfo",
			user:     &users[0],
			wantErr:  false,
		},
		{ // 错误的用户信息
			testName: "TestAddUserInfo-nil",
			user:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		var Err bool
		t.Run((tt.testName), func(t *testing.T) {
			err := dao.AddUserInfo(tt.user)
			if err != nil {
				Err = true
			} else {
				Err = false
			}
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	if dao.CheckIsExistByID(7) {
		dao.DB.Delete(&dao.UserInfo{}, 7)
	}
}

// 通过名字查询用户是否存在
func TestCheckIsExistByName(t *testing.T) {
	TestInit()

	users := []dao.UserInfo{
		{ //存在
			ID:            8,
			Name:          "TestCheckIsExistByName",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}

	type args struct {
		testName string
		name     string
		wantErr  bool
	}
	tests := []args{
		{ // 正确的用户名
			testName: "TestCheckIsExistByName-ok",
			name:     "TestCheckIsExistByName",
			wantErr:  true,
		},
		{ //不存在的用户名
			testName: "TestCheckIsExistByName-not-exist",
			name:     "TestCheckIsExistByName-not-exist",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		var Err bool
		t.Run((tt.testName), func(t *testing.T) {
			Err = dao.CheckIsExistByName(tt.name)
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	for _, user := range users {
		dao.DB.Delete(&user)
	}
}

// 通过id查询用户是否存在
func TestCheckIsExistByID(t *testing.T) {
	TestInit()
	users := []dao.UserInfo{
		{ //存在
			ID:            6,
			Name:          "TestCheckIsExistByID",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}

	type args struct {
		testName string
		ID       int64
		wantErr  bool
	}
	tests := []args{
		{ // 正确的用户id
			testName: "TestCheckIsExistByID-ok",
			ID:       6,
			wantErr:  true,
		},
		{ // 错误的用户id
			testName: "TestCheckIsExistByID-err",
			ID:       0,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		var Err bool
		t.Run((tt.testName), func(t *testing.T) {
			Err = dao.CheckIsExistByID(tt.ID)
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	for _, user := range users {
		dao.DB.Delete(&user)
	}
}

func TestGetIsFavorite(t *testing.T) {
	TestInit()

	users := []dao.UserInfo{
		{ //喜欢
			ID:            6,
			Name:          "TestGetIsFavorite",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
			FavorVideos: []*dao.Video{
				{
					ID:         6,
					UserInfoID: 6,
					Title:      "TestGetIsFavorite",
					PlayURL:    "TestGetIsFavorite",
					CoverURL:   "TestGetIsFavorite",
					Author: dao.UserInfo{
						ID: 6,
					},
				},
			},
		},
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}

	type args struct {
		testName string
		userInfo *dao.UserInfo
		videoId  int64
		wantErr  bool
	}
	tests := []args{
		{ //喜欢
			testName: "TestGetIsFavorite-true",
			userInfo: &users[0],
			videoId:  6,
			wantErr:  true,
		},
		{ //不喜欢
			testName: "TestGetIsFavorite-false",
			userInfo: &dao.UserInfo{
				ID: 6,
			},
			videoId: 5,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		var Err bool
		t.Run((tt.testName), func(t *testing.T) {
			Err = tt.userInfo.GetIsFavorite(tt.videoId)
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	for _, user := range users {
		dao.DB.Exec("delete from user_favor_videos where user_info_id = ?", user.ID)
		for _, video := range user.FavorVideos {
			dao.DB.Delete(&video)
		}
		dao.DB.Delete(&user)
	}

}

// FavoriteVideo 给视频点赞
func TestToFavoriteVideo(t *testing.T) {
	TestInit()

	users := []dao.UserInfo{
		{ //喜欢
			ID:            6,
			Name:          "TestToFavoriteVideo",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
			FavorVideos: []*dao.Video{
				{
					ID:            7,
					UserInfoID:    6,
					Title:         "TestToFavoriteVideo",
					PlayURL:       "TestToFavoriteVideo",
					CoverURL:      "TestToFavoriteVideo",
					FavoriteCount: 0,
					Author: dao.UserInfo{
						ID: 6,
					},
				},
			},
		},
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}

	type args struct {
		testName string
		userInfo *dao.UserInfo
		wantErr  bool
	}
	tests := []args{
		{ //成功
			testName: "TestToFavoriteVideo-ok",
			userInfo: &users[0],
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		var Err bool
		t.Run((tt.testName), func(t *testing.T) {
			err := tt.userInfo.ToFavoriteVideo(tt.userInfo.FavorVideos[0])
			if err != nil {
				Err = true
			} else {
				var favoriteCount int64
				dao.DB.Model(&dao.Video{}).Where("id=?", tt.userInfo.FavorVideos[0].ID).Select("favorite_count").Find(&favoriteCount)
				if favoriteCount == 1 {
					Err = false
				} else {
					Err = true
				}
			}
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	for _, user := range users {
		dao.DB.Exec("delete from user_favor_videos where user_info_id = ?", user.ID)
		for _, video := range user.FavorVideos {
			dao.DB.Delete(&video)
		}
		dao.DB.Delete(&user)
	}

}

// 通过id获取用户喜欢的视频列表
func TestGetFavoriteList(t *testing.T) {
	TestInit()

	users := []dao.UserInfo{
		{ //成功
			ID:            6,
			Name:          "TestGetFavoriteList",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
			FavorVideos: []*dao.Video{
				{
					ID:         6,
					UserInfoID: 6,
					Title:      "TestGetFavoriteList",
					PlayURL:    "TestGetFavoriteList",
					CoverURL:   "TestGetFavoriteList",
					Author: dao.UserInfo{
						ID: 6,
					},
				},
			},
		},
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}

	type args struct {
		testName string
		userInfo *dao.UserInfo
		wantErr  bool
	}

	tests := []args{
		{
			testName: "TestGetFavoriteList",
			userInfo: &users[0],
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		var Err bool
		t.Run((tt.testName), func(t *testing.T) {
			videos, err := dao.GetFavoriteList(tt.userInfo.ID)
			if err != nil {
				Err = true
			} else {
				if videos[0].ID == tt.userInfo.FavorVideos[0].ID {
					Err = false
				} else {
					Err = true
				}
			}
			assert.Equal(t, tt.wantErr, Err)
		})
	}
	for _, user := range users {
		dao.DB.Exec("delete from user_favor_videos where user_info_id = ?", user.ID)
		for _, video := range user.FavorVideos {
			dao.DB.Delete(&video)
		}
		dao.DB.Delete(&user)
	}
}

// 用户获赞数+1
func TestPlusFavCount(t *testing.T) {
	TestInit()

	users := []dao.UserInfo{
		{ //成功
			ID:            6,
			Name:          "TestPlusFavCount",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
			TotalFavorite: 0,
		},
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}

	type args struct {
		testName string
		userInfo *dao.UserInfo
		wantErr  bool
	}

	tests := []args{
		{
			testName: "TestPlusFavCount",
			userInfo: &users[0],
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		var Err bool
		t.Run((tt.testName), func(t *testing.T) {
			tt.userInfo.PlusFavCount()
			count := redis.New(redis.LIKED).GetUserReceivedLikeCount(tt.userInfo.ID)
			if count == tt.userInfo.TotalFavorite+1 {
				Err = false
			} else {
				Err = true
			}
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	for _, user := range users {
		dao.DB.Delete(&user)
	}
}

func TestFollowAct(t *testing.T) {
	TestInit()
	users := []dao.UserInfo{
		{ //关注者
			ID:            6,
			Name:          "TestFollowAct",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
		{ //被关注者
			ID:            7,
			Name:          "TestFollowAct-2",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}

	type args struct {
		testName string
		user     *dao.UserInfo
		toUser   *dao.UserInfo
		wantErr  bool
	}

	tests := []args{
		{
			testName: "TestFollowAct",
			user:     &users[0],
			toUser:   &users[1],
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run((tt.testName), func(t *testing.T) {
			var Err bool
			if err := users[0].FollowAct(tt.toUser); err != nil {
				Err = true
			} else {
				dao.DB.Model(&dao.UserInfo{}).Where("id=?", tt.user.ID).Select("follow_count").Find(&tt.user.FollowCount)
				dao.DB.Model(&dao.UserInfo{}).Where("id=?", tt.toUser.ID).Select("follower_count").Find(&tt.toUser.FollowerCount)

				if tt.user.FollowCount == 1 && tt.toUser.FollowerCount == 1 {
					Err = false
				} else {
					Err = true
				}
			}
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	for _, user := range users {
		dao.DB.Exec("delete from user_relations where user_info_id = ?", user.ID)
		dao.DB.Exec("delete from user_relations where follow_user_id = ?", user.ID)
		dao.DB.Delete(&user)
	}
}

// 取关
func TestUnFollowAct(t *testing.T) {
	TestInit()
	users := []dao.UserInfo{
		{ //关注者
			ID:            6,
			Name:          "TestUnFollowAct",
			FollowCount:   1,
			FollowerCount: 0,
			IsFollow:      true,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
		{ //被关注者
			ID:            7,
			Name:          "TestUnFollowAct-2",
			FollowCount:   0,
			FollowerCount: 1,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}

	dao.DB.Exec("insert into user_relations (user_info_id, follow_user_id) values (?, ?)", users[0].ID, users[1].ID)

	type args struct {
		testName string
		user     *dao.UserInfo
		toUser   *dao.UserInfo
		wantErr  bool
	}

	tests := []args{
		{
			testName: "TestUnFollowAct",
			user:     &users[0],
			toUser:   &users[1],
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run((tt.testName), func(t *testing.T) {
			var Err bool
			if err := users[0].UnFollowAct(tt.toUser); err != nil {
				Err = true
			} else {
				dao.DB.Model(&dao.UserInfo{}).Where("id=?", tt.user.ID).Select("follow_count").Find(&tt.user.FollowCount)
				dao.DB.Model(&dao.UserInfo{}).Where("id=?", tt.toUser.ID).Select("follower_count").Find(&tt.toUser.FollowerCount)

				if tt.user.FollowCount == 0 && tt.toUser.FollowerCount == 0 {
					Err = false
				} else {
					Err = true
				}
			}
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	for _, user := range users {
		dao.DB.Exec("delete from user_relations where user_info_id = ?", user.ID)
		dao.DB.Exec("delete from user_relations where follow_user_id = ?", user.ID)
		dao.DB.Delete(&user)
	}

}

// 获取用户关注列表
func TestGetFloList(t *testing.T) {
	TestInit()

	users := []dao.UserInfo{
		{ //关注者
			ID:            6,
			Name:          "TestGetFloList",
			FollowCount:   1,
			FollowerCount: 0,
			IsFollow:      true,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
		{ //被关注者
			ID:            7,
			Name:          "TestGetFloList-2",
			FollowCount:   0,
			FollowerCount: 1,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}

	dao.DB.Exec("insert into user_relations (user_info_id, follow_id) values (?, ?)", users[0].ID, users[1].ID)

	type args struct {
		testName string
		user     *dao.UserInfo
		wantErr  bool
	}

	tests := []args{
		{
			testName: "TestGetFloList",
			user:     &users[0],
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run((tt.testName), func(t *testing.T) {
			var Err bool
			floList, err := dao.GetFloList(tt.user.ID)
			if err != nil {
				Err = true
			} else {
				if floList[0].ID == users[1].ID {
					Err = false
				} else {
					Err = true
				}
			}
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	for _, user := range users {
		dao.DB.Exec("delete from user_relations where user_info_id = ?", user.ID)
		dao.DB.Exec("delete from user_relations where follow_id = ?", user.ID)
		dao.DB.Delete(&user)
	}

}

func TestGetFollowerList(t *testing.T) {
	TestInit()

	users := []dao.UserInfo{
		{ //关注者
			ID:            6,
			Name:          "TestGetFloList",
			FollowCount:   1,
			FollowerCount: 0,
			IsFollow:      true,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
		{ //被关注者
			ID:            7,
			Name:          "TestGetFloList-2",
			FollowCount:   0,
			FollowerCount: 1,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
		{ //没有关注关系
			ID:            8,
			Name:          "TestGetFloList-3",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}

	dao.DB.Exec("insert into user_relations (user_info_id, follow_id) values (?, ?)", users[0].ID, users[1].ID)
	type args struct {
		testName string
		user     *dao.UserInfo
		toUser   *dao.UserInfo
		wantErr  bool
	}

	tests := []args{
		{
			testName: "TestGetFollowerList",
			user:     &users[1],
			toUser:   &users[0],
			wantErr:  false,
		},
		{
			testName: "TestGetFollowerList-nil",
			user:     &users[2],
			toUser:   &users[2],
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run((tt.testName), func(t *testing.T) {
			var Err bool
			followerList, err := dao.GetFollowerList(tt.toUser.ID)

			if err != nil {
				Err = true
			} else {
				if followerList == nil {
					Err = false
				} else if followerList[0].ID == tt.user.ID {
					Err = false
				} else {
					Err = true
				}
			}
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	for _, user := range users {
		dao.DB.Exec("delete from user_relations where user_info_id = ?", user.ID)
		dao.DB.Exec("delete from user_relations where follow_id = ?", user.ID)
		dao.DB.Delete(&user)
	}
}

// GetUserRelation 判断两个用户之间是否存在关注关系
func TestGetUserRelation(t *testing.T) { //uid是否关注tid
	TestInit()

	users := []dao.UserInfo{
		{ //关注者
			ID:            6,
			Name:          "TestGetFloList",
			FollowCount:   1,
			FollowerCount: 0,
			IsFollow:      true,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
		{ //被关注者
			ID:            7,
			Name:          "TestGetFloList-2",
			FollowCount:   0,
			FollowerCount: 1,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
		{ //没有关注关系
			ID:            8,
			Name:          "TestGetFloList-3",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}

	dao.DB.Exec("insert into user_relations (user_info_id, follow_id) values (?, ?)", users[0].ID, users[1].ID)

	type args struct {
		testName string
		user     dao.UserInfo
		toUser   dao.UserInfo
		wantErr  bool
	}

	tests := []args{
		{ //存在关注关系
			testName: "TestGetUserRelation-ok",
			user:     users[0],
			toUser:   users[1],
			wantErr:  true,
		},
		{
			testName: "TestGetUserRelation-nil",
			user:     users[1],
			toUser:   users[2],
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run((tt.testName), func(t *testing.T) {
			var Err bool
			if dao.GetUserRelation(tt.user.ID, tt.toUser.ID) {
				Err = true
			} else {
				Err = false
			}
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	for _, user := range users {
		dao.DB.Exec("delete from user_relations where user_info_id = ?", user.ID)
		dao.DB.Exec("delete from user_relations where follow_id = ?", user.ID)
		dao.DB.Delete(&user)
	}

}
func TestGetMutualFriendListById(t *testing.T) {
	TestInit()

	users := []dao.UserInfo{
		{ //互相关注者
			ID:            6,
			Name:          "TestGetFloList",
			FollowCount:   1,
			FollowerCount: 1,
			IsFollow:      true,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
		{ //互相关注者
			ID:            7,
			Name:          "TestGetFloList-2",
			FollowCount:   1,
			FollowerCount: 1,
			IsFollow:      true,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
		{ //非互相关注者
			ID:            8,
			Name:          "TestGetFloList-3",
			FollowCount:   1,
			FollowerCount: 1,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			Signature:     "这个人很懒，什么都没写",
		},
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}
	dao.DB.Exec("insert into user_relations (user_info_id, follow_id) values (?, ?)", users[0].ID, users[1].ID)
	dao.DB.Exec("insert into user_relations (user_info_id, follow_id) values (?, ?)", users[1].ID, users[0].ID)

	type args struct {
		testName string
		user     dao.UserInfo
		toUser   dao.UserInfo
		wantErr  bool
	}

	tests := []args{
		{ //存在互相关注关系
			testName: "TestGetMutualFriendListById-ok",
			user:     users[0],
			toUser:   users[1],
			wantErr:  true,
		},
		{
			testName: "TestGetMutualFriendListById-nil",
			user:     users[1],
			toUser:   users[2],
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run((tt.testName), func(t *testing.T) {
			var Err bool
			mutualFriendList, err := dao.GetMutualFriendListById(tt.toUser.ID)
			if err != nil && mutualFriendList == nil {
				Err = false
			} else {
				if mutualFriendList[0].ID == tt.user.ID {
					Err = true
				} else {
					Err = false
				}
			}
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	for _, user := range users {
		dao.DB.Exec("delete from user_relations where user_info_id = ?", user.ID)
		dao.DB.Exec("delete from user_relations where follow_id = ?", user.ID)
		dao.DB.Delete(&user)
	}
}
