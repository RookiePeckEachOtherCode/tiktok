package jwt

import (
	"fmt"
	"net/http"
	"tiktok/configs"
	"tiktok/dao"
	"tiktok/util"
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
	//fmt.Printf("%v\n", token)
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
		// 如果token为空
		tokenStr := c.Query("token")
		if len(tokenStr) == 0 {
			tokenStr = c.PostForm("token")
		}
		if len(tokenStr) == 0 {
			util.PrintLog("token为空")
			util.PrintLog(fmt.Sprintf("token is %v", tokenStr))
			c.JSON(http.StatusUnauthorized, dao.Response{
				StatusCode: -1,
				StatusMsg:  "unauthorized: 用户不存在或者未登录",
			})
			c.Abort()
			return
		}
		//fmt.Printf("%v\n", tokenStr)
		token, ok := ParseToken(tokenStr)
		// 如果token无效
		if !ok {
			util.PrintLog("token无效")
			c.JSON(http.StatusForbidden, dao.Response{
				StatusCode: -1,
				StatusMsg:  "forbidden: token无效",
			})
			c.Abort()
			return
		}

		// 如果token过期
		if time.Now().Unix() > token.ExpiresAt {
			util.PrintLog("token已过期")
			c.JSON(http.StatusOK, dao.Response{
				StatusCode: -1,
				StatusMsg:  "token已过期",
			})

			c.Abort()
			return
		}
		// 将userId写入上下文
		util.PrintLog(fmt.Sprintf(" token没问题 user_id is %v", token.UserId))
		c.Set("user_id", token.UserId)
		// 继续执行
		c.Next()
	}
}

func FilterDirtyMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		_content := c.Query("content")
		content, _ := util.FilterDirty(_content)
		c.Set("content", content)
		c.Next()
	}
}
