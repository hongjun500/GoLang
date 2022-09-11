package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hongjun500/Go-master/gin-web-example/pkg/setting"
	"time"
)

/**
 * @author hongjun500
 * @date 2022/9/11 15:08
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description: jwt
 */

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 生成token 三小时有效期
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	// 过期时间为三小时
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-web-example",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken token的解析并鉴权
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
