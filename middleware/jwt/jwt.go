package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

func NewToken(userId int64) (string, error) {
	// 设置过期时间为当前时间+7天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expirationTime.Unix(),
			// 发布时间
			IssuedAt: time.Now().Unix(),
			// 发布者
			Issuer: "tiktok",
			// 主题
			Subject: "tiktok-token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}
