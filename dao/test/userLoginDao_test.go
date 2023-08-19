package dao_test

import (
	"testing"
	"tiktok/dao"

	"github.com/stretchr/testify/assert"
)

func TestJudgeUserPassword(t *testing.T) {
	TestInit()

	type args struct {
		ID       int64
		userName string
		password string
		wantErr  bool
	}
	tests := []args{
		{ // 正确的用户名和密码
			ID:       5,
			userName: "TestJudgeUserPassword",
			password: "123456",
			wantErr:  false,
		},
		{ // 正确的用户名密码错误
			ID:       5,
			userName: "TestJudgeUserPassword",
			password: "114514",
			wantErr:  true,
		},
		{ // 用户不存在
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

}

func TestIsExistUserLoginInfoByName(t *testing.T) {
	TestInit()

	type args struct {
		userName string
		wantErr  bool
	}

	tests := []args{
		{ // 存在的用户名
			userName: "TestJudgeUserPassword",
			wantErr:  true,
		},
		{ // 不存在的用户名
			userName: "TestIsExistUserLoginInfoByName-not-exist",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.userName, func(t *testing.T) {
			Err := dao.IsExistUserLoginInfoByName(tt.userName)
			assert.Equal(t, tt.wantErr, Err)
		})
	}

}
