/*
 * @Time : 2020/11/16 10:30
 * @Author : wangyl
 * @File : Jwt.go
 * @Software: GoLand
 */
package JwtAuth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"xing-doraemon/global"
)

var jwtSecret = []byte(global.GetAlterGatewayConfig().App.Secret)

type Claims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		id,
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "xing",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

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
