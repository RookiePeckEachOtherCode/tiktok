package hash

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"tiktok/configs"
	"tiktok/model"

	"github.com/gin-gonic/gin"
)

// HashPassword 中间件 - 对密码进行SHA1哈希
func CheckAndHashPassword() gin.HandlerFunc {
	return func(c *gin.Context) {

		pwd := getPassword(c)

		if len(pwd) == 0 {
			returnError(c, "密码不能为空")
			c.Abort()
			return
		}

		if len(pwd) >= configs.MAX_PASSWORD_LEN {
			returnError(c, "密码长度超过限制")
			c.Abort()
			return
		}

		c.Set("password", sha1Hash(pwd))
		c.Next()
	}
}

// CheckUserName 中间件 - 校验用户名
func CheckUserName() gin.HandlerFunc {
	return func(c *gin.Context) {

		uname := getUserName(c)

		if len(uname) == 0 {
			returnError(c, "用户名不能为空")
			c.Abort()
			return
		}

		if len(uname) >= configs.MAX_NAME_LEN {
			returnError(c, "用户名长度超过限制")
			c.Abort()
			return
		}

		c.Next()
	}
}

// 返回错误响应的统一处理
func returnError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 1,
		StatusMsg:  msg,
	})
}

// 获取密码的统一方法
func getPassword(c *gin.Context) string {
	// 优先从querystring获取
	pwd := c.Query("password")
	if len(pwd) == 0 {
		// 再从form表单获取
		pwd = c.PostForm("password")
	}
	return pwd
}

// 获取用户名的统一方法
func getUserName(c *gin.Context) string {
	// 优先从querystring获取
	uname := c.Query("username")
	if len(uname) == 0 {
		// 再从form表单获取
		uname = c.PostForm("username")
	}
	return uname
}

// SHA1哈希算法
func sha1Hash(str string) string {
	// 实现SHA1哈希...
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
