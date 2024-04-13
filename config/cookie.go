package config

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

//const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("Passwd!@#") // 生成签名的密钥

// GenJWT 生成JWT
func GenJWT(id uint, isAdmin bool) (string, error) {
	// 创建一个我们自己的声明
	c := jwt.MapClaims{
		"id":       id,
		"is_admin": isAdmin, // 自定义字段
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

func Verify(c *gin.Context) string {
	tokenString, _ := c.Cookie("gin_cookie")
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	switch {
	case err == nil:
		return ""
	case errors.Is(err, jwt.ErrTokenMalformed):
		return "That's not even a token"
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		// Invalid signature
		return "Invalid signature"
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		// Token is either expired or not active yet
		return "Timing is everything"
	default:
		return "Couldn't handle this token"
	}
}

func GetCookieId(c *gin.Context) int {
	cookie, _ := c.Cookie("gin_cookie")
	JWT, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	if err != nil {
		return -1
	}
	id := int(JWT.Claims.(jwt.MapClaims)["id"].(float64))
	return id
}
