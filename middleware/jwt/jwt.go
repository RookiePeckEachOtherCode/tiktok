package jwt

import (
	"fmt"
	"net/http"
	"tiktok/configs"
	"tiktok/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

func NewToken(userId int64) (string, error) {
	// 设置过期时间为当前时间7天
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
	return token.SignedString([]byte(configs.JWT_KEY))
}

// 解析token,返回带有userId的Claims
func ParseToken(tokenString string) (*Claims, bool) {
	// 解析token
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.JWT_KEY), nil //要用将密钥变成字节数组形式才行
	})
	fmt.Printf("%v\n", token)
	// 判断token是否有效
	if token != nil {
		if claim, ok := token.Claims.(*Claims); ok && token.Valid {
			return claim, true
		}
	}
	return nil, false
}

// 鉴权
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")

		// 如果token为空
		if len(tokenStr) == 0 {
			c.JSON(http.StatusUnauthorized, model.Response{
				StatusCode: -1,
				StatusMsg:  "unauthorized: 用户不存在或者未登录",
			})
			c.Abort()
			return
		}
		fmt.Printf("%v\n", tokenStr)
		token, ok := ParseToken(tokenStr)
		// 如果token无效
		if !ok {
			c.JSON(http.StatusForbidden, model.Response{
				StatusCode: -1,
				StatusMsg:  "forbidden: token无效",
			})
			c.Abort()
			return
		}

		// 如果token过期
		if time.Now().Unix() > token.ExpiresAt {
			c.JSON(http.StatusOK, model.Response{
				StatusCode: -1,
				StatusMsg:  "token已过期",
			})

			c.Abort()
			return
		}
		// 将userId写入上下文
		c.Set("userId", token.UserId)
		// 继续执行
		c.Next()
	}
}
