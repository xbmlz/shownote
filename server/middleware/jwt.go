package middleware

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"shownote/global"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 4010,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		auth, err := PaserToken(authHeader)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"code": 4012,
					"msg":  "Token已过期",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 4011,
					"msg":  "无效的Token",
				})
			}
			c.Abort()
			return
		}
		c.Set("type", auth.RepoType)
		c.Set("token", auth.AccessToken)
		c.Set("refresh", auth.RefreshToken)
		// 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		c.Next()
	}
}

// SignToken 签发token
func GenToken(auth global.AuthInfo) string {
	authJson, _ := json.Marshal(auth)
	encodeStr := base64.StdEncoding.EncodeToString(authJson)
	return string(encodeStr)
}

// PaserToken 验证token
func PaserToken(tokenStr string) (global.AuthInfo, error) {
	decodeStr, _ := base64.StdEncoding.DecodeString(tokenStr)
	var auth global.AuthInfo
	err := json.Unmarshal(decodeStr, &auth)
	if err != nil {
		return auth, TokenMalformed
	}
	// 验证是否超时
	if time.Now().Unix() >= auth.CreatedAt+auth.ExpiresIn {
		return auth, TokenExpired
	}
	return auth, nil
}
