package dao_test

import (
	"testing"
	"tiktok/dao"

	"github.com/stretchr/testify/assert"
)

func TestJudgeUserPassword(t *testing.T) {
	TestInit()

	users := []dao.UserInfo{
		{
			ID:            5,
			Name:          "TestJudgeUserPassword",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			UserLoginInfo: &dao.UserLogin{
				ID:         5,
				UserInfoID: 5,
				Username:   "TestJudgeUserPassword",
				Password:   "123456",
			},
		},
	}

	type args struct {
		testName string
		ID       int64
		userName string
		password string
		wantErr  bool
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}

	tests := []args{
		{ // 正确的用户名和密码
			testName: "TestJudgeUserPassword-ok",
			ID:       5,
			userName: "TestJudgeUserPassword",
			password: "123456",
			wantErr:  false,
		},
		{ // 正确的用户名密码错误
			testName: "TestJudgeUserPassword-error",
			ID:       5,
			userName: "TestJudgeUserPassword",
			password: "114514",
			wantErr:  true,
		},
		{ // 用户不存在
			testName: "TestJudgeUserPassword-not-exist",
			ID:       0,
			userName: "TestJudgeUserPassword-not-exist",
			password: "114514",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.userName, func(t *testing.T) {
			var Err bool
			id, err := dao.JudgeUserPassword(tt.userName, tt.password)
			if err != nil && id == 0 {
				switch err.Error() {
				case "该用户不存在":
					Err = true
				case "密码错误":
					Err = true
				}
			} else {
				Err = false
			}
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	for _, user := range users {
		dao.DB.Delete(&user)
	}

}

func TestIsExistUserLoginInfoByName(t *testing.T) {
	TestInit()

	users := []dao.UserInfo{
		{ // 存在的用户名
			ID:            5,
			Name:          "TestIsExistUserLoginInfoByName",
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
			Avatar:        "defualt.jpg",
			UserLoginInfo: &dao.UserLogin{
				ID:         5,
				UserInfoID: 5,
				Username:   "TestIsExistUserLoginInfoByName",
			},
		},
	}

	type args struct {
		testName string
		userName string
		wantErr  bool
	}

	for _, user := range users {
		dao.DB.Create(&user)
	}

	tests := []args{
		{ // 存在的用户名
			testName: "TestIsExistUserLoginInfoByName-ok",
			userName: "TestIsExistUserLoginInfoByName",
			wantErr:  true,
		},
		{ // 不存在的用户名
			testName: "TestIsExistUserLoginInfoByName-not-exist",
			userName: "TestIsExistUserLoginInfoByName-not-exist",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			Err := dao.IsExistUserLoginInfoByName(tt.userName)
			assert.Equal(t, tt.wantErr, Err)
		})
	}

	for _, user := range users {
		dao.DB.Delete(&user)
	}
}
