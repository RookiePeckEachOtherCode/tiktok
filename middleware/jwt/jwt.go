package jwt

import (
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

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
		//用户不存在
		if tokenStr == "" {
			c.JSON(http.StatusOK, dao.Response{
				StatusCode: 401,
				StatusMsg:  "该用户不存在",
			})
			c.Abort() //阻止执行
			return
		}
		//验证token
		tokenStruck, ok := ParseToken(tokenStr)
		if !ok {
			c.JSON(http.StatusOK, dao.Response{
				StatusCode: 403,
				StatusMsg:  "token不正确",
			})
			c.Abort() //阻止执行
			return
		}
		//token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			c.JSON(http.StatusOK, dao.Response{
				StatusCode: 402,
				StatusMsg:  "token过期",
			})
			c.Abort() //阻止执行
			return
		}
		c.Set("user_id", tokenStruck.UserId)
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
