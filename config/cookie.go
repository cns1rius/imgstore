package config

import (
	"time"
)

const TokenExpireDuration = time.Hour * 2

var MySecret = []byte("Passwd!@#") // 生成签名的密钥
// 登录成功后调用，传入UserInfo结构体

type MyClaims struct {
	User           interface{}
	StandardClaims interface{}
}

//func SetCookie(c *gin.Context, username string) (string, int) {
//	expirationTime := time.Now().Add(TokenExpireDuration) // 两个小时有效期
//	claims := &MyClaims{
//		User: userInfo,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expirationTime.Unix(),
//			Issuer:    "yourname",
//		},
//	}
//	// 生成Token，指定签名算法和claims
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	// 签名
//	if tokenString, err := token.SignedString(MySecret); err != nil {
//		return "", err
//	} else {
//		return tokenString, nil
//	}
//
//	c.SetCookie("gin-cookie", username, 3600, "/", "localhost", true, true)
//	return 0
//}
